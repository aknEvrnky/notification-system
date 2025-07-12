package api

import (
	"context"
	"github.com/aknEvrnky/notification-system/internal/application/notification/dispatcher"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func (a *Application) TriggerNotification(ctx context.Context, eventType string, payload map[string]any) error {
	zap.L().Info("Triggering notification", zap.String("type", eventType))

	// 1. find notifications to be triggered by event name
	handler, err := dispatcher.ResolveHandler(eventType)
	if err != nil {
		zap.L().Error("Failed to resolve handler", zap.String("event", eventType), zap.Error(err))
		return err
	}

	notifications, err := handler.Handle(ctx, payload)
	if err != nil {
		zap.L().Error("Failed to handle event", zap.String("event", eventType), zap.Error(err))
		return err
	}

	// 2. trigger each notification in a goroutine
	g, ctx := errgroup.WithContext(ctx)

	for _, notif := range notifications {
		n := notif // capture range variable
		g.Go(func() error {
			return a.Notifier.Send(ctx, n)
		})
	}

	if err := g.Wait(); err != nil {
		zap.L().Error("Failed to send notifications", zap.String("event", eventType), zap.Error(err))
		return err
	}

	return nil
}
