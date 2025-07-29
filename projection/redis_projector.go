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
			log.Println("Error unmarshalling transaction payload:", err)
			return
		}

		delta := payload.Amount
		if event.Type == "Withdrawn" {
			delta = -delta
		}

		key := "balance:" + payload.AccountID
		cache.RDB.IncrBy(ctx, key, delta)

	case "AccountCreated":
		var payload model.AccountCreatedPayload
		err := json.Unmarshal([]byte(event.Payload), &payload)
		if err != nil {
			log.Println("Error unmarshalling account payload:", err)
			return
		}

		// Set initial balance = 0 in Redis
		key := "balance:" + payload.AccountID
		err = cache.RDB.Set(ctx, key, 0, 0).Err()
		if err != nil {
			log.Println("Error setting initial balance in Redis:", err)
			return
		}
	default:
		log.Println("Unknown event type:", event.Type)
	}
}
