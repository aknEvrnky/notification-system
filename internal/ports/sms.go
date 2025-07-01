package ports

import (
	"context"
	"github.com/aknEvrnky/notification-system/internal/application/notification/messages"
)

type SmsPort interface {
	Send(ctx context.Context, message *messages.SmsMessage) error
}
