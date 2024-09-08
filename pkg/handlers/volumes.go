package handlers

import (
	"fmt"
	"net/http"
	"os"
	"poddy/pkg/helpers"
	"poddy/pkg/types"

	"github.com/labstack/echo/v4"
)

func Volumes(c echo.Context) error {
	cfg := c.Get("cfg").(types.ConfigFile)
	role := c.Get("role").(string)
	namespace := c.Param("namespace")

	if cfg.Roles[role].NamespacesAccess[namespace] == "" && role != "admin" {
		return c.NoContent(http.StatusForbidden)
	}
	files, err := helpers.GetAllFiles(fmt.Sprintf("%s/%s/volumes/", cfg.DataPath, namespace))
	if err != nil && files == nil {
		return c.NoContent(http.StatusNotFound)
	}
	var result [][]string
	for i := range files {
		result = append(result, []string{namespace, files[i]})
	}
	return c.JSON(http.StatusOK, result)
}

func VolumeGet(c echo.Context) error {
	cfg := c.Get("cfg").(types.ConfigFile)
	role := c.Get("role").(string)
	namespace := c.Param("namespace")
	volume := c.Param("volume")
	volumePath := fmt.Sprintf("%s/%s/volumes/%s", cfg.DataPath, namespace, volume)

	if cfg.Roles[role].NamespacesAccess[namespace] == "" && role != "admin" {
		return c.NoContent(http.StatusForbidden)
	}
	if helpers.FileExists(volumePath) {
		return c.NoContent(http.StatusOK)
	}
	return c.NoContent(http.StatusNotFound)
}

func VolumeDelete(c echo.Context) error {
	cfg := c.Get("cfg").(types.ConfigFile)
	role := c.Get("role").(string)
	namespace := c.Param("namespace")
	volume := c.Param("volume")
	volumePath := fmt.Sprintf("%s/%s/volumes/%s", cfg.DataPath, namespace, volume)

	if cfg.Roles[role].NamespacesAccess[namespace] != "rw" && role != "admin" {
		return c.NoContent(http.StatusForbidden)
	}
	if helpers.FileExists(volumePath) {
		err := os.RemoveAll(volumePath)
		if err != nil {
			fmt.Printf("Error removing file: %v\n", err)
			return c.NoContent(http.StatusInternalServerError)
		}
		return c.NoContent(http.StatusOK)
	}
	return c.NoContent(http.StatusNotFound)
}

func VolumeCreate(c echo.Context) error {
	cfg := c.Get("cfg").(types.ConfigFile)
	role := c.Get("role").(string)
	namespace := c.Param("namespace")
	volume := c.Param("volume")
	volumePath := fmt.Sprintf("%s/%s/volumes/%s", cfg.DataPath, namespace, volume)

	if cfg.Roles[role].NamespacesAccess[namespace] != "rw" && role != "admin" {
		return c.NoContent(http.StatusForbidden)
	}
	returnCode := http.StatusCreated
	if helpers.FileExists(volumePath) {
		returnCode = http.StatusOK
	} else {
		err := os.MkdirAll(volumePath, 0755)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Unable to create directory"})
		}
	}

	return c.NoContent(returnCode)
}
