package utils

import (
	"fmt"
	"reflect"
)

func FillStructFromStruct(src, dst any) error {
    srcVal := reflect.ValueOf(src)
    dstVal := reflect.ValueOf(dst)

    if dstVal.Kind() != reflect.Pointer {
        return fmt.Errorf("Destination is not ptr")
    }
	if srcVal.Kind() == reflect.Pointer && srcVal.IsNil() {
		return fmt.Errorf("Source is nil")
	}
	if dstVal.IsNil() {
		return fmt.Errorf("Destination is nil")
    }

	// Разыменование, если указатели
    srcVal = reflect.Indirect(srcVal) 
    dstVal = reflect.Indirect(dstVal) 

    if srcVal.Kind() != reflect.Struct{
		return fmt.Errorf("Source is not struct")
    }
	if dstVal.Kind() != reflect.Struct {
		return fmt.Errorf("Destination is not struct")
    }

    for i := 0; i < srcVal.NumField(); i++ {
        srcField := srcVal.Field(i)
        srcType := srcVal.Type().Field(i)

        if !srcType.IsExported() {
            continue
        }

        dstField := dstVal.FieldByName(srcType.Name)
        if !dstField.IsValid() {
            continue
        }
		if !dstField.CanSet() {
			return fmt.Errorf("Destination.%v is not set", srcType.Name)
        }

        if srcField.Type().AssignableTo(dstField.Type()) {
            dstField.Set(srcField)
        } else {
			return fmt.Errorf("Destination.%v different types: %v, %v", 
				srcType.Name, 
				srcField.Type().String(),
				dstField.Type().String(),
			)
		}
    }

	return nil
}