package handler

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/hexdeep/openbmc/backend/proc"
	"gorm.io/gorm"
)

type Handler struct {
	Proc   proc.Proc
	Config *Config
	Logs   chan *Log
	*Paginator
	*gorm.DB
}

func NewHandler(config *Config, logs chan *Log, paginator *Paginator, db *gorm.DB) *Handler {
	return &Handler{
		Config:    config,
		Logs:      logs,
		Paginator: paginator,
		DB:        db,
	}
}

var (
	ErrTokenInvalid  = errors.New("token invalid")
	ErrTokenNotFound = errors.New("token not found")
	ErrTokenExpired  = errors.New("token expired")
)

func (h *Handler) AuthToken(value string, ctx context.Context) error {

	id, err := strconv.Atoi(value)
	if err != nil {
		return ErrTokenInvalid
	}

	token, err := gorm.G[Token](h.DB).Where("id = ?", id).Take(ctx)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrTokenNotFound
	} else if err != nil {
		return err
	}

	if time.Now().After(token.ExpiresAt) {
		return ErrTokenExpired
	}

	return nil
}

func (h *Handler) NewToken(ctx context.Context) (string, error) {

	token := Token{ExpiresAt: time.Now().Add(time.Duration(h.Config.TokenDuration) * time.Second)}

	if err := gorm.G[Token](h.DB, gorm.WithResult()).Create(ctx, &token); err != nil {
		return "", err
	}

	return strconv.Itoa(int(token.ID)), nil
}
