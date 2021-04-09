package utils

import (
	"context"
	"fmt"
	"lireddit/cache"
	"lireddit/models"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/antihax/optional"
	"github.com/joho/godotenv"
	MailSlurpClient "github.com/mailslurp/mailslurp-client-go"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("cannot read env")
	}
}

// send email to user to change password
func SendEmail(user *models.User) bool {
	client, ctx := getMailSlurpClient()

	inbox, _, _ := client.InboxControllerApi.CreateInbox(ctx, nil)

	token := RandomString(20)

	redis := cache.ConnectRedis()
	redis.Set(context.Background(), fmt.Sprintf("forgot-password:%s", token), user.Id, time.Hour*24*3) // 3 days

	var body string = fmt.Sprintf(`<a href="%s/change-password/%s">reset password</a>`, os.Getenv("CORS_ORIGIN"), token)

	sendEmailOptions := MailSlurpClient.SendEmailOptions{
		To:      []string{inbox.EmailAddress},
		Subject: "change password",
		Body:    body,
		IsHTML:  true,
	}

	opts := &MailSlurpClient.SendEmailOpts{
		SendEmailOptions: optional.NewInterface(sendEmailOptions),
	}

	res, err := client.InboxControllerApi.SendEmail(ctx, inbox.Id, opts)
	if err != nil {
		return false
	}

	fmt.Println(res.StatusCode)
	return true
}

func getMailSlurpClient() (*MailSlurpClient.APIClient, context.Context) {
	ctx := context.WithValue(
		context.Background(),
		MailSlurpClient.ContextAPIKey,
		MailSlurpClient.APIKey{Key: os.Getenv("EMAIL_KEY")},
	)

	config := MailSlurpClient.NewConfiguration()
	client := MailSlurpClient.NewAPIClient(config)

	return client, ctx
}

// generate random string to work as a token
func RandomString(n int) string {
	var letters = []rune("abcde-fghijklmnop-qrstuvwxyzABCDEFGHIJ-KLMNOP-QRSTUVWXY-Z0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
