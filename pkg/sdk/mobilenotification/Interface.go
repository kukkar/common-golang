package mobilenotification

type SDK interface {
	SendMessage(number string, merchant string) (*SendOTPRes, error)
}
