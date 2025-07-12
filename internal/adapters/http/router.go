package http

import (
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func (a *Adapter) registerRoutes(authUser string, authPassword string) {
	api := a.fiberApp.Group("/api")

	api.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			authUser: authPassword,
		},
	}))

	// Register your routes here
	api.Post("/notifications/trigger", a.TriggerNotificationHandler)
}
