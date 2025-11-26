package main

import (
	"log"
	"time"

	"github.com/glebarez/sqlite"
	h "github.com/hexdeep/openbmc/backend/handler"
	"github.com/kelseyhightower/envconfig"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {

	config := h.Config{
		Address:         ":8080",
		DBFile:          "/data/data.db",
		CleanerInterval: 600,
		TokenDuration:   7 * 24 * 60,
		DefaultSize:     10,
		LogDuration:     7,
		LogLevel:        h.LogInfo,
		FilePath:        "/data/file",
	}
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

	if err := db.AutoMigrate(h.Tables...); err != nil {
		log.Fatalf("failed to auto migrate database: %v\n", err)
	}

	handler := &h.Handler{
		Config:    &config,
		Logs:      make(chan *h.Log, 100),
		Paginator: h.NewPaginator(config.DefaultSize),
		DB:        db,
	}

	go handler.Log()
	go handler.ClearData(time.Duration(config.CleanerInterval) * time.Second)

	router := h.GetRouter(handler)

	if config.SSL.Enabled {
		err = router.StartTLS(config.Address, config.SSL.Cert, config.SSL.Key)
	} else {
		err = router.Start(config.Address)
	}

	if err != nil {
		log.Fatalf("failed to start http server: %v\n", err)
	}
}
