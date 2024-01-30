package httpc

import (
	"net/http"
	"time"
)

type Client interface {
	Do(req *http.Request) (*http.Response, error)
	Post(url string, headers http.Header, body Body) (*http.Response, error)
	TelegramCall(method string, token string, params map[string]interface{}) (*http.Response, error)
}

type httpClient struct {
	url     string
	timeout time.Duration
	client  *http.Client
}

func NewHTTPClient(baseUrl string, timeOut time.Duration) Client {
	return &httpClient{
		url:     baseUrl,
		timeout: timeOut,
		client: &http.Client{
			Transport: http.DefaultTransport,
			Timeout:   timeOut,
		},
	}
}
