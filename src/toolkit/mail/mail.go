package mail

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/feerdim/boilerplate-golang/log"
	"github.com/feerdim/boilerplate-golang/src/constant"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/pkg/errors"
	"gopkg.in/gomail.v2"
)

type Mail struct {
	Driver      string
	Dialer      *gomail.Dialer
	Mailgun     *mailgun.MailgunImpl
	SenderName  string
	SenderEmail string
}

func NewMail() (mailDialer *Mail, err error) {
	mailDialer = &Mail{
		Driver:      os.Getenv("MAIL_DRIVER"),
		SenderName:  os.Getenv("MAIL_SENDER_NAME"),
		SenderEmail: os.Getenv("MAIL_SENDER_EMAIL"),
	}

	switch mailDialer.Driver {
	case "smtp":
		portStr := os.Getenv("MAIL_SMTP_PORT")

		port, err := strconv.Atoi(portStr)
		if err != nil {
			return nil, errors.Wrapf(err, "error parse int on port mail env : %s", portStr)
		}

		mailDialer.Dialer = gomail.NewDialer(os.Getenv("MAIL_SMTP_HOST"), port, os.Getenv("MAIL_SMTP_USERNAME"), os.Getenv("MAIL_SMTP_PASSWORD"))
	case "mailgun":
		mailDialer.Mailgun = mailgun.NewMailgun(os.Getenv("MAIL_MAILGUN_DOMAIN"), os.Getenv("MAIL_MAILGUN_API_KEY"))
	}

	return
}

func (m Mail) SendMail(ctx context.Context, subject, body string, recipients ...string) (msg, id string, err error) {
	sender := fmt.Sprintf("%s <%s>", m.SenderName, m.SenderEmail)

	switch m.Driver {
	case "smtp":
		message := gomail.NewMessage()
		message.SetHeader("From", sender)
		message.SetHeader("To", recipients...)
		message.SetHeader("Subject", subject)
		message.SetBody("text/html", body)

		if err = m.Dialer.DialAndSend(message); err != nil {
			log.WithContext(ctx).Error(err, "error send email with smtp driver")
			return
		}

		log.PrintDebug("success send email with smtp driver")
	case "mailgun":
		message := mailgun.NewMessage(sender, subject, body, recipients...)

		msg, id, err = m.Mailgun.Send(ctx, message)
		if err != nil {
			log.WithContext(ctx).Error(err, "error send email with mailgun driver")
			return
		}

		log.PrintDebug("success send email with mailgun driver", "id", id, "resp", msg)
	default:
		err = constant.ErrUnknownSource
		log.WithContext(ctx).Error(err, "error unknown mail driver")
	}

	return
}
