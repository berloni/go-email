package sendgrid

import (
	"errors"
	"fmt"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func getClient(apiKey string) *sendgrid.Client {
	client := sendgrid.NewSendClient(apiKey)

	return client
}

func getMessage(sender, recipient, subject, htmlBody, textBody string) *mail.SGMailV3 {
	from := mail.NewEmail("Sender", sender)
	to := mail.NewEmail("Recipient", recipient)
	plainTextContent := textBody
	htmlContent := htmlBody
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	return message
}

func Send(apiKey, sender, recipient, subject, htmlBody, textBody string) error {
	client := getClient(apiKey)
	message := getMessage(sender, recipient, subject, htmlBody, textBody)

	fmt.Println("API KEY: ", apiKey)
	fmt.Println("CLIENT: ", client)
	fmt.Println("MESSAGE: ", message)

	response, err := client.Send(message)
	if err != nil {
		return err
	}
	fmt.Println("RESPONSE: ", response)
	fmt.Println("STATUS CODE: ", response.StatusCode)
	if response.StatusCode >= 400 {
		return errors.New(response.Body)
	}

	return nil
}
