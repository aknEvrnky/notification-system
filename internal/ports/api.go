package ports

import "context"

type ApiPort interface {
	GetVersion() string
	TriggerNotification(ctx context.Context, eventType string, payload map[string]any) error
}
