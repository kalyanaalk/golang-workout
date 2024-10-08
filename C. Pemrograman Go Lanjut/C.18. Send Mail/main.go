package main

import (
	// "fmt"
	"log"
	// "net/smtp"
	// "strings"

	"gopkg.in/gomail.v2"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "PT. Makmur Subur Jaya <@gmail.com>"
const CONFIG_AUTH_EMAIL = ""
const CONFIG_AUTH_PASSWORD = ""

func main() {
	// to := []string{"@gmail.com", "@gmail.com"}
	// cc := []string{"@gmail.com"}
	// subject := "Test mail"
	// message := "Hellohihihi"

	// err := sendMail(to, cc, subject, message)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// log.Println("Mail sent!")

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", "recipient1@gmail.com", "emaillain@gmail.com")
	mailer.SetAddressHeader("Cc", "tralalala@gmail.com", "Tra Lala La")
	mailer.SetHeader("Subject", "Test mail")
	mailer.SetBody("text/html", "Hello, <b>have a nice day</b>")
	mailer.Attach("./sample.png")

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Mail sent!")
}

// func sendMail(to []string, cc []string, subject, message string) error {
// 	body := "From: " + CONFIG_SENDER_NAME + "\n" +
// 		"To: " + strings.Join(to, ",") + "\n" +
// 		"Cc: " + strings.Join(cc, ",") + "\n" +
// 		"Subject: " + subject + "\n\n" +
// 		message

// 	auth := smtp.PlainAuth("", CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD, CONFIG_SMTP_HOST)
// 	smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)

// 	err := smtp.SendMail(smtpAddr, auth, CONFIG_AUTH_EMAIL, append(to, cc...), []byte(body))
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
