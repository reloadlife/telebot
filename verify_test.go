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

func TestVerify(t *testing.T) {
	type MyStruct struct {
		Field1  string `verify:"required"`
		Field2  int    `verify:"nonzero"`
		Field3  float64
		Field4  string  `verify:"oneof:option1,option2"`
		Field5  int     `verify:"literalZero"`
		Field6  int     `verify:"exists_if:Field3"`
		Field7  string  `verify:"zero_if:Field5"`
		Field8  *string `verify:"nil_if:Field1"`
		Field9  string  `verify:"exists_if_not_exists:Field10"`
		Field10 string  // This field is checked for existence based on Field9
		Field11 *string `verify:"nil_if_not_nil:Field12"`
		Field12 *string // This field is checked for nil based on Field11
	}

	myStruct := MyStruct{
		Field1:  "value",
		Field2:  42,
		Field3:  3.14,
		Field4:  "option1",
		Field5:  0,
		Field6:  1,
		Field7:  "",
		Field8:  nil,
		Field9:  "dependency",
		Field10: "exists",
		Field11: nil,
		Field12: toPtr("not_nil"),
	}

	err := verify(myStruct)
	if err != nil {
		t.Logf("Verification failed: %s", err)
	} else {
		t.Logf("Verification passed")
	}
}
