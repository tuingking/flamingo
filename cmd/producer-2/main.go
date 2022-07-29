package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/segmentio/kafka-go"
)

const (
	new     = "hotel-new.json"
	newx    = "hotel-new-x.json"
	deleted = "hotel-deleted.json"
)

// NOTE: untuk test kafka listener TIX-JADE topic: com.tiket.tix.hotel.search-syncHotel

func main() {
	payload := new //! ganti di sini (case: new hotel or deleted hotel)

	b, err := ioutil.ReadFile("./" + payload)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))

	// init kafka client
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "com.tiket.tix.hotel.search-syncHotel",
		// Balancer: &kafka.Hash{},
	})
	defer w.Close()

	w.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte("🐶🐶🐶--Hotel ID--🐶🐶🐶"),
		Value: []byte(string(b)),
	})
}
