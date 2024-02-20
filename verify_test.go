package telebot

import (
	"reflect"
	"testing"
)

func TestReflections(t *testing.T) {
	typ := MaybeInaccessibleMessage{} // Create an instance of the struct directly

	valueOfT := reflect.ValueOf(typ)
	typeOfT := reflect.TypeOf(typ)
	numField := valueOfT.NumField()

	t.Logf("typeOf: %s\n", typeOfT.String())
	t.Logf("numField: %d\n", numField)

	field := typeOfT.Field(0)

	t.Logf("field: %s\n", field.Name)
}
