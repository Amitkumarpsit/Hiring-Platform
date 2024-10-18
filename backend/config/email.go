package config

import (
	"net/smtp"
)

type EmailConfig struct {
	SMTPHost     string
	SMTPPort     string
	SMTPUsername string
	SMTPPassword string
}

var MailConfig = EmailConfig{
	SMTPHost:     "smtp.gmail.com",
	SMTPPort:     "587",
	SMTPUsername: "siddharth63717@gmail.com", // Replace with your actual email
	SMTPPassword: "wpkmgavyvgxdmvbs",         // Replace with your actual app password
}

func SendEmail(to string, subject string, body string) error {
	from := MailConfig.SMTPUsername
	password := MailConfig.SMTPPassword

	// Set up authentication information.
	auth := smtp.PlainAuth("", from, password, MailConfig.SMTPHost)

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject = "Subject: " + subject + "\n"
	msg := []byte(subject + mime + body)

	// Connect to the server, authenticate, set up the message and send it
	err := smtp.SendMail(
		MailConfig.SMTPHost+":"+MailConfig.SMTPPort,
		auth,
		from,
		[]string{to},
		msg,
	)

	return err
}
