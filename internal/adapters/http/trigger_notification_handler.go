package http

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func (a *Adapter) TriggerNotificationHandler(ctx *fiber.Ctx) error {

	payload := make(map[string]any)

	if err := json.Unmarshal(ctx.Body(), &payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid payload")
	}

	eventRaw, ok := payload["event_type"]
	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, "event_type is required")
	}

	eventType, ok := eventRaw.(string)
	if !ok || eventType == "" {
		return fiber.NewError(fiber.StatusBadRequest, "event_type must be a string")
	}

	if err := a.api.TriggerNotification(ctx.Context(), eventType, payload); err != nil {
		return err
	}

	return ctx.Send([]byte("Notification triggered successfully"))
}
