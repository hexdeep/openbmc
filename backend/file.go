package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"sort"

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

func (h *Handler) ListFile(c echo.Context) error {

	return nil
}

func (h *Handler) UploadFile(c echo.Context) error {

	return nil
}

func (h *Handler) DeleteFile(c echo.Context) error {

	return nil
}
