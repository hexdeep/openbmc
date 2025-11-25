package main

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ListLogRequest struct {
	From     *time.Time `query:"from"`
	To       *time.Time `query:"to"`
	ClientIP string     `query:"clientIp"`
	Method   string     `query:"method"`
	Path     string     `query:"path"`
	Status   int        `query:"status"`
	Page     int        `query:"page"`
	Size     int        `query:"size"`
}

func ListLog(h *Handler, c echo.Context, r *ListLogRequest) error {

	type Log struct {
		ID        uint      `json:"id"`
		CreatedAt time.Time `json:"createdAt"`
		ClientIP  string    `json:"clientIp"`
		Method    string    `json:"method"`
		Path      string    `json:"path"`
		Status    int       `json:"status"`
	}

	q := gorm.G[Log](h.DB).Scopes()

	if r.From != nil {
		q = q.Where("created_at >= ?", r.From)
	}

	if r.To != nil {
		q = q.Where("created_at <= ?", r.To)
	}

	if r.ClientIP != "" {
		q = q.Where("client_ip = ?", r.ClientIP)
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

	total, err := q.Count(c.Request().Context(), "*")
	if err != nil {
		return err
	}

	logs, err := q.Scopes(h.Paginate(r.Page, r.Size)).Find(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(200, Res("", NewList(logs, total)))
}

func (h *Handler) DeleteLoga(c echo.Context) error {

	return nil
}

func ClearLog(h *Handler, c echo.Context, r *struct {
	Method []string `json:"method"`
	Status []int    `json:"status"`
}) error {

	q := gorm.G[Log](h.DB).Scopes()

	if len(r.Method) != 0 {
		q = q.Where("method IN ?", r.Method)
	}

	if len(r.Status) != 0 {
		q = q.Where("status IN ?", r.Status)
	}

	row, err := q.Delete(c.Request().Context())
	if err != nil {
		return err
	} else if row == 0 {
		return c.JSON(200, Res("没有匹配的数据", nil))
	}

	return c.JSON(200, Res(fmt.Sprintf("成功删除了 %d 条数据", row), nil))
}
