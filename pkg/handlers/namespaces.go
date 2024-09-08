package handlers

import (
	"net/http"
	"poddy/pkg/helpers"
	"poddy/pkg/types"
	"sort"

	"github.com/labstack/echo/v4"
)

func Namespaces(c echo.Context) error {
	cfg := c.Get("cfg").(types.ConfigFile)
	role := c.Get("role").(string)

	var result [][]string
	if role == "admin" {
		for i := range cfg.Namespaces {
			ns := []string{i, "rw", cfg.Namespaces[i].Network}
			result = append(result, ns)
		}
	} else {
		for i := range cfg.Roles[role].NamespacesAccess {
			ns := []string{i, cfg.Roles[role].NamespacesAccess[i], cfg.Namespaces[i].Network}
			result = append(result, ns)
		}
	}

	sort.Sort(helpers.ByKey(result))
	return c.JSON(http.StatusOK, result)
}

func NamespaceGet(c echo.Context) error {
	cfg := c.Get("cfg").(types.ConfigFile)
	role := c.Get("role").(string)
	namespace := c.Param("namespace")

	if cfg.Roles[role].NamespacesAccess[namespace] == "" && role != "admin" {
		return c.NoContent(http.StatusForbidden)
	}
	if _, ok := cfg.Namespaces[namespace]; !ok {
		return c.NoContent(http.StatusNotFound)
	}
	var ns types.Namespace
	ns.Name = namespace
	if role == "admin" {
		ns.AccessType = "rw"
	} else {
		ns.AccessType = cfg.Roles[role].NamespacesAccess[namespace]
	}
	ns.Network = cfg.Namespaces[namespace].Network
	ns.ENV = cfg.Namespaces[namespace].ENV
	return c.JSON(http.StatusOK, ns)
}
