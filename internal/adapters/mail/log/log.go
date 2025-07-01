package log

import (
	"context"
	"github.com/aknEvrnky/notification-system/internal/application/notification/messages"
	"go.uber.org/zap"
)

type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (a *Adapter) Send(ctx context.Context, message *messages.MailMessage) error {
	zap.L().Info("Sending email", zap.String("to", message.To), zap.String("subject", message.Subject), zap.String("body", message.Body))

	return nil
}
