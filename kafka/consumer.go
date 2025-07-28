package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Thedarkmatter10/ledger-service/model"
	"github.com/Thedarkmatter10/ledger-service/projection"
	"github.com/segmentio/kafka-go"
)

func StartConsumer() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "ledger-events",
		GroupID: "projection-service",
	})

	for {
		msg, _ := r.ReadMessage(context.Background())
		var event model.Event
		json.Unmarshal(msg.Value, &event)
		log.Println("Consumed:", event.Type)
		projection.ProcessEvent(event)
	}
}
