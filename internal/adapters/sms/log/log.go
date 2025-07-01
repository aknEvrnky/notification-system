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

func (a *Adapter) Send(ctx context.Context, message *messages.SmsMessage) error {
	zap.L().Info("Sending SMS", zap.String("phoneNumber", message.PhoneNumber), zap.String("text", message.Message))

	return nil
}
