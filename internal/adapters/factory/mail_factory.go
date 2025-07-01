package factory

import (
	"errors"
	"github.com/aknEvrnky/notification-system/internal/adapters/mail/log"
	"github.com/aknEvrnky/notification-system/internal/ports"
	"github.com/aknEvrnky/notification-system/pkg/config"
)

// ErrInvalidMailPort is returned when an invalid mail port is specified.
var ErrInvalidMailPort = errors.New("invalid mail port specified")

func NewMailPort(cfg *config.Config) (ports.MailPort, error) {
	switch cfg.MailDriver {
	case "log":
		return log.NewAdapter(), nil
	default:
		return nil, ErrInvalidMailPort
	}
}
