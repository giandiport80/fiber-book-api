package dto

type Response[T any] struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

func CreateResponseError(message string) Response[string] {
	return Response[string]{
		Code:    "99",
		Message: message,
	}
}

func CreateResponseSuccess[T any](data T) Response[T] {
	return Response[T]{
		Code:    "00",
		Message: "success",
		Data:    data,
	}
}
