package telebot

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (b *bot) GetFile(fileID string) (*File, error) {
	var resp struct {
		Result File
	}

	params := getFileRequest{
		FileID: fileID,
	}

	req, err := b.sendMethodRequest(methodGetFile, params)

	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(req, &resp); err != nil {
		return nil, err
	}

	return &resp.Result, nil
}

// URL currently only supports TELEGRAM's OFFICIAL API URL
// todo: implement SelfHosted API Support
func (f *File) URL() string {
	return "https://api.telegram.org/file/bot" + telegramSecretToken + "/" + f.FilePath
}

func (f *File) Download() ([]byte, error) {
	resp, err := http.Get(f.URL())

	if err != nil {
		return nil, err
	}

	defer panicIf(wrapError(resp.Body.Close()))

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if len(body) > 0 {
		return body, nil
	}

	return nil, errors.New("failed to download file")
}
