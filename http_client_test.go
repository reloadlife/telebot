package telebot

import (
	"net/http"
	"reflect"
	"testing"
)

type TestStruct struct {
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Score float64 `json:"-"`
}

// Mock HTTPClient for testing purposes
type mockHTTPClient struct{}

func (m *mockHTTPClient) TelegramCall(method, token string, params map[string]interface{}) (*http.Response, error) {
	// Implement the mock behavior here
	// Return a mock http.Response and error for testing
	return nil, nil
}

func TestSendMethodRequest(t *testing.T) {
	// Create an instance of your bot with the mockHTTPClient
	b := &bot{
		httpClient: &mockHTTPClient{},
		token:      "your_token",
	}

	// Test case 1: Normal case
	method := SomeMethod // Replace with your actual method type
	params := SomeParams // Replace with your actual params type
	_, err := b.sendMethodRequest(method, params)
	if err != nil {
		t.Errorf("Test case 1 failed. Expected no error, got %v", err)
	}

	// Test case 2: Error case (modify the mockHTTPClient to return an error)
	// ...

	// Add more test cases as needed
}

func TestRaw(t *testing.T) {
	// Create an instance of your bot with the mockHTTPClient
	b := &bot{
		httpClient: &mockHTTPClient{},
		token:      "your_token",
	}

	// Test case 1: Normal case
	method := "some_method"                          // Replace with your actual method string
	params := map[string]interface{}{"key": "value"} // Replace with your actual params
	_, err := b.Raw(method, params)
	if err != nil {
		t.Errorf("Test case 1 failed. Expected no error, got %v", err)
	}

	// Test case 2: Error case (modify the mockHTTPClient to return an error)
	// ...

	// Add more test cases as needed
}

func TestStructToMap(t *testing.T) {
	// Test case 1: Normal case
	input := TestStruct{Name: "John", Age: 25, Score: 90.5}
	expected := map[string]interface{}{"name": "John", "age": 25}
	result := structToMap(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case 1 failed. Expected %v, got %v", expected, result)
	}

	// Test case 2: Struct with a field having json tag "-"
	input2 := TestStruct{Name: "Alice", Age: 30, Score: 85.0}
	expected2 := map[string]interface{}{"name": "Alice", "age": 30}
	result2 := structToMap(input2)
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("Test case 2 failed. Expected %v, got %v", expected2, result2)
	}

	// Test case 3: Empty struct
	input3 := struct{}{}
	expected3 := map[string]interface{}{}
	result3 := structToMap(input3)
	if !reflect.DeepEqual(result3, expected3) {
		t.Errorf("Test case 3 failed. Expected %v, got %v", expected3, result3)
	}
}
