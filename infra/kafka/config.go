package kafka

import "github.com/segmentio/kafka-go"

type KafkaConfig struct {
	Enable  bool
	Brokers []string
	GroupID string

	Producer struct {
		Ack string
	}

	Consumer struct {
		MinBytes    int // min buffer bytes, IMPORTANT: need to set MaxByte if MinBytes is set
		MaxBytes    int // max qty of data that the cluster can response with when polled
		MaxWait     int
		StartOffset int // earliest,
	}
}

func InitReaderConfigDefault(topic string) kafka.ReaderConfig {
	return kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		GroupID:  "go.local.playground",
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	}
}
