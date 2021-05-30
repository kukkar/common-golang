package kafka

import (
	"fmt"
	"strings"
	"time"

	concurrenthashmap "github.com/kukkar/common-golang/pkg/utils/concurrenthashmap"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
)

const DEFAULT_POOL = "default"

//store cache Adapter
var cacheMap = concurrenthashmap.New()
var keyToConfigMap = make(map[string]interface{})

func InitConfigMap(configMap map[string]interface{}) {

	for key, value := range configMap {
		keyToConfigMap[key] = value
	}
}

func GetPool(topic string) (*kafka.Writer, error) {
	if val, ok := cacheMap.Get(topic); !ok {
		cache, err := getAdapter(topic)
		if err != nil {
			return nil, err
		}
		cacheMap.Put(topic, cache)
		return cache, nil
	} else {
		return val.(*kafka.Writer), nil
	}
}

func getAdapter(topic string) (*kafka.Writer, error) {
	var w *kafka.Writer
	config, err := getConfig(topic)
	if err != nil {
		return nil, err
	}
	dialer := &kafka.Dialer{
		Timeout:  10 * time.Second,
		ClientID: config.ClientID,
	}

	kafkaCluster := strings.Split(config.Addr, ",")
	kafkaconfig := kafka.WriterConfig{
		Brokers:          kafkaCluster,
		Topic:            topic,
		Balancer:         &kafka.LeastBytes{},
		Dialer:           dialer,
		WriteTimeout:     10 * time.Second,
		ReadTimeout:      10 * time.Second,
		CompressionCodec: snappy.NewCompressionCodec(),
	}
	w = kafka.NewWriter(kafkaconfig)
	return w, nil
}

func getConfig(key string) (*KafkaConfig, error) {

	var config KafkaConfig
	if val, ok := keyToConfigMap[key]; ok {
		config = val.(KafkaConfig)
	} else {
		return nil, fmt.Errorf("Wrong Config passed unable to assert to KafkaConfig")
	}
	return &config, nil
}
