package mail_api

import (
	"context"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/mail"
	"gopkg.in/gomail.v2"
)

func (a *apiImpl) Send(ctx context.Context, mail *mail.Message) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", mail.From)
	msg.SetHeader("To", mail.To)
	msg.SetHeader("Subject", mail.Subject)
	msg.SetBody("text/html", mail.Body)

	mailer := gomail.NewDialer(a.Host, a.Port, a.Username, a.Password)
	err := mailer.DialAndSend(msg)
	if err != nil {
		return err
	}

	return nil
}
