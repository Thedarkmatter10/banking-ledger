package repository

import (
	"github.com/Thedarkmatter10/ledger-service/db"
	"github.com/Thedarkmatter10/ledger-service/model"
)

func SaveEvent(event model.Event) error {
	_, err := db.DB.Exec(`
        INSERT INTO events (id, aggregate_id, type, payload, timestamp)
        VALUES ($1, $2, $3, $4, $5)
    `, event.ID, event.AggregateID, event.Type, event.Payload, event.Timestamp)
	return err
}

func GetEventsByAggregateID(aggregateID string) ([]model.Event, error) {
	rows, err := db.DB.Query(`SELECT id, aggregate_id, type, payload, timestamp FROM events WHERE aggregate_id = $1 ORDER BY timestamp DESC`, aggregateID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []model.Event
	for rows.Next() {
		var event model.Event
		if err := rows.Scan(&event.ID, &event.AggregateID, &event.Type, &event.Payload, &event.Timestamp); err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func AccountExists(accountID string) (bool, error) {
	var count int
	err := db.DB.QueryRow(`SELECT COUNT(*) FROM accounts WHERE id = $1`, accountID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
