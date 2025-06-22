package main

import (
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

	application := api.NewApplication()
	zap.L().Info("Application initialized")

	httpAdapter := http.NewAdapter(application, cfg.ApplicationPort)
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
