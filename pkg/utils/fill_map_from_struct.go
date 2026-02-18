package utils

import (
	"fmt"
	"reflect"
)

func FillMapFromStruct(src any, dict *map[string]string) (map[string]any, error) {
    if src == nil {
        return nil, fmt.Errorf("Source is nil")
    }
    if dict == nil || *dict == nil {
        return nil, fmt.Errorf("Dict is nil")
    }

    srcVal := reflect.ValueOf(src)
    srcVal = reflect.Indirect(srcVal) 

    if srcVal.Kind() != reflect.Struct {
        return nil, fmt.Errorf("Source is not a struct")
    }

    result := make(map[string]any)

    for dbField, structField := range *dict {
        field := srcVal.FieldByName(structField)
        if !field.IsValid() {
            return nil, fmt.Errorf("%s: Not found in struct", structField)
        }

        if !field.CanInterface() {
            return nil, fmt.Errorf("%s: Cannot get value (type %v)", structField, field.Type())
        }

        value := field.Interface()

        if value == nil {
            return nil, fmt.Errorf("%s: Value is nil", structField)
        }  

        result[dbField] = value
    }

    return result, nil
}