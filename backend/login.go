package main

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) Login(c echo.Context) error {

	type LoginRequest struct {
		Password string `json:"password"`
	}

	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		c.Error(err)
		return c.JSON(400, Res("请求格式有误", nil))
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(h.GetPassword()), []byte(req.Password),
	); errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return c.JSON(401, Res("密码不正确", nil))
	} else if err != nil {
		c.Error(err)
		return c.JSON(500, Res("密码校验失败", nil))
	}

	token := uuid.New().String()
	if err := h.SetToken(token); err != nil {
		c.Error(err)
		return c.JSON(500, Res("存储用户凭证失败", nil))
	}

	c.SetCookie(&http.Cookie{
		Name:     "token",
		Value:    token,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})

	return c.JSON(200, Res("登录成功", nil))
}
