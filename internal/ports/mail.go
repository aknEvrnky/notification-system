package ports

import (
	"context"
	"github.com/aknEvrnky/notification-system/internal/application/notification/messages"
)

type MailPort interface {
	Send(ctx context.Context, message *messages.MailMessage) error
}
