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

func (a *Adapter) Send(ctx context.Context, message *messages.PushMessage) error {
	zap.L().Info("Sending push notification", zap.String("title", message.Title), zap.String("body", message.Message), zap.String("deviceToken", message.DeviceToken))

	return nil
}
