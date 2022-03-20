package kafka

import (
	"github.com/segmentio/kafka-go"
)

func InitProducer(cfg KafkaConfig, topic string) *kafka.Writer {
	writter := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	return writter
}
