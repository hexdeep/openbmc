package main

import (
	"log"

	"github.com/kelseyhightower/envconfig"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {

	var config Config
	if err := envconfig.Process("", &config); err != nil {
		log.Fatalf("failed to load config: %v\n", err)
	}

	db, err := gorm.Open(sqlite.Open(config.DBFile), &gorm.Config{
		TranslateError: true,
		Logger:         logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("failed to connect to database: %v\n", err)
	}

	if err := db.AutoMigrate(tables...); err != nil {
		log.Fatalf("failed to auto migrate database: %v\n", err)
	}

	handler := &Handler{
		Config: &config,
		DB:     db,
	}

	go handler.Logger()

	if err := GetRouter(handler).Start(config.Port); err != nil {
		log.Fatalf("failed to start http server: %v\n", err)
	}
}
