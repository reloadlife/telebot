package telebot

import (
	"net/http"
	"reflect"
)

func (b *bot) sendMethodRequest(method method, params any) (*http.Response, error) {
	return b.httpClient.TelegramCall(method.String(), b.token, structToMap(params))
}

func (b *bot) Raw(method string, params map[string]any) (*http.Response, error) {
	return b.httpClient.TelegramCall(method, b.token, params)
}

func structToMap(input any) map[string]any {
	result := make(map[string]any)

	val := reflect.ValueOf(input)
	typ := reflect.TypeOf(input)

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" || jsonTag == "-" {
			continue
		}

		result[jsonTag] = val.Field(i).Interface()
	}

	return result
}
