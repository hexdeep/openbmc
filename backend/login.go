package main

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) Login(c echo.Context) error {

	type LoginRequest struct {
		Password string `json:"password"`
	}

	var req LoginRequest
	if c.Bind(&req) != nil {
		return c.JSON(400, Res("请求格式有误", nil))
	}

	h.Config.Mu.RLock()
	password := h.Config.Password
	h.Config.Mu.RUnlock()

	if err := bcrypt.CompareHashAndPassword(
		[]byte(password), []byte(req.Password),
	); errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return c.JSON(401, Res("密码不正确", nil))
	} else if err != nil {
		return err
	}

	token, err := h.NewToken(c.Request().Context())
	if err != nil {
		return err
	}

	c.SetCookie(&http.Cookie{
		Name:     "token",
		Value:    token,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   24 * 60 * 60,
		Domain:   "axogc.net",
		Path:     "/",
	})

	return c.JSON(200, Res("登录成功", nil))
}
