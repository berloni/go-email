package email

import (
	"bytes"
	"errors"
	"html/template"

	"github.com/berloni/go-email/sendgrid"
	"github.com/berloni/go-email/ses"
)

func Send(client ClientVariables, template *template.Template, templateVars interface{}, txtTemplate string, sender string, recipient string, subject string) error {
	// set template variables
	var tByt bytes.Buffer
	err := template.Execute(&tByt, templateVars)
	if err != nil {
		return err
	}
	htmlBody := tByt.String()

	// TODO: Add more clients
	switch client.Name {
	case "sendgrid":
		err = sendgrid.Send(client.SendgridVars.ApiKey, sender, recipient, subject, htmlBody, txtTemplate)

	case "aws-ses":
		_, err = ses.Send(client.SesVars.AwsRegion, client.SesVars.AccessKeyID, client.SesVars.SecretAccessKey, sender, recipient, subject, htmlBody, txtTemplate)

	default:
		return errors.New("invalid client")
	}

	return err
}
