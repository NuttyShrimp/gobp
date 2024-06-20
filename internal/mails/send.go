package mails

import (
	"os"

	"github.com/cli/browser"
	"github.com/studentkickoff/gobp/pkg/config"
	"github.com/wneessen/go-mail"
)

func sendMailViaSMTP(HTMLtmpl, TextTmpl, recipient, subject string) error {
	m := mail.NewMsg()
	if err := m.FromFormat("Student Kick-Off", "am@studentkickoff.be"); err != nil {
		return err
	}
	if err := m.To(recipient); err != nil {
		return err
	}

	m.Subject(subject)
	m.SetBodyString(mail.TypeTextHTML, HTMLtmpl)
	m.SetBodyString(mail.TypeTextPlain, TextTmpl)

	c, err := mail.NewClient(config.GetString("mail.smtp.address"), mail.WithPort(config.GetInt("mail.smtp.port")), mail.WithSMTPAuth(mail.SMTPAuthLogin),
		mail.WithUsername(config.GetString("mail.smtp.username")), mail.WithPassword(config.GetString("mail.smtp.password")))
	if err != nil {
		return err
	}

	if err := c.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

func sendMail(HTMLTmpl, TextTmpl, recipient, subject string) error {
	env := config.GetDefaultString("app.env", "development")
	if env == "development" {
		f, err := os.CreateTemp("", "mail.html")
		if err != nil {
			return err
		}
		_, err = f.WriteString(HTMLTmpl)
		if err != nil {
			return err
		}
		if err := browser.OpenFile(f.Name()); err != nil {
			return err
		}
	} else {
		return sendMailViaSMTP(HTMLTmpl, TextTmpl, recipient, subject)
	}
	return nil
}
