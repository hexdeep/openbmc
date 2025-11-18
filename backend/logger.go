package main

import (
	"context"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (ctx *Context) LoggerMiddleWare(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		err := next(c)

		ctx.Logs <- &Log{
			Method: c.Request().Method,
			Path:   c.Request().URL.Path,
			Status: c.Response().Status,
		}

		return err
	}
}

func (c *Context) Log() {
	for l := range c.Logs {
		ctx, canc := context.WithTimeout(context.Background(), 5*time.Second)
		defer canc()
		if err := gorm.G[Log](c.DB).Create(ctx, l); err != nil {
			log.Printf("failed to store request log: %v\n", err)
		}
	}
}
