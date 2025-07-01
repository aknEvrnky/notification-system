package http

import (
	"fmt"
	"github.com/aknEvrnky/notification-system/internal/ports"
	"github.com/aknEvrnky/notification-system/pkg/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"time"
)

type Adapter struct {
	api      ports.ApiPort
	httpPort int
	fiberApp *fiber.App
}

func NewAdapter(api ports.ApiPort, cfg *config.Config) *Adapter {
	fiberApp := newFiberApp()

	a := &Adapter{
		api:      api,
		httpPort: cfg.ApplicationPort,
		fiberApp: fiberApp,
	}

	a.registerRoutes(cfg.BasicAuthUsername, cfg.BasicAuthPassword)

	return a
}

func newFiberApp() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())

	return app
}

func (a *Adapter) Run() error {
	return a.fiberApp.Listen(fmt.Sprintf(":%d", a.httpPort))
}

func (a *Adapter) Shutdown() error {
	return a.fiberApp.ShutdownWithTimeout(5 * time.Second)
}
