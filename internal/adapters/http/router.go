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

	// Notification routes
	api.Post("/notifications/trigger", a.TriggerNotificationHandler)

	// User management routes
	api.Post("/users", a.CreateUserHandler)
	api.Put("/users/:id", a.UpdateUserHandler)
	api.Delete("/users/:id", a.DeleteUserHandler)
}
