package handlers

import (
	"net/http"
	"poddy/pkg/types"

	"github.com/labstack/echo/v4"
)

func Info(c echo.Context) error {
	var result types.Info
	result.Commit = types.Commit
	result.Version = types.Version
	return c.JSON(http.StatusOK, result)
}
