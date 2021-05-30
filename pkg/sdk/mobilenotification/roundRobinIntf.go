package mobilenotification

type RoundRobinOTPService interface {
	SendMessage(number string, merchant string) (*SendOTPRes, error)
}
