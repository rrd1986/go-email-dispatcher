package main

import (
	"log"

	"github.com/rrd1986/go-email-dispatcher/dispatcher"
)

func main() {
	// Example usage:
	exchangeEmailDispatcher := dispatcher.NewEmailDispatcher(
		"exchange.example.com",
		587,
		"your_username",
		"your_password",
	)

	emailToSend := exchangeEmailDispatcher.Compose(
		[]string{"recipient@example.com"},
		"Subject of your email",
		"This is the plain text body of your email.",
		"<b>This is the HTML body of your email.</b>",
	)

	if err := exchangeEmailDispatcher.Send(emailToSend); err != nil {
		log.Fatal(err)
	}
	log.Println("Email sent successfully")
}
