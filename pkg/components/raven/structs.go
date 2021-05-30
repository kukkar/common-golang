package raven

import (
	"github.com/kukkar/raven"
)

// Define A Farm Type
type FARM_TYPE string

// Allowed farms
const REDIS_CLUSTER_FARM FARM_TYPE = raven.FARM_TYPE_REDISCLUSTER
const REDIS_FARM FARM_TYPE = raven.FARM_TYPE_REDIS

// Config for Raven
type RavenConfig struct {
	Ftype  FARM_TYPE
	Config interface{}
	Logger raven.Logger
	//NewRelicApp newrelic.Application
}

// Config for RedisConf
type RedisConf struct {
	Addr     string
	Password string
	PoolSize int
}
