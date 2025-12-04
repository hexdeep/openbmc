package log

import (
	"fmt"

	"github.com/hexdeep/openbmc/backend/utils"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	Logger Logger
}

func NewHandler(logger Logger) *Handler {
	return &Handler{Logger: logger}
}

func (h *Handler) List(c echo.Context, filter *LogFilter) error {

	logs, total, err := h.Logger.Find(c.Request().Context(), filter)
	if err != nil {
		return err
	} else if total == 0 {
		return c.JSON(400, utils.Res("找不到符合条件的数据", nil))
	}

	return c.JSON(200, utils.ListRes("", total, logs))
}

func (h *Handler) Delete(c echo.Context, filter *LogFilter) error {

	row, err := h.Logger.Delete(c.Request().Context(), filter)
	if err != nil {
		return err
	} else if row == 0 {
		return c.JSON(400, utils.Res("找不到符合条件的数据", false))
	}

	return c.JSON(200, utils.Res(fmt.Sprintf("成功删除了 %v 条数据", row), true))
}
