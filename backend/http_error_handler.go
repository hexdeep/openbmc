package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func HTTPErrorHandler(err error, c echo.Context) {
	fmt.Printf("error encountered: %v\n", err)
	c.JSON(500, Res("服务器内部错误", nil))
}
