package utils

import "fmt"
import "github.com/mailgun/mailgun-go"

var (
	publicApiKey string = "xxxx"
	apiKey       string = "xxxx"
	domain       string = "xxxx.mailgun.org"
)

func SendSimpleMessage(domain, apiKey string) (string, error) {
	mg := mailgun.NewMailgun(domain, apiKey, publicApiKey)
	m := mg.NewMessage(
		"Hackman-Postmaster <hackman-postmaster@xxxx.mailgun.org>",
		"Hello",
		"Testing some Mailgun awesomeness!",
		"hackpravj@gmail.com",
	)
	_, id, err := mg.Send(m)
	return id, err
}
