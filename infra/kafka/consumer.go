package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

func InitConsumer(cfg KafkaConfig, topic string) {
	reader := kafka.NewReader(InitReaderConfigDefault(topic))
	defer reader.Close()

	logrus.Info("start consuming... !!")

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			logrus.Fatalf("failed to read message", err)
		}
		logrus.Infof("message at topic:%v partition:%v offset:%v	%s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}
