package main

import (
	"gorm.io/gorm"
)

type Config struct {
	Port     string
	Password string
	DBFile   string
}

type Handler struct {
	Config *Config
	Logs   chan *Log
	*gorm.DB
}
