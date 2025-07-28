package kafka

import (
	"context"
	"encoding/json"

	"github.com/Thedarkmatter10/ledger-service/model"
	"github.com/Thedarkmatter10/ledger-service/repository"
	"github.com/segmentio/kafka-go"
)

var Writer = kafka.NewWriter(kafka.WriterConfig{
	Brokers: []string{"localhost:9092"},
	Topic:   "ledger-events",
})

func PublishEvent(event model.Event) error {
	err := repository.SaveEvent(event)
	if err != nil {
		return err
	}
	msg, _ := json.Marshal(event)
	return Writer.WriteMessages(context.Background(), kafka.Message{Value: msg})
}
