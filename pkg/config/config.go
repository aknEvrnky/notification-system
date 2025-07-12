package config

import (
	"go.uber.org/zap"
	"os"
	"strconv"
)

type Config struct {
	ApplicationPort   int
	BasicAuthUsername string
	BasicAuthPassword string

	MailDriver   string
	MailUser     string
	MailPassword string
	MailHost     string
	MailPort     int

	SmsDriver string

	PushDriver string

	Dsn string
}

func NewConfig() *Config {
	appPort, err := strconv.Atoi(os.Getenv("APPLICATION_PORT"))
	if err != nil {
		zap.L().Error("Invalid APPLICATION_PORT value, using default 8080", zap.Error(err))
		appPort = 8080 // Default application port
	}

	smtpPort, err := strconv.Atoi(os.Getenv("MAIL_PORT"))
	if err != nil {
		zap.L().Error("Invalid MAIL_PORT value, using default 587", zap.Error(err))
		smtpPort = 587 // Default SMTP port
	}

	return &Config{
		ApplicationPort:   appPort,
		BasicAuthUsername: os.Getenv("BASIC_AUTH_USERNAME"),
		BasicAuthPassword: os.Getenv("BASIC_AUTH_PASSWORD"),
		MailDriver:        os.Getenv("MAIL_DRIVER"),
		MailUser:          os.Getenv("MAIL_USER"),
		MailPassword:      os.Getenv("MAIL_PASSWORD"),
		MailHost:          os.Getenv("MAIL_HOST"),
		MailPort:          smtpPort,
		SmsDriver:         os.Getenv("SMS_DRIVER"),
		PushDriver:        os.Getenv("PUSH_DRIVER"),
		Dsn:               os.Getenv("DSN"),
	}
}
