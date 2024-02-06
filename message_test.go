package telebot

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestMessageSerialization(t *testing.T) {
	t.Run("Marshal and Unmarshal", func(t *testing.T) {
		// Create a sample Message instance for testing
		message := &Message{
			ID:   123,
			Date: 1627954567,
			Chat: &Chat{
				ID:   456,
				Type: "private",
			},
			Text: stringPtr("Hello, world!"),
		}

		// Serialize the Message to JSON
		jsonData, err := json.Marshal(message)
		if err != nil {
			t.Fatalf("Failed to marshal Message to JSON: %v", err)
		}

		// Deserialize the JSON back to Message
		var deserializedMessage Message
		err = json.Unmarshal(jsonData, &deserializedMessage)
		if err != nil {
			t.Fatalf("Failed to unmarshal JSON to Message: %v", err)
		}

		// Check if the original and deserialized messages are equal
		if !areMessagesEqual(message, &deserializedMessage) {
			t.Errorf("Original and deserialized messages are not equal.")
		}
	})

	t.Run("String representation", func(t *testing.T) {
		// Create a sample Message instance for testing
		message := &Message{
			ID:   123,
			Date: 1627954567,
			Chat: &Chat{
				ID:   456,
				Type: "private",
			},
			Text: stringPtr("Hello, world!"),
		}

		// Ensure String() method does not panic
		strRepresentation := message.String()
		if strRepresentation == "" {
			t.Error("String representation is empty.")
		}
	})
}

func TestMessageEquality(t *testing.T) {
	// Test equality with different scenarios
	t.Run("Equal Messages", func(t *testing.T) {
		message1 := &Message{ID: 1, Date: 1627954567}
		message2 := &Message{ID: 1, Date: 1627954567}

		if !areMessagesEqual(message1, message2) {
			t.Error("Equal messages are not considered equal.")
		}
	})

	t.Run("Different Messages", func(t *testing.T) {
		message1 := &Message{ID: 1, Date: 1627954567}
		message2 := &Message{ID: 2, Date: 1627954567}

		if areMessagesEqual(message1, message2) {
			t.Error("Different messages are considered equal.")
		}
	})
}

func areMessagesEqual(m1, m2 *Message) bool {
	if m1 == nil || m2 == nil {
		return m1 == m2
	}

	// Get the reflect.Value of m1 and m2
	val1 := reflect.ValueOf(m1).Elem()
	val2 := reflect.ValueOf(m2).Elem()

	// Iterate over fields
	for i := 0; i < val1.NumField(); i++ {
		field1 := val1.Field(i)
		field2 := val2.Field(i)

		// Skip comparison if field is unexported (lowercase)
		if !field1.CanInterface() || !field2.CanInterface() {
			continue
		}

		// Skip comparison if field is a pointer to a struct and is nil in both structs
		if field1.Kind() == reflect.Ptr && field1.IsNil() && field2.Kind() == reflect.Ptr && field2.IsNil() {
			continue
		}

		// Perform field comparison
		if !reflect.DeepEqual(field1.Interface(), field2.Interface()) {
			return false
		}
	}

	return true
}

func stringPtr(s string) *string {
	return &s
}
