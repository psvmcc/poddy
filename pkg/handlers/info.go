package handlers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"poddy/pkg/types"

	"github.com/containers/podman/v4/pkg/bindings"
	"github.com/containers/podman/v4/pkg/bindings/system"
	"github.com/labstack/echo/v4"
)

func Info(c echo.Context) error {
	var result types.Info
	result.Commit = types.Commit
	result.Version = types.Version
	socket := "unix:///var/folders/54/ckkdp57x2c39lr2q1pzrstwm0000gn/T/podman/podman-machine-default-api.sock"
	ctx, err := bindings.NewConnection(context.Background(), socket)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	info, err := system.Info(ctx, &system.InfoOptions{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	result.Podman.Version = info.Version
	return c.JSON(http.StatusOK, result)
}
