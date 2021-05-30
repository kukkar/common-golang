package emailnotificaiton

type SDK interface {
	SendMail(req SendOTPReq) error
}
