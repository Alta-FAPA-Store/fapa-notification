package email

import (
	"fmt"
	"os"

	"gopkg.in/mail.v2"
)

func SendEmail(emailAddress []byte) {
	from := os.Getenv("FROM_EMAIL")
	password := os.Getenv("EMAIL_PASS")
	fmt.Println(from, password)

	email := mail.NewMessage()

	email.SetHeader("From", from)
	email.SetHeader("To", string(emailAddress))
	email.SetHeader("Subject", "Welcome to FAPA-STORE!")
	email.SetBody("text/plain", "Thank you for register")

	send := mail.NewDialer("smtp.gmail.com", 587, from, password)

	if err := send.DialAndSend(email); err != nil {
		fmt.Println(err)
		panic(err)
	}

}
