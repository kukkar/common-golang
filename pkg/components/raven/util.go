package raven

import (
	"github.com/kukkar/raven"
)

// Define Source Object
func CreateSource(name string, boxes int) raven.Source {
	return raven.CreateSource(name, boxes)
}

// Define Destination Object
func CreateDestination(name string, boxes int, shardlogic func(raven.Message, int) (string, error)) raven.Destination {
	return raven.CreateDestination(name, boxes, raven.ShardHandler(shardlogic))
}

// Define PrepareMessage object.
func PrepareMessage(id, mtype, data, shardkey string) raven.Message {
	return raven.PrepareMessage(id, mtype, data, shardkey)
}
