package telebot

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type Verifiable interface {
	Verify() error
}

func splitOptions(options string) []string {
	return splitNonEmpty(options, ",")
}

func splitNonEmpty(s, sep string) []string {
	var result []string
	for _, v := range split(s, sep) {
		if v != "" {
			result = append(result, v)
		}
	}
	return result
}

func split(s, sep string) []string {
	return filterEmpty(strings.Split(s, sep))
}

func filterEmpty(s []string) []string {
	var result []string
	for _, v := range s {
		if v != "" {
			result = append(result, v)
		}
	}
	return result
}

func tagOptions(tag string) []string {
	options := tag[len("oneof:"):]
	return splitOptions(options)
}

func validateEitherExists(field reflect.Value, dependentFieldNames []string) bool {
	existsCount := 0
	for _, depFieldName := range dependentFieldNames {
		depField := field.FieldByName(depFieldName)
		if !depField.IsZero() {
			existsCount++
		}
	}
	return existsCount == 1
}

func verify(t interface{}) error {
	valueOfT := reflect.ValueOf(t)
	if valueOfT.Kind() == reflect.Ptr {
		if valueOfT.IsNil() {
			// Handle nil pointer gracefully if needed
			return errors.New("nil pointer provided")
		}
		valueOfT = valueOfT.Elem()
	}

	var existsCount int
	var hasEitherExistsTag = false
	var eitehrExistsDependentFieldNames []string

	typeOfT := valueOfT.Type()

	for i := 0; i < valueOfT.NumField(); i++ {
		field := valueOfT.Field(i)
		tag := typeOfT.Field(i).Tag.Get("verify")

		if !(field.Kind() == reflect.Ptr && field.Elem().Kind() == reflect.Struct) {
			if field.Kind() == reflect.Struct {
				if err := verify(field.Interface()); err != nil {
					return err
				}
			}
			continue
		}

		switch strings.ToLower(tag) {

		case "required":
			if field.IsZero() || field.IsNil() {
				return fmt.Errorf("field %s is required but has no value", typeOfT.Field(i).Name)
			}
		case "nonzero":
			if field.IsNil() {
				continue // NonZero can be NIL
			}

			switch field.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				if field.Int() == 0 {
					return fmt.Errorf("field %s must be nonzero", typeOfT.Field(i).Name)
				}
			case reflect.Float32, reflect.Float64:
				if field.Float() == 0.0 {
					return fmt.Errorf("field %s must be nonzero", typeOfT.Field(i).Name)
				}

			case reflect.Struct:
				if field.IsZero() {
					return fmt.Errorf("field %s is required but has zero value", typeOfT.Field(i).Name)
				}

			default:
				return fmt.Errorf("field %s has unsupported type %s", typeOfT.Field(i).Name, field.Kind())
			}

		case "oneof":
			validOptions := make(map[string]struct{})
			options := tagOptions(tag)
			for _, option := range options {
				validOptions[option] = struct{}{}
			}
			fieldValue := field.Interface().(string)
			if _, ok := validOptions[fieldValue]; !ok {
				return fmt.Errorf("field %s must be one of %v", typeOfT.Field(i).Name, options)
			}

		case "literalzero":
			switch field.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				if field.Int() != 0 {
					return fmt.Errorf("field %s must have a literal zero value", typeOfT.Field(i).Name)
				}
			case reflect.Float32, reflect.Float64:
				if field.Float() != 0.0 {
					return fmt.Errorf("field %s must have a literal zero value", typeOfT.Field(i).Name)
				}
			default:
				return fmt.Errorf("field %s has unsupported type %s, literalZero is only supported for ints and floats", typeOfT.Field(i).Name, field.Kind())
			}

		case "exists_if":
			dependentField := valueOfT.FieldByName(tag)
			if dependentField.IsZero() {
				return fmt.Errorf("field %s depends on %s, which has zero value", typeOfT.Field(i).Name, tag)
			}

		case "zero_if":
			dependentField := valueOfT.FieldByName(tag)
			if !dependentField.IsZero() && field.IsZero() {
				return fmt.Errorf("field %s must be zero because %s is not zero", typeOfT.Field(i).Name, tag)
			}

		case "nil_if":
			dependentField := valueOfT.FieldByName(tag)
			if !dependentField.IsZero() && field.IsNil() {
				return fmt.Errorf("field %s must be nil because %s is not zero", typeOfT.Field(i).Name, tag)
			}

		case "exists_if_not_exists":
			dependentField := valueOfT.FieldByName(tag)
			if !dependentField.IsZero() && field.IsZero() {
				return fmt.Errorf("field %s must exist because %s exists", typeOfT.Field(i).Name, tag)
			}
		case "nil_if_not_nil":
			dependentField := valueOfT.FieldByName(tag)
			if !dependentField.IsNil() && field.IsNil() {
				return fmt.Errorf("field %s must be nil because %s is not nil", typeOfT.Field(i).Name, tag)
			}

		case "either_exists":
			hasEitherExistsTag = true
			dependentFieldName := strings.Split(tag, ":")[1]
			dependentField := field.FieldByName(dependentFieldName)
			eitehrExistsDependentFieldNames = append(eitehrExistsDependentFieldNames, tag)
			if !dependentField.IsZero() {
				existsCount++
			}

		case "not_nil_if_nil":
			dependentField := valueOfT.FieldByName(tag)
			if dependentField.IsNil() && !field.IsNil() {
				return fmt.Errorf("field %s must not be nil because %s is nil", typeOfT.Field(i).Name, tag)
			}
		case "not_exists_if_exists":
			dependentField := valueOfT.FieldByName(tag)
			if dependentField.IsZero() {
				return fmt.Errorf("field %s must not exist because %s exists", typeOfT.Field(i).Name, tag)
			}
		}
	}

	if hasEitherExistsTag && existsCount != 1 {
		return fmt.Errorf("exactly one of the %s should exist", strings.Join(eitehrExistsDependentFieldNames, ", "))
	}

	return nil
}
