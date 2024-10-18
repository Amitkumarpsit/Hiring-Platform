package config

import (
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

type EmailConfig struct {
	SMTPHost     string
	SMTPPort     string
	SMTPUsername string
	SMTPPassword string
}

var MailConfig EmailConfig

func init() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Initialize the MailConfig with environment variables
	MailConfig = EmailConfig{
		SMTPHost:     os.Getenv("SMTP_HOST"),
		SMTPPort:     os.Getenv("SMTP_PORT"),
		SMTPUsername: os.Getenv("SMTP_USERNAME"),
		SMTPPassword: os.Getenv("SMTP_PASSWORD"),
	}
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
