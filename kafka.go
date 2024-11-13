package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func consume(kafkaBroker, kafkaTopic string, c chan int) error {
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
			break
		}
		c <- 1 // Send a trigger message to the channel
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		if err := r.CommitMessages(ctx, m); err != nil {
			return fmt.Errorf("failed to commit messages: %e", err)
		}
	}

	if err := r.Close(); err != nil {
		return fmt.Errorf("failed to close connection: %e", err)
	}

	return nil
}
