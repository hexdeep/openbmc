package main

import (
	"embed"
	"io/fs"
	"mime"
	"net/http"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

// 将 dist 内容 embed 进二进制
//
//go:embed frontend/*
var embeddedFrontend embed.FS

// 清理路径并返回内容
func HandleFrontend(c echo.Context) error {
	reqPath := c.Param("*")
	if reqPath == "" {
		reqPath = "index.html"
	}

	// 强制路径不能以 "/" 开头
	reqPath = filepath.Clean(reqPath)
	if reqPath == "." {
		reqPath = "index.html"
	}

	// 尝试打开文件
	fsys, _ := fs.Sub(embeddedFrontend, "frontend")

	data, err := fs.ReadFile(fsys, reqPath)
	if err != nil {
		// 任何文件找不到都 fallback 到 index.html
		data, _ = fs.ReadFile(fsys, "index.html")
		c.Response().Header().Set("Content-Type", "text/html; charset=utf-8")
		return c.Blob(http.StatusOK, "text/html; charset=utf-8", data)
	}

	// 自动根据文件扩展名设置 MIME
	ext := filepath.Ext(reqPath)
	contentType := mime.TypeByExtension(ext)
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	return c.Blob(http.StatusOK, contentType, data)
}
