package httpc

import (
	"os"
	"testing"
	"time"
)

func TestHttp(t *testing.T) {
	c := NewHTTPClient("https://api.telegram.org/", time.Minute)

	telegramToken := os.Getenv("TELEBOT_SECRET")

	call, err := c.TelegramCall("getMe", telegramToken, nil)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	if call.StatusCode != 200 {
		t.Fatalf("Error: %v", call.StatusCode)
	}
}
