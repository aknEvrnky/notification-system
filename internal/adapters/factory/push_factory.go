package factory

import (
	"errors"
	"github.com/aknEvrnky/notification-system/internal/adapters/push/log"
	"github.com/aknEvrnky/notification-system/internal/ports"
	"github.com/aknEvrnky/notification-system/pkg/config"
)

// ErrInvalidPushPort is returned when an invalid push port is specified.
var ErrInvalidPushPort = errors.New("invalid push port specified")

func NewPushPort(cfg *config.Config) (ports.PushPort, error) {
	switch cfg.PushDriver {
	case "log":
		return log.NewAdapter(), nil
	default:
		return nil, ErrInvalidPushPort
	}
}
