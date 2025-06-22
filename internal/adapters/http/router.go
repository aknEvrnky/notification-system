package http

import (
	"github.com/aknEvrnky/notification-system/pkg/config"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func (a *Adapter) registerRoutes() {
	api := a.fiberApp.Group("/api")
	cfg := config.NewConfig()

	api.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			cfg.BasicAuthUsername: cfg.BasicAuthPassword,
		},
	}))

	// Register your routes here
}
