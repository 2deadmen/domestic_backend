package utils

import (
	"log"
	"os"

	"gopkg.in/gomail.v2"
)

func SendEmail(to string, subject string, body string) error {
	var emailPasscode string = os.Getenv("EMAIL_PASS")
	var emailEmail string = os.Getenv("EMAIL")
	log.Println(to, subject, body, emailEmail, emailPasscode)
	m := gomail.NewMessage()
	m.SetHeader("From", emailEmail) // Replace with your email
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer("smtp.gmail.com", 587, emailEmail, emailPasscode) // Replace SMTP details

	if err := d.DialAndSend(m); err != nil {
		log.Println("Failed to send email:", err)
		return err
	}

	return nil
}
