package model

import "time"

type Event struct {
	ID          string    `json:"id"`
	AggregateID string    `json:"aggregate_id"`
	Type        string    `json:"type"` // e.g. AccountCreated, Deposited, Withdrawn
	Payload     string    `json:"payload"`
	Timestamp   time.Time `json:"timestamp"`
}

type TransactionPayload struct {
	AccountID string `json:"account_id"`
	Amount    int64  `json:"amount"`
}
