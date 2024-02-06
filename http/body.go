package httpc

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"os"
)

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
	files    []file

	buffer *bytes.Buffer // Use bytes.Buffer instead of io.Pipe

	w *multipart.Writer

	isNil bool
}

func NewBody(p map[string]any) Body {
	return &body{
		b:      p,
		isNil:  p == nil,
		buffer: &bytes.Buffer{},
	}
}

type file struct {
	name string
	data []byte
}

func (b *body) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.b)
}

func (b *body) Add(key string, value any) {
	switch value.(type) {
	case []byte:
		b.hasFiles = true
		b.files = append(b.files, file{
			name: key,
			data: value.([]byte),
		})
	case Uploadable:
		v := value.(Uploadable)
		switch {
		case v.InCloud():
			b.b[key] = v.GetFileID()
		case v.OnDisk():
			fileData, err := os.ReadFile(v.GetFileLocal())
			if err != nil {
				panic("httpc: failed to read file")
			}
			b.files = append(b.files, file{
				name: key,
				data: fileData,
			})
		case v.GetFileReader() != nil:
			fileData, err := io.ReadAll(v.GetFileReader())
			if err != nil {
				panic("httpc: failed to read file")
			}

			b.files = append(b.files, file{
				name: key,
				data: fileData,
			})

		case v.GetFileURL() != "":
			b.b[key] = v.GetFileURL()

		default:
			panic("httpc: file for field " + key + " doesn't exist")
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
			part, err := b.w.CreateFormFile(f.name, f.name)
			if err != nil {
				panic("httpc: failed to create form file")
			}
			_, _ = part.Write(f.data)
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
