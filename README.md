# GO-EMAIL

`Golang` package to send emails.

At the moment can be used only with `sendgrid` or `aws - simple email service`.

## Examples

### HTML Template

```html
<!DOCTYPE html>

<body>
    <h3>Hi {{ .Username }}, welcome to {{ .AppName }}</h3>
    <p>We are excited to have you started. First, you need to verify your account.</p>
    <p>Use the following link to complete your authentication:</p>
    <a href="{{ .VerificationLink }}" target="_blank">{{ .VerificationLink }}</a>
</body>
```

### Sendgrid

```go
clientVariablesSendgrid := email.ClientVariables{
	Name: "sendgrid",
	SendgridVars: email.SendgridVariables{
		ApiKey: "SG.XXXXXXXXXXXXXXXXXXXXXXX",
	,
}
templateVarsSendgrid := templates.TemplateVariables{
	Username:         "username",
	AppName:          "application Name",
	VerificationLink: "https://api.domanin/verification/email?email=email@email.com",
}
err := email.Send(clientVariablesSendgrid, htmlTemplate, templateVarsSendgrid, "TXT TEMPLATE", "email@email.com", "email@email.com", "Email Subject")
```

### AWS - Ses

```go
clientVariablesAWS := email.ClientVariables{
			Name: "aws-ses",
			SesVars: email.SesVariables{
				AwsRegion:       "us-west-1",
				AccessKeyID:     "XXXXXXXXXXXXXXXXXXXX",
				SecretAccessKey: "XXXXXXXXXXXXXXXXXXXX",
			},
		}
		templateVarsAWS := templates.TemplateVariables{
			Username:         "username",
			AppName:          "Application Name",
			VerificationLink: "https://domanin/verification/email?email=email@email.com",
		}
err := email.Send(clientVariablesAWS, htmlTemplate, templateVarsAWS, "TXT TEMPLATE", "email@email.com", "email@email.com", "Email Subject")
```



