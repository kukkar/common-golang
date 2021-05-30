package raven

import (
	"fmt"
	"strings"

	concurrenthashmap "github.com/kukkar/common-golang/pkg/utils/concurrenthashmap"
	"github.com/kukkar/raven"
)

//
// Define a RavenMap.
//
var ravenMap = concurrenthashmap.New()

// Define a Farm.
type Farm struct {
	*raven.Farm
}

// Define a Logger.
type Logger interface {
	raven.Logger
}

// Define a Message.
type Message struct {
	*raven.Message
}

//Initiate Farm.
func InitFarm(key string, rc RavenConfig) (*Farm, error) {
	farm, err := GetFarm(key)
	if err == nil {
		return farm, nil
	}

	var config interface{}
	switch rc.Ftype {
	case REDIS_CLUSTER_FARM:
		redisConf, ok := rc.Config.(RedisConf)
		if !ok {
			return nil, fmt.Errorf("Expected raven.RedisClusterConf, Got: %T", rc.Config)
		}
		addrs := strings.Split(redisConf.Addr, ",")
		config = raven.RedisClusterConfig{
			Addrs:    addrs,
			Password: redisConf.Password,
			PoolSize: redisConf.PoolSize,
		}
	case REDIS_FARM:
		redisSimpleConf, ok := rc.Config.(RedisConf)
		if !ok {
			return nil, fmt.Errorf("Expected raven.RedisSimpleConf, Got: %T", rc.Config)
		}
		config = raven.RedisSimpleConfig{
			Addr:     redisSimpleConf.Addr,
			PoolSize: redisSimpleConf.PoolSize,
			Password: redisSimpleConf.Password,
		}

	default:
		return nil, fmt.Errorf("How did you landed here?")
	}

	farm1, err1 := raven.InitializeFarm(string(rc.Ftype), config, rc.Logger)
	if err1 != nil {
		return nil, err1
	}
	//make sure to attach newrelic application.
	//farm1.AttachNewRelicApp(rc.NewRelicApp)

	f := &Farm{farm1}
	ravenMap.Put(key, f)
	return f, nil
}

//
// Get farm by Key.
//
func GetFarm(key string) (*Farm, error) {
	finalkey := key
	if val, ok := ravenMap.Get(finalkey); ok {
		return val.(*Farm), nil
	}
	return nil, fmt.Errorf("No Such Farm exists.")
}
