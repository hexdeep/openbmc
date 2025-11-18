package main

import (
	"context"
	"log"
	"time"

	"gorm.io/gorm"
)

func (h *Handler) ClearToken(interval time.Duration) {

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			ctx, canc := context.WithTimeout(context.Background(), 5*time.Second)
			if _, err := gorm.G[Token](h.DB).Where("expires_at <= ?", time.Now()).Delete(ctx); err != nil {
				log.Printf("failed to clear expired token: %v\n", err)
			}
			canc()
		}
	}
}
