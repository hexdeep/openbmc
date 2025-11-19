package main

import (
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ListLogRequest struct {
	Time   *[2]time.Time `form:"time"`
	Method string        `form:"method"`
	Path   string        `form:"path"`
	Status int           `form:"status"`
	Page   int           `form:"page"`
	Size   int           `form:"size"`
}

func ListLog(h *Handler, c echo.Context, r *ListLogRequest) error {

	type Log struct {
		ID        uint      `json:"id"`
		CreatedAt time.Time `json:"createdAt"`
		Method    string    `json:"method"`
		Path      string    `json:"path"`
		Status    int       `json:"status"`
	}

	q := gorm.G[Log](h.DB).Scopes(h.Paginate(r.Page, r.Size))

	if r.Time != nil {
		q = q.Where("created_at BETWEEN ? AND ?", r.Time[0], r.Time[1])
	}

	if r.Method != "" {
		q = q.Where("method = ?", r.Method)
	}

	if r.Path != "" {
		q = q.Where("path = ?", r.Path)
	}

	if r.Status != 0 {
		q = q.Where("status = ?", r.Status)
	}

	logs, err := q.Find(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(200, Res("", logs))
}
