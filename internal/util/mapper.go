package util

import "reflect"

func CopyStruct(src interface{}, dst interface{}) {
	srcVal := reflect.ValueOf(src)
	dstVal := reflect.ValueOf(dst)

	if srcVal.Kind() == reflect.Ptr {
		if srcVal.IsNil() {
			return
		}
		srcVal = srcVal.Elem()
	}

	if dstVal.Kind() != reflect.Ptr || dstVal.Elem().Kind() != reflect.Struct {
		return
	}

	dstVal = dstVal.Elem()

	for i := 0; i < srcVal.NumField(); i++ {
		srcField := srcVal.Type().Field(i)
		srcFieldValue := srcVal.Field(i)

		dstFieldValue := dstVal.FieldByName(srcField.Name)

		if dstFieldValue.IsValid() && dstFieldValue.CanSet() {
			if dstFieldValue.Type() == srcFieldValue.Type() {
				dstFieldValue.Set(srcFieldValue)
			} else if srcFieldValue.Kind() == reflect.Ptr && !srcFieldValue.IsNil() {
				if dstFieldValue.Type() == srcFieldValue.Elem().Type() {
					dstFieldValue.Set(srcFieldValue.Elem())
				}
			}
		}
	}
}
