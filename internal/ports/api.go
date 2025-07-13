package ports

import (
	"context"
	"github.com/aknEvrnky/notification-system/internal/application/core/domain"
)

type ApiPort interface {
	GetVersion() string
	TriggerNotification(ctx context.Context, eventType string, payload map[string]any) error
	CreateUser(ctx context.Context, user domain.User) error
	UpdateUser(ctx context.Context, user domain.User) error
	DeleteUser(ctx context.Context, id string) error
}
