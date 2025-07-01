package smtp

import (
	"context"
	"fmt"
	"github.com/aknEvrnky/notification-system/internal/application/notification/messages"
	"net/smtp"
)

type Adapter struct {
	user     string
	password string
	host     string
	port     int
}

func NewAdapter(user, password, host string, port int) *Adapter {
	return &Adapter{
		user:     user,
		password: password,
		host:     host,
		port:     port,
	}
}

func (a Adapter) Send(ctx context.Context, message *messages.MailMessage) error {
	auth := smtp.PlainAuth("", a.user, a.password, a.host)
	host := fmt.Sprintf("%s:%d", a.host, a.port)

	to := []string{message.To}
	body := fmt.Sprintf("Subject: %s\r\n\r\n%s", message.Subject, message.Body)

	return smtp.SendMail(host, auth, a.user, to, []byte(body))
}
