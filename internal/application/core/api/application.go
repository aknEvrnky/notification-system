package api

import (
	"github.com/aknEvrnky/notification-system/internal/application/notification"
	"github.com/aknEvrnky/notification-system/internal/ports"
	"github.com/aknEvrnky/notification-system/pkg/config"
)

type Application struct {
	config   *config.Config
	Notifier *notification.NotificationService
}

func NewApplication(cfg *config.Config, mailPort ports.MailPort, smsPort ports.SmsPort, pushPort ports.PushPort) *Application {
	return &Application{
		config:   cfg,
		Notifier: notification.NewNotificationService(mailPort, smsPort, pushPort),
	}
}

func (a *Application) GetVersion() string {
	return "0.0.1-dev"
}
