package telebot

import (
	"fmt"
	"reflect"
	"time"
)

func setFieldByName(obj any, fieldName string, value any) error {
	structValue := reflect.ValueOf(obj).Elem()
	field := structValue.FieldByName(fieldName)

	if !field.IsValid() {
		return fmt.Errorf("field '%s' not found in struct", fieldName)
	}

	if !field.CanSet() {
		return fmt.Errorf("cannot set value for field '%s'", fieldName)
	}

	fieldType := field.Type()
	newValue := reflect.ValueOf(value)

	if !newValue.Type().AssignableTo(fieldType) {
		return fmt.Errorf("type mismatch for field '%s'", fieldName)
	}

	field.Set(newValue)
	return nil
}

func appendSliceFieldByName(obj any, fieldName string, values ...any) error {
	structValue := reflect.ValueOf(obj).Elem()
	field := structValue.FieldByName(fieldName)

	if !field.IsValid() {
		return fmt.Errorf("field '%s' not found in struct", fieldName)
	}

	if field.Kind() != reflect.Slice {
		return fmt.Errorf("field '%s' is not a slice type", fieldName)
	}

	for _, value := range values {
		elementValue := reflect.ValueOf(value)

		if !elementValue.Type().AssignableTo(field.Type().Elem()) {
			return fmt.Errorf("type mismatch for field '%s'", fieldName)
		}

		field.Set(reflect.Append(field, elementValue))
	}

	return nil
}

func format(params any, options ...any) bool {
	param := reflect.TypeOf(params)
	if param.Kind() != reflect.Ptr {
		panic("telebot: first argument must be a pointer")
	}

	hasAnythingChanged := false

	for _, option := range options {
		var err error
		switch v := option.(type) {
		case *MessageThreadID:
			err = setFieldByName(params, "ThreadID", v)
			if err == nil {
				hasAnythingChanged = true
			}
		case ReplyMarkup:
			err = setFieldByName(params, "ReplyMarkup", v)
			if err == nil {
				hasAnythingChanged = true
			}
		case *ReplyParameters:
			err = setFieldByName(params, "ReplyParameters", v)
			if err == nil {
				hasAnythingChanged = true
			}
		case ParseMode:
			err = setFieldByName(params, "ParseMode", v)
			if err == nil {
				hasAnythingChanged = true
			}
		case Entity:
			err = appendSliceFieldByName(params, "Entities", v)
			if err == nil {
				hasAnythingChanged = true
			}
		case []Entity:
			var entities []any
			for _, entity := range v {
				entities = append(entities, entity)
			}
			err = appendSliceFieldByName(params, "Entities", entities...)
			if err == nil {
				hasAnythingChanged = true
			}

		case time.Duration:
			err = setFieldByName(params, "Duration", v)
			if err == nil {
				hasAnythingChanged = true
			}

		case Performer:
			err = setFieldByName(params, "Performer", toPtr(string(v)))
			if err == nil {
				hasAnythingChanged = true
			}

		case Title:
			err = setFieldByName(params, "Title", toPtr(string(v)))
			if err == nil {
				hasAnythingChanged = true
			}

		case Width:
			err = setFieldByName(params, "Width", toPtr(int(v)))
		case Height:
			err = setFieldByName(params, "Height", toPtr(int(v)))

		case Option:
			switch v {
			case Silent:
				err = setFieldByName(params, "DisableNotification", toPtr(true))
				if err == nil {
					hasAnythingChanged = true
				}
			case Protected:
				err = setFieldByName(params, "Protected", toPtr(true))
				if err == nil {
					hasAnythingChanged = true
				}
			case HasSpoiler:
				err = setFieldByName(params, "HasSpoiler", toPtr(true))
				if err == nil {
					hasAnythingChanged = true
				}
			}
		}
		if err != nil {
			panic("telebot: " + err.Error())
		}
	}

	return hasAnythingChanged
}
