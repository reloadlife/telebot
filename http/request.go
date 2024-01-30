package httpc

import (
	"fmt"
	"net/http"
)

func (h *httpClient) Do(req *http.Request) (*http.Response, error) {
	return h.client.Do(req)
}

func (h *httpClient) Post(url string, headers http.Header, body Body) (*http.Response, error) {
	_body, _content := body.Encode()
	URL := h.url + url
	req, err := http.NewRequest(http.MethodPost, URL, _body)
	if err != nil {
		return nil, err
	}
	req.Header = headers
	req.Header.Set("Content-Type", _content)
	return h.Do(req)
}

func (h *httpClient) TelegramCall(method string, token string, params map[string]interface{}) (*http.Response, error) {
	headers := http.Header{}
	headers.Set("Accept", "application/json")
	return h.Post(fmt.Sprintf("bot%s/%s", token, method), headers, NewBody(params))
}
