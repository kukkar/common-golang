package apiGateway

import "errors"

const (
	HeaderHash       = "hash"
	HeaderClientName = "clientName"
	HeaderMID        = "mid"
)

var ErrorServiceFailed = errors.New("service return with error")
