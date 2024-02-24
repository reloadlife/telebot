package httpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"reflect"
	"strings"
)

type FileeRepresenter interface {
	FileRepresent() string
}

type Body interface {
	MarshalJSON() ([]byte, error)
	Add(key string, value any)
	Get(key string) any
	Delete(key string)
	Encode() (io.Reader, string)
}

type Uploadable interface {
	InCloud() bool
	OnDisk() bool

	GetFileID() string
	GetFileSize() int64
	GetFilePath() string
	GetFileLocal() string
	GetFileURL() string
	GetFileName() string
	GetFileReader() io.Reader
}

type body struct {
	b map[string]any

	hasFiles bool
	files    []File

	buffer *bytes.Buffer // Use bytes.Buffer instead of io.Pipe

	w *multipart.Writer

	isNil bool
}

func NewBody(p map[string]any) Body {
	files := make([]File, 0)
	for k, l := range p {
		switch v := l.(type) {
		case Uploadable:
			if v.GetFileReader() != nil {
				_b, err := io.ReadAll(v.GetFileReader())
				if err != nil {
					panic("httpc: failed to read File")
				}
				files = append(files, File{
					Name:     k,
					DATA:     _b,
					FileName: v.GetFileName(),
				})
				delete(p, k)
				continue
			}
		}
	}
	return &body{
		files:    files,
		hasFiles: len(files) > 0,
		b:        p,
		isNil:    p == nil,
		buffer:   &bytes.Buffer{},
	}
}

type File struct {
	Name     string
	FileName string
	DATA     []byte
}

func (b *body) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.b)
}

func (b *body) Add(key string, value any) {
	switch v := value.(type) {
	case []byte:
		b.hasFiles = true
		b.files = append(b.files, File{
			Name: key,
			DATA: v,
		})
	case Uploadable:
		switch {
		case v.InCloud():
			b.b[key] = v.GetFileID()
		case v.OnDisk():
			fileData, err := os.ReadFile(v.GetFileLocal())
			if err != nil {
				panic("httpc: failed to read File")
			}
			b.files = append(b.files, File{
				Name: key,
				DATA: fileData,
			})
		case v.GetFileReader() != nil:
			fileData, err := io.ReadAll(v.GetFileReader())
			if err != nil {
				panic("httpc: failed to read File")
			}

			b.files = append(b.files, File{
				Name: key,
				DATA: fileData,
			})

		case v.GetFileURL() != "":
			b.b[key] = v.GetFileURL()

		default:
			panic("httpc: File for field " + key + " doesn't exist")
		}

	default:
		b.b[key] = value
	}
}

func (b *body) Get(key string) any {
	return b.b[key]
}

func (b *body) Delete(key string) {
	delete(b.b, key)
}

func (b *body) Encode() (io.Reader, string) {
	if b.isNil {
		return nil, "application/json"
	}

	if b.hasFiles {
		b.w = multipart.NewWriter(b.buffer)

		for _, f := range b.files {
			part, err := b.w.CreateFormFile(f.Name, f.FileName)
			if err != nil {
				panic("httpc: failed to create form File")
			}
			fileType := http.DetectContentType(f.DATA)
			if !strings.Contains(fileType, "image") {
				_, _ = part.Write([]byte(fmt.Sprintf("Content-Type: %s\r\n\r\n", fileType)))
			}
			_, _ = part.Write(f.DATA)
		}

		for key, value := range b.b {
			var strValue string
			switch v := value.(type) {
			case string:
				strValue = v
			case *string:
				if v != nil {
					strValue = *v
				}
			case int, int32, int64, float32, float64:
				strValue = fmt.Sprintf("%v", v)
			case json.Marshaler:
				j, err := v.MarshalJSON()
				if err != nil {
					panic("httpc: failed to encode field value to json: " + key + " " + err.Error())
				}
				strValue = string(j)
			case fmt.Stringer:
				strValue = v.String()

			default:
				t := reflect.TypeOf(v)
				if t.Kind() == reflect.Ptr {
					strValue = fmt.Sprintf("%p", v)
				} else {
					strValue = fmt.Sprintf("%v", v)
				}
			}

			err := b.w.WriteField(key, strValue)
			if err != nil {
				panic("httpc: failed to add form field")
			}
		}

		wrapError(b.w.Close())
		return b.buffer, b.w.FormDataContentType()
	}

	j, err := b.MarshalJSON()
	if err != nil {
		panic("httpc: failed to encode body!")
	}
	return bytes.NewReader(j), "application/json"
}

func wrapError(err error) {
	if err != nil {
		panic("httpc: " + err.Error())
	}
}
