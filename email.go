package main

import (
	"fmt"
	"net/smtp"
	"os"

	emailLib "github.com/scorredoira/email"
)

func sendMail(email Email) error {

	var err error

	auth := smtp.PlainAuth(
		"",
		email.Mailbox,
		email.Password,
		email.Host,
	)

	if os.Getenv("EMAIL_SEND") != "0" {

		m := emailLib.NewMessage(email.Subject, string(email.body))
		m.From = email.From
		m.To = email.To

		for _, filename := range email.attachments {
			m.Attach(filename)
		}

		err = emailLib.Send(fmt.Sprintf("%s:%s", email.Host, email.Port), auth, m)
	}

	return err
}
