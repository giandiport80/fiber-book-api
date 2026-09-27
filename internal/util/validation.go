package util

import (
	"fmt"

	"github.com/go-playground/validator"
)

func Validate[T any](data T) map[string]string {
	err := validator.New().Struct(data)
	if err != nil {
		res := map[string]string{}
		for _, v := range err.(validator.ValidationErrors) {
			res[v.StructField()] = TranslateTag(v)
		}
		return res
	}
	return nil
}

func TranslateTag(fd validator.FieldError) string {
	switch fd.ActualTag() {
	case "required":
		return fmt.Sprintf("field %s wajib diisi", fd.StructField())
	case "email":
		return fmt.Sprintf("field %s harus berupa format email yang valid", fd.StructField())
	case "min":
		if fd.Kind().String() == "string" {
			return fmt.Sprintf("field %s minimal berisi %s karakter", fd.StructField(), fd.Param())
		}
		return fmt.Sprintf("field %s minimal bernilai %s", fd.StructField(), fd.Param())
	case "max":
		if fd.Kind().String() == "string" {
			return fmt.Sprintf("field %s maksimal berisi %s karakter", fd.StructField(), fd.Param())
		}
		return fmt.Sprintf("field %s maksimal bernilai %s", fd.StructField(), fd.Param())
	case "numeric":
		return fmt.Sprintf("field %s harus berupa angka", fd.StructField())
	case "alphanum":
		return fmt.Sprintf("field %s hanya boleh berisi huruf dan angka", fd.StructField())
	case "eqfield":
		return fmt.Sprintf("field %s harus sama dengan field %s", fd.StructField(), fd.Param())
	case "nefield":
		return fmt.Sprintf("field %s tidak boleh sama dengan field %s", fd.StructField(), fd.Param())
	case "gt":
		return fmt.Sprintf("field %s harus lebih besar dari %s", fd.StructField(), fd.Param())
	case "gte":
		return fmt.Sprintf("field %s harus lebih besar dari atau sama dengan %s", fd.StructField(), fd.Param())
	case "lt":
		return fmt.Sprintf("field %s harus lebih kecil dari %s", fd.StructField(), fd.Param())
	case "lte":
		return fmt.Sprintf("field %s harus lebih kecil dari atau sama dengan %s", fd.StructField(), fd.Param())
	case "oneof":
		return fmt.Sprintf("field %s harus bernilai salah satu dari [%s]", fd.StructField(), fd.Param())
	case "url":
		return fmt.Sprintf("field %s harus berupa format URL yang valid", fd.StructField())
	}

	return fmt.Sprintf("field %s tidak valid", fd.StructField())
}
