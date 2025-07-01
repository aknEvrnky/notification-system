package factory

import (
	"errors"
	"github.com/aknEvrnky/notification-system/internal/adapters/sms/log"
	"github.com/aknEvrnky/notification-system/internal/ports"
	"github.com/aknEvrnky/notification-system/pkg/config"
)

// ErrInvalidSmsPort is returned when an invalid sms port is specified.
var ErrInvalidSmsPort = errors.New("invalid sms port specified")

func NewSmsPort(cfg *config.Config) (ports.SmsPort, error) {
	switch cfg.SmsDriver {
	case "log":
		return log.NewAdapter(), nil
	default:
		return nil, ErrInvalidSmsPort
	}
}
