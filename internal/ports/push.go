package ports

import (
	"context"
	"github.com/aknEvrnky/notification-system/internal/application/notification/messages"
)

type PushPort interface {
	Send(ctx context.Context, message *messages.PushMessage) error
}
