package dto

type Response[T any] struct {
	Code    string            `json:"code"`
	Message string            `json:"message"`
	Data    T                 `json:"data,omitempty"`
	Errors  map[string]string `json:"errors,omitempty"`
}

func CreateResponseError(message string, errors map[string]string) Response[string] {
	return Response[string]{
		Code:    "99",
		Message: message,
		Errors:  errors,
	}
}

func CreateResponseSuccess[T any](data T) Response[T] {
	return Response[T]{
		Code:    "00",
		Message: "success",
		Data:    data,
	}
}
