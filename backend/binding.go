package main

import "github.com/labstack/echo/v4"

func WithBind[R any](h *Handler, fn func(h *Handler, c echo.Context, r *R) error) echo.HandlerFunc {
	return func(c echo.Context) error {
		var r R
		if c.Bind(&r) != nil {
			return c.JSON(400, Res("非法请求", nil))
		}
		return fn(h, c, &r)
	}
}
