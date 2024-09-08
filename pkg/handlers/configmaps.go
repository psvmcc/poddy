package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"poddy/pkg/helpers"
	"poddy/pkg/types"

	"github.com/labstack/echo/v4"
)

func ConfigMaps(c echo.Context) error {
	cfg := c.Get("cfg").(types.ConfigFile)
	role := c.Get("role").(string)
	namespace := c.Param("namespace")

	if cfg.Roles[role].NamespacesAccess[namespace] == "" && role != "admin" {
		return c.NoContent(http.StatusForbidden)
	}
	files, err := helpers.GetAllFiles(fmt.Sprintf("%s/%s/configmaps/", cfg.DataPath, namespace))
	if err != nil && files == nil {
		return c.NoContent(http.StatusNotFound)
	}
	result := make([][]string, len(files))
	for i := range files {
		result = append(result, []string{namespace, files[i]})
	}
	return c.JSON(http.StatusOK, result)
}

func ConfigMapGet(c echo.Context) error {
	cfg := c.Get("cfg").(types.ConfigFile)
	role := c.Get("role").(string)
	namespace := c.Param("namespace")
	configmap := c.Param("configmap")
	configmapPath := fmt.Sprintf("%s/%s/configmaps/%s", cfg.DataPath, namespace, configmap)

	if cfg.Roles[role].NamespacesAccess[namespace] == "" && role != "admin" {
		return c.NoContent(http.StatusForbidden)
	}
	sha256, err := helpers.GetSHA256(configmapPath)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	c.Response().Header().Add("Poddy-sha256", sha256)
	return c.File(configmapPath)
}

func ConfigMapDelete(c echo.Context) error {
	cfg := c.Get("cfg").(types.ConfigFile)
	role := c.Get("role").(string)
	namespace := c.Param("namespace")
	configmap := c.Param("configmap")
	configmapPath := fmt.Sprintf("%s/%s/configmaps/%s", cfg.DataPath, namespace, configmap)

	if cfg.Roles[role].NamespacesAccess[namespace] != "rw" && role != "admin" {
		return c.NoContent(http.StatusForbidden)
	}
	if helpers.FileExists(configmapPath) {
		err := os.Remove(configmapPath)
		if err != nil {
			fmt.Printf("Error removing file: %v\n", err)
			return c.NoContent(http.StatusInternalServerError)
		}
		return c.NoContent(http.StatusOK)
	}
	return c.NoContent(http.StatusNotFound)
}

func ConfigMapCreate(c echo.Context) error {
	cfg := c.Get("cfg").(types.ConfigFile)
	role := c.Get("role").(string)
	namespace := c.Param("namespace")
	configmap := c.Param("configmap")
	configmapPath := fmt.Sprintf("%s/%s/configmaps/%s", cfg.DataPath, namespace, configmap)

	if cfg.Roles[role].NamespacesAccess[namespace] != "rw" && role != "admin" {
		return c.NoContent(http.StatusForbidden)
	}
	returnCode := http.StatusCreated
	if helpers.FileExists(configmapPath) {
		returnCode = http.StatusOK
	} else {
		err := os.MkdirAll(fmt.Sprintf("%s/%s/configmaps", cfg.DataPath, namespace), 0o750)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Unable to create directory"})
		}
	}

	dst, err := os.OpenFile(filepath.Clean(configmapPath), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Unable to open file"})
	}
	defer dst.Close()

	if _, err := io.Copy(dst, c.Request().Body); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to write file"})
	}

	return c.NoContent(returnCode)
}
