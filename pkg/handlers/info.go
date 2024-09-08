package handlers

import (
	"net/http"
	"poddy/pkg/types"

	"github.com/labstack/echo/v4"
)

func Info(c echo.Context) error {
	// cfg := c.Get("cfg").(types.ConfigFile)
	// role := c.Get("role").(string)
	// logger := c.Get("logger").(*zap.SugaredLogger)

	var result types.Info
	result.Commit = types.Commit
	result.Version = types.Version
	return c.JSON(http.StatusOK, result)
}
