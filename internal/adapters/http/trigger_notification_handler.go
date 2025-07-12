package http

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func (a *Adapter) TriggerNotificationHandler(ctx *fiber.Ctx) error {

	payload := make(map[string]any)

	if err := json.Unmarshal(ctx.Body(), &payload); err != nil {
		return err
	}

	eventType := payload["event_type"].(string)

	if eventType == "" {
		return fiber.NewError(fiber.StatusBadRequest, "event_type is required")
	}

	return a.api.TriggerNotification(ctx.Context(), eventType, payload)
}
