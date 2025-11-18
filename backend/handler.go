package main

import (
	"context"
	"errors"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type Handler interface {
	GetPassword() string
	NewToken(context.Context) (string, error)
	AuthToken(string, context.Context) error
}

type Context struct {
	Config *Config
	Logs   chan *Log
	*gorm.DB
}

func (c *Context) GetPassword() string {

	return ""
}

var (
	ErrTokenInvalid  = errors.New("token invalid")
	ErrTokenNotFound = errors.New("token not found")
	ErrTokenExpired  = errors.New("token expired")
)

func (c *Context) AuthToken(value string, ctx context.Context) error {

	id, err := strconv.Atoi(value)
	if err != nil {
		return ErrTokenInvalid
	}

	token, err := gorm.G[Token](c.DB).Where("id = ?", id).Take(ctx)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrTokenNotFound
	}

	if time.Now().After(token.ExpiresAt) {
		return ErrTokenExpired
	}

	return nil
}

func (c *Context) NewToken(ctx context.Context) (string, error) {

	return "", nil
}
