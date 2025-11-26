package handler

import (
	"context"
	"log"
	"time"

	"gorm.io/gorm"
)

func (h *Handler) ClearData(interval time.Duration) {

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			ctx, canc := context.WithTimeout(context.Background(), 5*time.Second)
			if _, err := gorm.G[Token](h.DB).Where("expires_at <= ?", time.Now()).Delete(ctx); err != nil {
				log.Printf("failed to clear expired token: %v\n", err)
			}
			if _, err := gorm.G[Log](h.DB).Where(
				"created_at <= ?",
				time.Now().Add(time.Duration(h.Config.LogDuration)*(-24)*time.Hour),
			).Delete(ctx); err != nil {
				log.Printf("failed to clear old log: %v\n", err)
			}
			canc()
		}
	}
}
