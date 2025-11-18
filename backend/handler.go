package main

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type Handler struct {
	Config *Config
	Logs   chan *Log
	*gorm.DB
}

func (h *Handler) GetPassword() string {

	return ""
}

var (
	ErrTokenNotFound = errors.New("token not found")
	ErrTokenExpired  = errors.New("token expired")
)

func (h *Handler) AuthToken(token string, ctx context.Context) error {

	return nil
}

func (h *Handler) SetToken(token string) error {

	return nil
}
