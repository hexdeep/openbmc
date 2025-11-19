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

type List struct {
	Data  any   `json:"data"`
	Total int64 `json:"total"`
}

func NewList(data any, total int64) *List {
	return &List{
		Data:  data,
		Total: total,
	}
}
