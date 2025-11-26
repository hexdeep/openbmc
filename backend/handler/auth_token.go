package handler

import (
	"errors"

	"github.com/labstack/echo/v4"
)

func (h *Handler) WithAuthToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		cookie, err := c.Cookie("token")
		if errors.Is(err, echo.ErrCookieNotFound) {
			return c.JSON(401, Res("未认证", nil))
		} else if err != nil {
			return err
		}

		if err := h.AuthToken(cookie.Value, c.Request().Context()); errors.Is(err, ErrTokenNotFound) {
			return c.JSON(401, Res("用户未登录", nil))
		} else if errors.Is(err, ErrTokenExpired) {
			return c.JSON(401, Res("用户凭证已过期", nil))
		} else if errors.Is(err, ErrTokenInvalid) {
			return c.JSON(400, Res("用户凭证不合法", nil))
		} else if err != nil {
			return err
		}

		return next(c)
	}
}
