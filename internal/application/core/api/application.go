package api

import (
	"github.com/aknEvrnky/notification-system/internal/application/notification"
	"github.com/aknEvrnky/notification-system/internal/application/notification/dispatcher"
	"github.com/aknEvrnky/notification-system/internal/ports"
	"github.com/aknEvrnky/notification-system/pkg/config"
)

type Application struct {
	config         *config.Config
	Notifier       *notification.NotificationService
	UserRepository ports.UserRepository
}

func NewApplication(cfg *config.Config, mailPort ports.MailPort, smsPort ports.SmsPort, pushPort ports.PushPort, userRepo ports.UserRepository) *Application {
	dispatcher.RegisterHandler("user_followed", dispatcher.NewUserFollowedHandler(userRepo))

	return &Application{
		config:         cfg,
		Notifier:       notification.NewNotificationService(mailPort, smsPort, pushPort),
		UserRepository: userRepo,
	}
}

func (a *Application) GetVersion() string {
	return "0.0.1-dev"
}
