package otpLess

import "errors"

const (
	HeaderClientId     = "clientid"
	HeaderClientSecret = "clientsecret"
)

var ErrorServiceFailed = errors.New("service return with error")
