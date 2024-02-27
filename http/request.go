package httpc

import (
	"fmt"
	"go.mamad.dev/gtb/log"
	"net/http"
	"net/url"
)

func (h *httpClient) Do(req *http.Request) (*http.Response, error) {
	return h.client.Do(req)
}

func (h *httpClient) Post(_url string, headers http.Header, body Body) (*http.Response, error) {
	log.Debugf("Content-Type: %#v", body)
	_body, _content := body.Encode()
	log.Debugf("Content-Type: %s", _content)
	baseURL, err := url.Parse(h.url)
	if err != nil {
		return nil, err
	}
	relURL, err := url.Parse(_url)
	if err != nil {
		return nil, err
	}
	fullURL := baseURL.ResolveReference(relURL).String()
	req, err := http.NewRequest(http.MethodPost, fullURL, _body)
	if err != nil {
		return nil, err
	}
	req.Header = headers
	req.Header.Set("Content-Type", _content)
	return h.Do(req)
}

func (h *httpClient) TelegramCall(method string, token string, params map[string]interface{}, files ...File) (*http.Response, error) {
	headers := http.Header{}
	headers.Set("Accept", "application/json")
	b := NewBody(params)
	for _, file := range files {
		b.Add(file.Name, file)
	}
	return h.Post(fmt.Sprintf("/bot%s/%s", token, method), headers, b)
}
