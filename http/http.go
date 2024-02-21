package httpc

import (
	"net/http"
	"os"
	"time"
)

const DefaultTelegramAPIURL = "https://api.telegram.org"

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

func NewHTTPClient(baseURL string, timeOut time.Duration) Client {
	if baseURL == "" {
		defaultURL, has := os.LookupEnv("TELEBOT_API_URL")
		baseURL = DefaultTelegramAPIURL
		if has {
			baseURL = defaultURL
		}
	}
	return &httpClient{
		url:     baseURL,
		timeout: timeOut,
		client: &http.Client{
			Transport: &LoggingTransport{
				Transport: http.DefaultTransport,
			},
			Timeout: timeOut,
		},
	}
}
