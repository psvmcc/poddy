package commands

import (
	"fmt"
	"log"
	"net/http"
	"poddy/pkg/handlers"
	"poddy/pkg/logging"
	"poddy/pkg/types"
	"poddy/pkg/victoriametrics"
	"strconv"
	"strings"
	"time"

	"github.com/VictoriaMetrics/metrics"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

var (
	cfg   types.ConfigFile
	creds types.AuthFile
)

func StartServer(c *cli.Context) error {
	metrics.GetOrCreateCounter(fmt.Sprintf("poddy_app_version{version=%q,commit=%q}", types.Version, types.Commit)).Inc()

	cfg.Load(c.String("config"))
	// fmt.Printf("%+v\n", cfg)
	creds.Load(cfg.AuthFile)
	// fmt.Printf("%+v\n", creds)

	logger := logging.Build(c.Bool("verbose"))
	zap.ReplaceGlobals(logger)

	echo.NotFoundHandler = func(c echo.Context) error {
		return c.String(http.StatusNotFound, "404 page not found")
	}
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	httpLogger := zap.S().Named("http")

	e.Use(middleware.RequestID())
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ce echo.Context) error {
			ce.Set("user", "unauthorized")
			ce.Set("role", "none")
			authHeader := ce.Request().Header.Get("Authorization")
			if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
				token := strings.TrimPrefix(authHeader, "Bearer ")
				for i := range creds {
					if creds[i].Token == token {
						ce.Set("user", creds[i].User)
						ce.Set("role", creds[i].Role)
					}
				}
			}
			ce.Set("cfg", cfg)
			ce.Set("logger", httpLogger)
			ce.Response().Header().Set("Server", fmt.Sprintf("poddy/%s (%s)", types.Version, types.Commit))
			return next(ce)
		}
	})

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			req := c.Request()
			res := c.Response()
			start := time.Now()
			if err = next(c); err != nil {
				c.Error(err)
			}
			stop := time.Now()

			if req.RequestURI == "/health" {
				return
			}

			metrics.GetOrCreateCounter(fmt.Sprintf("poddy_requests{status=%q,user=%q,ip=%q}", strconv.Itoa(res.Status), c.Get("user"), c.RealIP())).Inc()
			message := fmt.Sprintf(
				"%s %s requested from %s @%s with status %d in %s",
				req.Method,
				req.RequestURI,
				c.RealIP(),
				c.Get("user"),
				res.Status,
				stop.Sub(start).String(),
			)

			logger := c.Get("logger").(*zap.SugaredLogger)

			if res.Status >= 100 && res.Status <= 399 {
				logger.Named("req").Info(message)
			} else if res.Status >= 400 && res.Status <= 499 {
				logger.Named("req").Warn(message)
			} else {
				logger.Named("req").Error(message)
			}

			return
		}
	})

	e.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		Skipper: skipperFn([]string{"/health"}),
		Validator: func(key string, _ echo.Context) (bool, error) {
			for i := range creds {
				if creds[i].Token == key {
					return true, nil
				}
			}
			return false, nil
		},
	}))

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	e.GET("/api/v1/info", handlers.Info)

	e.GET("/api/v1/namespaces", handlers.Namespaces)
	e.GET("/api/v1/namespaces/:namespace", handlers.NamespaceGet)
	e.HEAD("/api/v1/namespaces/:namespace", handlers.NamespaceGet)

	e.GET("/api/v1/namespaces/:namespace/configmaps", handlers.ConfigMaps)
	e.GET("/api/v1/namespaces/:namespace/configmaps/:configmap", handlers.ConfigMapGet)
	e.HEAD("/api/v1/namespaces/:namespace/configmaps/:configmap", handlers.ConfigMapGet)
	e.POST("/api/v1/namespaces/:namespace/configmaps/:configmap", handlers.ConfigMapCreate)
	e.DELETE("/api/v1/namespaces/:namespace/configmaps/:configmap", handlers.ConfigMapDelete)

	e.GET("/api/v1/namespaces/:namespace/volumes", handlers.Volumes)
	e.GET("/api/v1/namespaces/:namespace/volumes/:volume", handlers.VolumeGet)
	e.HEAD("/api/v1/namespaces/:namespace/volumes/:volume", handlers.VolumeGet)
	e.POST("/api/v1/namespaces/:namespace/volumes/:volume", handlers.VolumeCreate)
	e.DELETE("/api/v1/namespaces/:namespace/volumes/:volume", handlers.VolumeDelete)

	go func() {
		log.Fatal(e.Start(c.String("bind")))
	}()
	return victoriametrics.ListenMetricsServer(c.String("self-exporter-bind"))
}

func skipperFn(skipURLs []string) func(echo.Context) bool {
	return func(context echo.Context) bool {
		for _, url := range skipURLs {
			if url == context.Request().URL.Path {
				return true
			}
		}
		return false
	}
}
