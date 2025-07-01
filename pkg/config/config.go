package config

import (
	"os"
	"strconv"
)

type Config struct {
	ApplicationPort   int
	BasicAuthUsername string
	BasicAuthPassword string

	MailDriver string

	SmsDriver string

	PushDriver string
}

func NewConfig() *Config {
	port, err := strconv.Atoi(os.Getenv("APPLICATION_PORT"))
	if err != nil {
		panic("Invalid APPLICATION_PORT value: " + os.Getenv("APPLICATION_PORT"))
	}

	return &Config{
		ApplicationPort:   port,
		BasicAuthUsername: os.Getenv("BASIC_AUTH_USERNAME"),
		BasicAuthPassword: os.Getenv("BASIC_AUTH_PASSWORD"),
		MailDriver:        os.Getenv("MAIL_DRIVER"),
		SmsDriver:         os.Getenv("SMS_DRIVER"),
		PushDriver:        os.Getenv("PUSH_DRIVER"),
	}
}
