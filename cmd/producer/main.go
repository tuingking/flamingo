package main

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

const topic = "payment.success"

var wg sync.WaitGroup

func main() {
	wg.Add(1)

	//*** Producer
	go func() {
		spam(1_000_000)
		wg.Done()
	}()

	wg.Wait()
}

func spam(n int64) {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9093"},
		Topic:    topic,
		Balancer: &kafka.Hash{},
		// Balancer: &kafka.LeastBytes{},
	})
	defer w.Close()

	keys := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var count int64
	for {
		key := strconv.Itoa(keys[rand.Intn(10)])
		val := fmt.Sprintf("Product %d", count)

		logrus.Infof("[SEND] %s:%s", key, val)
		if err := w.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(key),
			Value: []byte(val),
		}); err != nil {
			logrus.Fatal(errors.Wrap(err, "failed to write message"))
		}

		time.Sleep(10 * time.Millisecond)
		count++

		if count == n {
			break
		}
	}
}
