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

	eventType, ok := payload["event_type"]
	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, "event_type is required")
	}

	eventTypeStr, ok := eventType.(string)

	if eventType == "" {
		return fiber.NewError(fiber.StatusBadRequest, "event_type is required")
	}

	ctx.Send([]byte("Notification triggered successfully"))

	return a.api.TriggerNotification(ctx.Context(), eventTypeStr, payload)
}
