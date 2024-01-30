package httpc

import (
	"testing"
	"time"
)

func TestRequests(t *testing.T) {
	c := NewHTTPClient("https://example.com", time.Minute)

	call, err := c.Post("/", map[string][]string{}, NewBody(nil))
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	if call.StatusCode != 200 {
		t.Fatalf("Error: %v", call.StatusCode)
	}
}
