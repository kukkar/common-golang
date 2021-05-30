package swipe

import "errors"

const (
	HeaderHash       = "hash"
	HeaderClientName = "clientName"
)

var ErrorServiceFailed = errors.New("service return with error")
