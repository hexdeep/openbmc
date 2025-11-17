package main

import (
	"sync"

	"gorm.io/gorm"
)

type Config struct {
	Password string
}

type Handler struct {
	Config Config
	Token  Token
	gorm.DB
}

type Token struct {
	Map map[string]struct{}
	Mu  sync.RWMutex
}
