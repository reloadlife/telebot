package telebot

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	httpc "go.mamad.dev/gtb/http"
	"go.mamad.dev/gtb/log"
	"io"
	"net/http"
	"reflect"
	"strings"
	"time"
)

func (b *bot) sendMethodRequest(method method, params any, files ...httpc.File) ([]byte, error) {
	if b.offlineMode {
		return nil, nil
	}

	v, ok := params.(Verifiable)
	if ok {
		if err := v.Verify(); err != nil {
			return nil, err
		}
	}

	req, err := b.httpClient.TelegramCall(method.String(), b.token, structToMap(params), files...)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	return body, extractOk(body)
}

func (b *bot) Raw(method string, params map[string]any) (*http.Response, error) {
	return b.httpClient.TelegramCall(method, b.token, params)
}

func structToMap(input any) map[string]any {
	result := make(map[string]any)

	val := reflect.ValueOf(input)
	if val.Kind() != reflect.Struct {
		// Handle the case when input is not a struct
		return result
	}

	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		jsonTag := field.Tag.Get("json")

		if jsonTag == "" || jsonTag == "-" {
			continue
		}

		v := val.Field(i).Interface()
		if strings.Contains(jsonTag, "omitempty") {
			if reflect.DeepEqual(v, reflect.Zero(field.Type).Interface()) {
				continue
			}
		}

		jsonTag = strings.ReplaceAll(jsonTag, ",omitempty", "")

		switch vi := v.(type) {
		case *ReplyParameters:
			result[jsonTag] = structToMap(*vi)
			continue

		case time.Duration:
			result[jsonTag] = int(vi.Seconds())
			continue

		case Recipient:
			result[jsonTag] = vi.Recipient()
			continue

		case Userable:
			result[jsonTag] = vi.User()

		case *File:
			if vi.InCloud() {
				result[jsonTag] = vi.FileID
				continue
			}

			if vi.OnDisk() {
				result[jsonTag] = vi
				continue
			}

			result[jsonTag] = vi.GetFileURL()
			continue

		case []InputMedia:
			medias := make([]map[string]any, 0)
			for _, m := range vi {
				media := structToMap(m)
				media["type"] = m.MediaType
				if m.Repr != "" {
					media["media"] = m.FileRepresent()
				}
				medias = append(medias, media)
				log.Infof("media: %v", media)
			}
			result[jsonTag] = medias
			continue
		case []InputSticker:
			stickers := make([]map[string]any, 0)
			for _, s := range vi {
				sticker := structToMap(s)
				sticker["sticker"] = s.FileRepresent()
				stickers = append(stickers, sticker)
			}
			result[jsonTag] = stickers
			continue
		case InputSticker:
			sticker := structToMap(vi)
			sticker["sticker"] = vi.FileRepresent()
			result[jsonTag] = sticker
			continue

		case []StickerEmoji:
			var list []string
			for _, s := range vi {
				list = append(list, s.String())
			}
			result[jsonTag] = list
			continue
		}

		result[jsonTag] = v
	}

	return result
}

func extractOk(data []byte) error {

	var e struct {
		Ok          bool           `json:"ok"`
		Code        int            `json:"error_code"`
		Description string         `json:"description"`
		Parameters  map[string]any `json:"parameters"`
	}
	if json.NewDecoder(bytes.NewReader(data)).Decode(&e) != nil {
		return nil
	}

	if e.Ok {
		return nil
	}

	err := Err(e.Description)
	switch err {
	case nil:
	case ErrGroupMigrated:
		migratedTo, ok := e.Parameters["migrate_to_chat_id"]
		if !ok {
			return NewError(e.Code, e.Description)
		}

		return GroupError{
			err:        err.(*Error),
			MigratedTo: int64(migratedTo.(float64)),
		}
	default:
		return err
	}

	switch e.Code {
	case http.StatusTooManyRequests:
		retryAfter, ok := e.Parameters["retry_after"]
		if !ok {
			return NewError(e.Code, e.Description)
		}

		err = FloodError{
			err:        NewError(e.Code, e.Description),
			RetryAfter: int(retryAfter.(float64)),
		}
	default:
		err = fmt.Errorf("telegram: %s (%d)", e.Description, e.Code)
	}

	return err
}

func (b *bot) resAsMessage(data []byte, errs ...error) (*AccessibleMessage, error) {
	var m struct {
		Result AccessibleMessage `json:"result"`
	}
	err := json.Unmarshal(data, &m)
	if err != nil {
		return nil, errors.Join(append(errs, err)...)
	}
	return &m.Result, errors.Join(errs...)
}
