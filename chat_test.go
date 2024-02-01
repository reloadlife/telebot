package telebot

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChat_MarshalJSON(t *testing.T) {
	chat := &Chat{
		ID:    123,
		Type:  "private",
		Title: nil,
		// Add more fields with sample values for testing
	}

	expectedJSON := `{"id":123,"type":"private"}`

	resultJSON, err := json.Marshal(chat)
	assert.NoError(t, err)
	assert.Equal(t, expectedJSON, string(resultJSON))
}

func TestChat_UnmarshalJSON(t *testing.T) {
	jsonStr := `{"id":123,"type":"private"}`

	expectedChat := &Chat{
		ID:    123,
		Type:  "private",
		Title: nil,
		// Add more fields with expected values
	}

	var resultChat Chat
	err := json.Unmarshal([]byte(jsonStr), &resultChat)
	assert.NoError(t, err)
	assert.Equal(t, expectedChat, &resultChat)
}

func TestChat_String(t *testing.T) {
	chat := &Chat{
		ID:    123,
		Type:  "private",
		Title: nil,
		// Add more fields with sample values for testing
	}

	expectedString := "Chat{ID: (private) 123, Title: @<nil>}\n{}"

	resultString := chat.String()
	assert.Equal(t, expectedString, resultString)
}

func TestBan(t *testing.T) {

}

func TestUnban(t *testing.T) {

}
