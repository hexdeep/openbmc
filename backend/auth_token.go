package main

import (
	"errors"

	"github.com/labstack/echo/v4"
)

func (h *Handler) AuthToken(c echo.Context) error {

	cookie, err := c.Cookie("token")
	if errors.Is(err, echo.ErrCookieNotFound) {
		return c.JSON(401, Res("未认证", nil))
	} else if err != nil {
		c.Error(err)
		return c.JSON(500, Res("读取用户凭证失败", nil))
	}

	h.Token.Mu.RLock()
	_, exists := h.Token.Map[cookie.Value]
	h.Token.Mu.RUnlock()
	if !exists {
		return c.JSON(401, Res("未登录", nil))
	}

	return nil
}
