package emailnotificaiton

import (
	"fmt"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

//SendGridAccessor
type SendGridAccessor struct {
	APIKey string
}

func (this SendGridAccessor) SendMail(req SendOTPReq) error {

	from := mail.NewEmail(req.From.Name, req.From.Email)
	subject := req.Subject
	var toEmail *mail.Email
	if len(req.To) > 0 {
		toEmail = mail.NewEmail(req.To[0].Name, req.To[0].Email)
	}

	message := mail.NewSingleEmail(from, subject, toEmail, req.TextContent, req.HTMLContent)
	client := sendgrid.NewSendClient(this.APIKey)
	response, err := client.Send(message)
	if err != nil {
		return err
	}
	if response.StatusCode > 299 {
		return fmt.Errorf("error with code %v and response %v", response.StatusCode, response.Body)
	}
	return nil
}
