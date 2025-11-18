package main

import (
	"errors"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (h *Handler) AuthToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		cookie, err := c.Cookie("token")
		if errors.Is(err, echo.ErrCookieNotFound) {
			return c.JSON(401, Res("未认证", nil))
		} else if err != nil {
			c.Error(err)
			return c.JSON(500, Res("读取用户凭证失败", nil))
		}

		session, err := gorm.G[Session](h.DB).Where("id = ?", cookie.Value).Take(c.Request().Context())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(401, Res("用户未登录", nil))
		} else if err != nil {
			c.Error(err)
			return c.JSON(500, Res("读取用户凭证失败", nil))
		}

		if time.Now().After(session.ExpiresAt) {
			return c.JSON(401, Res("用户凭证已过期", nil))
		}

		return next(c)
	}
}
