package notification

import (
	"context"
	"github.com/aknEvrnky/notification-system/internal/ports"
	"golang.org/x/sync/errgroup"
)

type NotificationService struct {
	mailSender ports.MailPort
	smsSender  ports.SmsPort
	pushSender ports.PushPort
}

func NewNotificationService(mailPort ports.MailPort, smsPort ports.SmsPort, pushPort ports.PushPort) *NotificationService {
	return &NotificationService{
		mailSender: mailPort,
		smsSender:  smsPort,
		pushSender: pushPort,
	}
}

func (s *NotificationService) Send(ctx context.Context, n Notification) error {
	g, ctx := errgroup.WithContext(ctx)

	for _, ch := range n.Channels() {
		ch := ch // capture range variable
		g.Go(func() error {
			switch ch {
			case ChannelMail:
				if m, ok := n.(MailNotifiable); ok {
					msg, err := m.ToMail()
					if err != nil {
						return err
					}
					return s.mailSender.Send(ctx, msg)
				}

				return ErrMissingImplementation
			case ChannelSms:
				if m, ok := n.(SmsNotifiable); ok {
					msg, err := m.ToSms()
					if err != nil {
						return err
					}
					return s.smsSender.Send(ctx, msg)
				}

				return ErrMissingImplementation
			case ChannelPush:
				if m, ok := n.(PushNotifiable); ok {
					msg, err := m.ToPush()
					if err != nil {
						return err
					}
					return s.pushSender.Send(ctx, msg)
				}

				return ErrMissingImplementation
			default:
				return ErrChannelNotSupported
			}
		})
	}

	return g.Wait()
}
