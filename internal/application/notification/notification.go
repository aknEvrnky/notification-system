package notification

import (
	"errors"
	"github.com/aknEvrnky/notification-system/internal/application/notification/messages"
)

type NotificationChannel string

const (
	ChannelMail NotificationChannel = "mail"
	ChannelSms  NotificationChannel = "sms"
	ChannelPush NotificationChannel = "push"
)

var (
	ErrChannelNotSupported   = errors.New("notification channel not supported")
	ErrMissingImplementation = errors.New("missing implementation for notification channel")
)

type Notification interface {
	Channels() []NotificationChannel
}

type MailNotifiable interface {
	ToMail() (*messages.MailMessage, error)
}

type SmsNotifiable interface {
	ToSms() (*messages.SmsMessage, error)
}

type PushNotifiable interface {
	ToPush() (*messages.PushMessage, error)
}
