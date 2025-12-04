package utils

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Res(message string, data any) *Response {
	return &Response{
		Message: message,
		Data:    data,
	}
}

type List struct {
	Data  any   `json:"data"`
	Total int64 `json:"total"`
}

func ListRes(message string, total int64, data any) *Response {
	return &Response{
		Message: message,
		Data: &List{
			Total: total,
			Data:  data,
		},
	}
}
