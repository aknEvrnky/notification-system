package main

import (
	"github.com/aknEvrnky/notification-system/internal/adapters/factory"
	"github.com/aknEvrnky/notification-system/internal/adapters/http"
	"github.com/aknEvrnky/notification-system/internal/application/core/api"
	"github.com/aknEvrnky/notification-system/pkg/config"
	_ "github.com/aknEvrnky/notification-system/pkg/logger"
	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.NewConfig()
	zap.L().Info("Configuration loaded")

	// create ports with factory
	mailer, err := factory.NewMailPort(cfg)
	if err != nil {
		zap.L().Fatal("Failed to create mail port", zap.Error(err))
	}

	zap.L().Info("Mail port created", zap.String("driver", cfg.MailDriver))
	sms, err := factory.NewSmsPort(cfg)
	if err != nil {
		zap.L().Fatal("Failed to create SMS port", zap.Error(err))
	}

	pusher, err := factory.NewPushPort(cfg)
	if err != nil {
		zap.L().Fatal("Failed to create push port", zap.Error(err))
	}

	application := api.NewApplication(cfg, mailer, sms, pusher)
	zap.L().Info("Application initialized")

	httpAdapter := http.NewAdapter(application, cfg)
	zap.L().Info("HTTP adapter created", zap.Int("port", cfg.ApplicationPort))

	go func() {
		zap.L().Info("Starting HTTP server on port", zap.Int("port", cfg.ApplicationPort))
		if err := httpAdapter.Run(); err != nil {
			panic(err)
		}
	}()

	// Wait for term signal
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan

	zap.L().Info("Received shutdown signal, shutting down server...")

	if err := httpAdapter.Shutdown(); err != nil {
		panic(err)
	}
}
