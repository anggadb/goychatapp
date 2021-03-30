package lib

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func SendMail(subject, from, body string, to, cc, attachments []string) error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}
	port, err := strconv.Atoi(os.Getenv("MAIL_PORT"))
	if err != nil {
		return err
	}
	recipients := strings.Join(to, ",")
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", from)
	mailer.SetHeader("To", recipients)
	if len(cc) != 0 {
		for a := 0; a < len(cc); a++ {
			mailer.SetAddressHeader("Cc", cc[a], "")
		}
	}
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)
	if len(attachments) != 0 {
		for a := 0; a < len(attachments); a++ {
			mailer.Attach(attachments[a])
		}
	}
	dialer := gomail.NewDialer(os.Getenv("SMTP"), port, os.Getenv("MAIL_USER"), os.Getenv("MAIL_PASSWORD"))
	err = dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}
	return nil
}
