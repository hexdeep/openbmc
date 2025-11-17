package main

type Resp struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Res(message string, data any) *Resp {
	return &Resp{
		Message: message,
		Data:    data,
	}
}
