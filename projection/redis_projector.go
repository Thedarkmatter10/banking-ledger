package projection

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Thedarkmatter10/ledger-service/cache"
	"github.com/Thedarkmatter10/ledger-service/model"
)

func ProcessEvent(event model.Event) {
	ctx := context.Background()

	switch event.Type {
	case "Deposited", "Withdrawn":
		var payload model.TransactionPayload
		err := json.Unmarshal([]byte(event.Payload), &payload)
		if err != nil {
			log.Println("Error unmarshalling:", err)
			return
		}

		delta := payload.Amount
		if event.Type == "Withdrawn" {
			delta = -delta
		}

		key := "balance:" + payload.AccountID
		cache.RDB.IncrBy(ctx, key, delta)
	default:
		log.Println("Unknown event type:", event.Type)
	}
}
