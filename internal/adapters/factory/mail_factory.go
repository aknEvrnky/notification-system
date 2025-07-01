package factory

import (
	"errors"
	"github.com/aknEvrnky/notification-system/internal/adapters/mail/log"
	"github.com/aknEvrnky/notification-system/internal/adapters/mail/smtp"
	"github.com/aknEvrnky/notification-system/internal/ports"
	"github.com/aknEvrnky/notification-system/pkg/config"
)

// ErrInvalidMailPort is returned when an invalid mail port is specified.
var ErrInvalidMailPort = errors.New("invalid mail port specified")

func NewMailPort(cfg *config.Config) (ports.MailPort, error) {
	switch cfg.MailDriver {
	case "log":
		return log.NewAdapter(), nil
	case "smtp":
		if cfg.MailUser == "" || cfg.MailPassword == "" || cfg.MailHost == "" || cfg.MailPort <= 0 {
			return nil, errors.New("missing required SMTP configuration parameters")
		}
		return smtp.NewAdapter(cfg.MailUser, cfg.MailPassword, cfg.MailHost, cfg.MailPort), nil
	default:
		return nil, ErrInvalidMailPort
	}
}
