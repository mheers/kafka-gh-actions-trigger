package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func consume(kafkaBroker, kafkaTopic string, c chan int) {
	// to consume messages
	partition := 0
	groupID := "my-group"

	r := kafka.NewReader(
		kafka.ReaderConfig{
			Brokers:   []string{kafkaBroker},
			Topic:     kafkaTopic,
			Partition: partition,
			GroupID:   groupID,
			MaxBytes:  10e6, // 10MB
		},
	)

	ctx := context.Background()
	for {
		m, err := r.FetchMessage(ctx)
		if err != nil {
			log.Fatalf("failed to fetch message: %e", err)
		}
		c <- 1 // Send a trigger message to the channel
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		if err := r.CommitMessages(ctx, m); err != nil {
			log.Fatalf("failed to commit messages: %e", err)
		}
	}
}
