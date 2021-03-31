package utils

import (
	"fmt"
	"lireddit/env"
	"log"
	"net/smtp"

	"github.com/ilyakaznacheev/cleanenv"
)

func init() {
	if err := cleanenv.ReadEnv(&env.Cfg); err != nil {
		log.Fatal("cannot read rend")
	}
}

func SendEmail(to []string) bool {
	from := "bob@bob.com"
	msg := fmt.Sprintf(
		"To: %s \r\n"+
			"From: hello@schadokar.dev\r\n"+
			"Subject: Hello Gophers!\r\n"+
			"\r\n"+
			"This is the email is sent using golang and sendinblue.\r\n",
		to,
	)

	smtpAddress := "localhost:1025"
	err := smtp.SendMail(smtpAddress, nil, from, to, []byte(msg))

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}
