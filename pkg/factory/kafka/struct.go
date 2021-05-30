package kafka

type KafkaConfig struct {
	Addr          string `json:"Addr"`
	WriterTimeout int    `json:"WriterTimeout"`
	ReaderTimeout int    `json:"ReaderTimeout"`
	ClientID      string `json:"ClientID"`
}
