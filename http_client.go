package telebot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
)

func (b *bot) sendMethodRequest(method method, params any) ([]byte, error) {
	if b.offlineMode {
		return nil, nil
	}

	req, err := b.httpClient.TelegramCall(method.String(), b.token, structToMap(params))
	if err != nil {
		return nil, err
	}
	defer panicIf(wrapError(req.Body.Close()))
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

		result[jsonTag] = val.Field(i).Interface()
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
