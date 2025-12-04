package log

import (
	"context"
	"time"
)

type Logger interface {
	Create(ctx context.Context, log *Log) error
	Find(ctx context.Context, filter *LogFilter) ([]Log, int64, error)
	Delete(ctx context.Context, filter *LogFilter) (int, error)
}

type Log struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	ClientIP  string    `json:"clientIp"`
	Method    string    `json:"method"`
	Path      string    `json:"path"`
	Status    int       `json:"status"`
}

type LogFilter struct {
	From     time.Time `query:"from"`
	To       time.Time `query:"to"`
	ClientIP string    `query:"clientIp"`
	Method   string    `query:"method"`
	Path     string    `query:"path"`
	Status   int       `query:"status"`
	Page     int       `query:"page"`
	Size     int       `query:"size"`
}
