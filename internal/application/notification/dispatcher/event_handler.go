package dispatcher

import (
	"context"
	"github.com/aknEvrnky/notification-system/internal/application/notification"
)

type EventHandler interface {
	Handle(ctx context.Context, payload map[string]any) ([]notification.Notification, error)
}
