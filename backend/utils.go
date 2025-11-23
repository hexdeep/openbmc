package main

import (
	"io"
	"os"
)

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

func Echo(filename, data string) (string, error) {

	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = file.Write([]byte(data))
	if err != nil {
		return "", err
	}

	value, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(value), nil
}
