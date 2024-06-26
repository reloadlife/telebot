package telebot

import (
	"fmt"
	"io"
	"os"
)

// File object represents any sort of file.
type File struct {
	FileID   string `json:"file_id"`
	UniqueID string `json:"file_unique_id"`
	FileSize int64  `json:"file_size"`

	// FilePath is used for files on Telegram server.
	FilePath string `json:"file_path"`

	// FileLocal is used for files on local file system.
	FileLocal string `json:"file_local"`

	// FileURL is used for file on the internet.
	FileURL string `json:"file_url"`

	// FileReader is used for file backed with io.Reader.
	FileReader io.Reader `json:"-"`

	fileName string
}

// FromDisk constructs a new local (on-disk) file object.
//
// Note, it returns File, not *File for a very good reason:
// in telebot, File is pretty much an embeddable struct,
// so upon uploading media you'll need to set embedded File
// with something. NewFile() returning File makes it a one-liner.
//
//	photo := &tele.Photo{File: tele.FromDisk("chicken.jpg")}
func FromDisk(filename string) File {
	return File{FileLocal: filename}
}

// FromURL constructs a new file on provided HTTP URL.
//
// Note, it returns File, not *File for a very good reason:
// in telebot, File is pretty much an embeddable struct,
// so upon uploading media you'll need to set embedded File
// with something. NewFile() returning File makes it a one-liner.
//
//	photo := &tele.Photo{File: tele.FromURL("https://site.com/picture.jpg")}
func FromURL(url string) File {
	return File{FileURL: url}
}

// FromReader constructs a new file from io.Reader.
//
// Note, it returns File, not *File for a very good reason:
// in telebot, File is pretty much an embeddable struct,
// so upon uploading media you'll need to set embedded File
// with something. NewFile() returning File makes it a one-liner.
//
//	photo := &tele.Photo{File: tele.FromReader(bytes.NewReader(...))}
func FromReader(reader io.Reader) File {
	return File{FileReader: reader}
}

// FromFileID constructs a new file from file_id.
func FromFileID(fileID string) File {
	return File{FileID: fileID}
}

func (f *File) stealRef(g *File) {
	if g.OnDisk() {
		f.FileLocal = g.FileLocal
	}

	if g.FileURL != "" {
		f.FileURL = g.FileURL
	}
}

// InCloud tells whether the file is present on Telegram servers.
func (f *File) InCloud() bool {
	return f.FileID != ""
}

// OnDisk will return true if file is present on disk.
func (f *File) OnDisk() bool {
	_, err := os.Stat(f.FileLocal)
	return err == nil
}

func (f *File) GetFileID() string {
	return f.FileID
}
func (f *File) GetFileSize() int64 {
	return f.FileSize
}
func (f *File) GetFilePath() string {
	return f.FilePath
}
func (f *File) GetFileLocal() string {
	return f.FileLocal
}
func (f *File) GetFileURL() string {
	return f.FileURL
}
func (f *File) GetFileName() string {
	return f.fileName
}
func (f *File) GetFileReader() io.Reader {
	if f.OnDisk() {
		file, err := os.Open(f.FileLocal)
		if err != nil {
			return nil
		}
		return file
	}
	return f.FileReader
}

func (f *File) ReflectType() string {
	return fmt.Sprintf("%T", f)
}
func (f *File) Type() string {
	return "File"
}

func (f *File) SetFileName(fileName string) {
	f.fileName = fileName
}
