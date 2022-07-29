package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

var wg sync.WaitGroup

const topic = "payment.success"

func main() {

	wg.Add(3)

	go func() {
		consumer1(topic)
		wg.Done()
	}()

	// go func() {
	// 	consumer2(topic)
	// 	wg.Done()
	// }()

	// go func() {
	// 	consumer3(topic)
	// 	wg.Done()
	// }()

	// go func() {
	// 	consumer4(topic)
	// 	wg.Done()
	// }()

	wg.Wait()
}

func consumer1(topic string) {
	logrus.Info("Start Consumer 1")
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9093"},
		GroupID: "shipping",
		Topic:   topic,
		// Partition:   0,
		StartOffset: kafka.LastOffset,
		// Logger:      kafka.LoggerFunc(logInfof),
		// ErrorLogger: kafka.LoggerFunc(logErrorf),
	})
	defer r.Close()

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			logrus.Error(errors.Wrap(err, "read message"))
			break
		}

		logrus.Infof("[CONSUMER-01] GroupID: %v, GroupTopic: %v", r.Config().GroupID, r.Config().GroupTopics)
		logrus.Infof("[CONSUMER-01] Lag: %v", r.Stats().Lag)
		logrus.Infof("[CONSUMER-01] p:%v offset:%v key:%s val:%s\n", m.Partition, m.Offset, string(m.Key), string(m.Value))
	}

	logrus.Info("End Consumer 1")
}

func consumer2(topic string) {

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		GroupID: "shipping",
		Topic:   topic,
		// Partition:   1,
		StartOffset: kafka.LastOffset,
		// MinBytes:    5,
		// MaxBytes:    1e6,
		// wait for at most 3 seconds before receiving new data
		MaxWait: 1 * time.Millisecond,
		// Logger:      kafka.LoggerFunc(logInfof),
		// ErrorLogger: kafka.LoggerFunc(logErrorf),
	})
	defer r.Close()

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			logrus.Error(errors.Wrap(err, "read message"))
			break
		}

		logrus.Infof("[CONSUMER-02] GroupID: %v, GroupTopic: %v, Topic: %v", r.Config().GroupID, r.Config().GroupTopics, r.Config().Topic)
		logrus.Infof("[CONSUMER-02] Lag: %v", r.Stats().Lag)
		logrus.Infof("[CONSUMER-02] p:%v offset:%v key:%s val:%s\n", m.Partition, m.Offset, string(m.Key), string(m.Value))
	}

	logrus.Info("End Consumer 2")
}

func consumer3(topic string) {
	logrus.Info("Start Consumer 3")
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		GroupID: "shipping",
		Topic:   topic,
		// Partition:   1,
		StartOffset: kafka.LastOffset,
		// MinBytes:    5,
		// MaxBytes:    1e6,
		// wait for at most 3 seconds before receiving new data
		MaxWait: 1 * time.Millisecond,
		// Logger:      kafka.LoggerFunc(logInfof),
		// ErrorLogger: kafka.LoggerFunc(logErrorf),
	})
	defer r.Close()

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			logrus.Error(errors.Wrap(err, "read message"))
			break
		}

		logrus.Infof("[CONSUMER-03] p:%v offset:%v key:%s val:%s\n", m.Partition, m.Offset, string(m.Key), string(m.Value))
	}

	logrus.Info("End Consumer 3")
}

func consumer4(topic string) {
	logrus.Info("Start Consumer 4")
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		GroupID: "shipping",
		Topic:   topic,
		// Partition:   1,
		StartOffset: kafka.LastOffset,
		// MinBytes:    5,
		// MaxBytes:    1e6,
		// wait for at most 3 seconds before receiving new data
		MaxWait: 1 * time.Millisecond,
		// Logger:      kafka.LoggerFunc(logInfof),
		// ErrorLogger: kafka.LoggerFunc(logErrorf),
	})
	defer r.Close()

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			logrus.Error(errors.Wrap(err, "read message"))
			break
		}

		logrus.Infof("[CONSUMER-04] p:%v offset:%v key:%s val:%s\n", m.Partition, m.Offset, string(m.Key), string(m.Value))
	}

	logrus.Info("End Consumer 4")
}

func logInfof(msg string, a ...interface{}) {
	logrus.Infof(msg, a...)
	fmt.Println()
}

func logErrorf(msg string, a ...interface{}) {
	logrus.Errorf(msg, a...)
	fmt.Println()
}
