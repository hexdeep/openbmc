package file

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/hexdeep/openbmc/backend/utils"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
)

type Handler struct {
	Filer Filer
}

func NewHandler(filer Filer) *Handler {
	return &Handler{Filer: filer}
}

func (h *Handler) ListFolder(c echo.Context, r *struct {
	Path string `json:"path"`
}) error {

	entries, err := h.Filer.ListFolder(r.Path)
	if err != nil {
		return fmt.Errorf("failed to list folder: %w", err)
	}

	type Result struct {
		Name    string    `json:"name"`
		Path    string    `json:"path"`
		Size    int64     `json:"size"`
		IsDir   bool      `json:"isDir"`
		ModTime time.Time `json:"modTime"`
	}

	results := lo.Map(entries, func(item fs.DirEntry, i int) Result {
		info, _ := item.Info()
		return Result{
			Name:    item.Name(),
			Path:    filepath.Join(r.Path, item.Name()),
			Size:    info.Size(),
			IsDir:   item.IsDir(),
			ModTime: info.ModTime(),
		}
	})

	sort.Slice(results, func(i, j int) bool {
		if results[i].IsDir != results[j].IsDir {
			return results[i].IsDir
		}
		return results[i].Name < results[j].Name
	})

	return c.JSON(200, utils.Res("", results))
}

func (h *Handler) CreateFolder(c echo.Context, r *struct {
	Path string `json:"path"`
}) error {

	if !IsPathValid(r.Path) {
		return c.JSON(400, utils.Res("非法路径", nil))
	}

	if err := h.Filer.CreateFolder(r.Path); err != nil {
		return err
	}

	return c.JSON(200, utils.Res("文件夹创建成功", true))
}

func Delete(h *Handler, c echo.Context, r *struct {
	Path string `json:"path"`
}) error {

	if !IsPathValid(r.Path) {
		return c.JSON(400, utils.Res("非法路径", nil))
	}

	if err := h.Filer.Delete(r.Path); err != nil {
		return err
	}

	return c.JSON(200, utils.Res("删除成功", true))
}

func (h *Handler) UploadFile(c echo.Context, r *struct {
	Path string `json:"path"`
}) error {

	if !IsPathValid(r.Path) {
		return c.JSON(200, utils.Res("非法路径", nil))
	}

	fileHeader, err := c.FormFile("file")
	if err != nil {
		return err
	}

	if err := h.Filer.UploadFile(r.Path, fileHeader); err != nil {
		return err
	}

	return c.JSON(200, utils.Res("文件上传成功", true))
}

func (h *Handler) Delete(c echo.Context, r *struct {
	Path string `query:"path"`
}) error {

	if !IsPathValid(r.Path) {
		return c.JSON(400, utils.Res("非法路径", nil))
	}

	if err := h.Filer.Delete(r.Path); err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}

	return c.JSON(200, utils.Res("删除成功", nil))
}

func IsPathValid(path string) bool {
	return !strings.Contains(path, "..")
}
