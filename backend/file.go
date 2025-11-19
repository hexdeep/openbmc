package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
)

func (h *Handler) GetFolder(c echo.Context) error {

	path := c.QueryParam("path")
	fullpath := filepath.Join(h.Config.FilePath, path)

	entries, err := os.ReadDir(fullpath)
	if err != nil {
		return err
	}

	if strings.Contains(path, "..") {
		return c.JSON(400, Res("禁止遍历上级目录", nil))
	}

	type Result struct {
		Name  string `json:"name"`
		Path  string `json:"path"`
		Size  int64  `json:"size"`
		IsDir bool   `json:"isDir"`
	}

	results := lo.Map(entries, func(item fs.DirEntry, i int) Result {
		info, _ := item.Info()
		return Result{
			Name:  item.Name(),
			Path:  filepath.Join(path, item.Name()),
			Size:  info.Size(),
			IsDir: item.IsDir(),
		}
	})

	sort.Slice(results, func(i, j int) bool {
		if results[i].IsDir != results[j].IsDir {
			return results[i].IsDir
		}
		return results[i].Name < results[j].Name
	})

	return c.JSON(200, Res("", results))
}

func CreateFolder(h *Handler, c echo.Context, r *struct {
	Name string `json:"name"`
}) error {

	if err := os.Mkdir(filepath.Join(h.Config.FilePath, r.Name), 0755); err != nil {
		return err
	}

	return c.JSON(200, Res("文件夹创建成功", true))
}

func DeleteFolder(h *Handler, c echo.Context, r *struct {
	Path string `json:"path"`
}) error {

	if err := os.RemoveAll(filepath.Join(h.Config.FilePath, r.Path)); err != nil {
		return err
	}

	return c.JSON(200, Res("文件夹删除成功", true))
}

func (h *Handler) ListFile(c echo.Context) error {

	return nil
}

func (h *Handler) UploadFile(c echo.Context) error {

	return nil
}

func (h *Handler) DeleteFile(c echo.Context) error {

	return nil
}
