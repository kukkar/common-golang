package emailnotificaiton

import "fmt"

//Config
type Config struct {
	ServiceToUse string
	APIKey       string
}

func (this *Config) validateConfig() error {
	if this.ServiceToUse == SendGrid && this.APIKey == "" {
		return fmt.Errorf("sendgrid service required api key")
	}
	return nil
}

type SendOTPReq struct {
	From        UserInfo
	To          []UserInfo
	TextContent string
	HTMLContent string
	Subject     string
}

type UserInfo struct {
	Name  string
	Email string
}
