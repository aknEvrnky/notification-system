package http

import (
	"fmt"
	"github.com/aknEvrnky/notification-system/internal/ports"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"time"
)

type Adapter struct {
	api      ports.ApiPort
	httpPort int
	fiberApp *fiber.App
}

func NewAdapter(api ports.ApiPort, httpPort int) *Adapter {
	fiberApp := newFiberApp()

	a := &Adapter{
		api:      api,
		httpPort: httpPort,
		fiberApp: fiberApp,
	}

	a.registerRoutes()

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
