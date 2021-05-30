package PG

import "errors"

const (
	HeaderHash       = "hash"
	HeaderClientName = "clientName"
	HeaderMID        = "mid"
)

var ErrorServiceFailed = errors.New("PG-service return with error")
