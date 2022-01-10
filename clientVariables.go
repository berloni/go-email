package email

type SendgridVariables struct {
	ApiKey string `json:"apiKey"`
}

type SesVariables struct {
	AwsRegion       string `json:"awsRegion"`
	AccessKeyID     string `json:"accessKey"`
	SecretAccessKey string `json:"secretKey"`
}

// TODO: Add more clients templates and possible variables
type ClientVariables struct {
	Name         string            `json:"name"`
	SendgridVars SendgridVariables `json:"sendgridVars"`
	SesVars      SesVariables      `json:"sesVars"`
}
