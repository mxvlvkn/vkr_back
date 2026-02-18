package utils

import (
	"fmt"
	"reflect"
)

func GetStructFieldByString(src any, fieldName string) (any, error) {
    srcVal := reflect.ValueOf(src)

	if srcVal.Kind() == reflect.Pointer && srcVal.IsNil() {
		return nil, fmt.Errorf("%v: Source is nil", fieldName)
	}

	if fieldName == "" {
		return nil, fmt.Errorf("%v: FieldName is empty", fieldName)
	}

	// Разыменование, если указатель
    srcVal = reflect.Indirect(srcVal) 

    if srcVal.Kind() != reflect.Struct{
		return nil, fmt.Errorf("%v: Source is not struct", fieldName)
    }

	field := srcVal.FieldByName(fieldName)
    if field.IsValid() {
		return field.Interface(), nil
	}

	return nil, fmt.Errorf("%v: Field not found", fieldName)
}