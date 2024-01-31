package telebot

import (
	"encoding/json"
	"testing"
)

func TestUser_MarshalUnmarshalJSON(t *testing.T) {
	// Test case 1: Normal case
	user := User{
		ID:           123,
		IsBot:        true,
		FirstName:    "John",
		LastName:     nil,
		Username:     nil,
		LanguageCode: nil,
		IsPremium:    nil,
	}

	// Marshal to JSON
	jsonData, err := json.Marshal(&user)
	if err != nil {
		t.Errorf("Test case 1 failed. Error marshaling JSON: %v", err)
	}

	// Unmarshal from JSON
	var newUser User
	err = json.Unmarshal(jsonData, &newUser)
	if err != nil {
		t.Errorf("Test case 1 failed. Error unmarshaling JSON: %v", err)
	}

	// Check if the unmarshaled user is equal to the original user
	if newUser.ID != user.ID || newUser.IsBot != user.IsBot || newUser.FirstName != user.FirstName {
		t.Errorf("Test case 1 failed. Unmarshaled user is not equal to the original user.")
	}

	// Test case 2: With additional fields set
	userWithFields := User{
		ID:           456,
		IsBot:        false,
		FirstName:    "Jane",
		LastName:     String("Doe"),
		Username:     String("jane_doe"),
		LanguageCode: String("en"),
		IsPremium:    Bool(true),
	}

	// Marshal to JSON
	jsonData2, err := json.Marshal(&userWithFields)
	if err != nil {
		t.Errorf("Test case 2 failed. Error marshaling JSON: %v", err)
	}

	// Unmarshal from JSON
	var newUserWithFields User
	err = json.Unmarshal(jsonData2, &newUserWithFields)
	if err != nil {
		t.Errorf("Test case 2 failed. Error unmarshaling JSON: %v", err)
	}

	// Check if the unmarshaled user is equal to the original user with additional fields
	if newUserWithFields.ID != userWithFields.ID || newUserWithFields.IsBot != userWithFields.IsBot || newUserWithFields.FirstName != userWithFields.FirstName {
		t.Errorf("Test case 2 failed. Unmarshaled user is not equal to the original user with additional fields.")
	}
}

func TestUser_String(t *testing.T) {
	// Test case 1: Bot user
	user := User{
		ID:        123,
		IsBot:     true,
		FirstName: "BotName",
		Username:  String("bot_username"),
	}

	expectedString := "User{(Bot) ID: 123, User: @bot_username}\n{\n  \"id\": 123,\n  \"is_bot\": true,\n  \"first_name\": \"BotName\",\n  \"username\": \"bot_username\"\n}"
	resultString := user.String()

	if resultString != expectedString {
		t.Errorf("Test case 1 failed. Expected:\n%s\nGot:\n%s", expectedString, resultString)
	}

	// Test case 2: Regular user
	user2 := User{
		ID:        456,
		IsBot:     false,
		FirstName: "John",
		Username:  String("john_doe"),
	}

	expectedString2 := "User{ID: 456, User: @john_doe}\n{\n  \"id\": 456,\n  \"is_bot\": false,\n  \"first_name\": \"John\",\n  \"username\": \"john_doe\"\n}"
	resultString2 := user2.String()

	if resultString2 != expectedString2 {
		t.Errorf("Test case 2 failed. Expected:\n%s\nGot:\n%s", expectedString2, resultString2)
	}
}

// Helper functions to create pointers for optional fields

func String(value string) *string {
	return &value
}

func Bool(value bool) *bool {
	return &value
}
