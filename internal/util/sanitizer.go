package util

import (
	"reflect"
	"strings"
)

// TrimStrings akan memeriksa semua field bertipe string dalam struct dan melakukan TrimSpace
func TrimStrings(v interface{}) {
	val := reflect.ValueOf(v)

	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		return
	}

	val = val.Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		if field.Kind() == reflect.String && field.CanSet() {
			trimmed := strings.TrimSpace(field.String())
			field.SetString(trimmed)
		}
	}
}
