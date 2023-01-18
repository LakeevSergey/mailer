package sender

import (
	"bytes"
	"fmt"

	"github.com/LakeevSergey/mailer/internal/domain/entity"
	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
	"gopkg.in/gomail.v2"
)

type SMTPSender struct {
	host     string
	port     int
	user     string
	password string
}

func NewSMTPSender(host string, port int, user string, password string) *SMTPSender {
	return &SMTPSender{
		host:     host,
		port:     port,
		user:     user,
		password: password,
	}
}

func (s *SMTPSender) Send(mail entity.Mail) error {
	m := gomail.NewMessage()
	m.SetHeader("To", mail.SendTo...)
	m.SetHeader("From", fmt.Sprintf("\"%s\" <%s>", mail.SendFrom.Name, mail.SendFrom.Email))
	m.SetHeader("Subject", mail.Title)
	m.SetBody("text/html", mail.Body)

	for _, file := range mail.Attachments {
		m.Attach(
			file.Info.Path,
			gomail.Rename(file.Info.FileName),
			gomail.SetHeader(
				map[string][]string{
					"Content-Type": {file.Info.Mime},
				},
			))
	}

	buf := new(bytes.Buffer)

	m.WriteTo(buf)

	addr := fmt.Sprintf("%s:%d", s.host, s.port)

	client := sasl.NewLoginClient(s.user, s.password)

	err := smtp.SendMail(addr, client, mail.SendFrom.Email, mail.SendTo, buf)

	if err != nil {
		return err
	}

	return nil
}
