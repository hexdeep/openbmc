package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func WithSSE(h *Handler, fn func(h *Handler, c echo.Context, send chan<- string) error) func(c echo.Context) error {
	return func(c echo.Context) error {

		c.Response().Header().Set(echo.HeaderContentType, "text/event-stream")
		c.Response().Header().Set("Cache-Control", "no-cache")
		c.Response().Header().Set("Connection", "keep-alive")
		c.Response().WriteHeader(http.StatusOK)

		send := make(chan string)

		go func() {
			for msg := range send {
				_, err := c.Response().Write([]byte("data: " + msg + "\n\n"))
				if err != nil {
					close(send)
					return
				}
				c.Response().Flush()
			}
		}()

		return fn(h, c, send)
	}
}
