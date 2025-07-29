package command

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/Thedarkmatter10/ledger-service/cache"
	"github.com/Thedarkmatter10/ledger-service/kafka"
	"github.com/Thedarkmatter10/ledger-service/model"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

type CreateAccountRequest struct {
	UserName string `json:"user_name" binding:"required"`
}

func CreateAccount(c *gin.Context) {
	var req CreateAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_name is required"})
		return
	}

	accountID := uuid.New().String()

	// Create payload
	payload := model.AccountCreatedPayload{
		AccountID: accountID,
		UserName:  req.UserName,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to encode payload"})
		return
	}

	// Create event
	event := model.Event{
		ID:          uuid.New().String(),
		AggregateID: accountID,
		Type:        "AccountCreated",
		Payload:     string(payloadBytes),
		Timestamp:   time.Now(),
	}

	if err := kafka.PublishEvent(event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"account_id": accountID,
		"user_name":  req.UserName,
	})
}

func HandleDeposit(c *gin.Context) {
	accountID := c.Param("id")

	var req struct {
		Amount int64 `json:"amount"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	payload := map[string]any{
		"amount": req.Amount,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not encode payload"})
		return
	}

	event := model.Event{
		ID:          uuid.New().String(),
		AggregateID: accountID,
		Type:        "Deposited",
		Payload:     string(jsonPayload),
		Timestamp:   time.Now(),
	}

	if err := kafka.PublishEvent(event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to publish event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Deposit event published"})
}

type WithdrawPayload struct {
	Amount int64 `json:"amount"`
}

func HandleWithdraw(c *gin.Context) {
	accountID := c.Param("id")
	var body struct {
		Amount int64 `json:"amount"`
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "invalid payload"})
		return
	}

	// 🔍 Check current balance
	balance, err := cache.RDB.Get(context.Background(), "balance:"+accountID).Int64()
	if err != nil && err != redis.Nil {
		c.JSON(500, gin.H{"error": "could not fetch balance"})
		return
	}

	if balance < body.Amount {
		c.JSON(400, gin.H{"error": "insufficient balance"})
		return
	}

	payload := model.TransactionPayload{
		AccountID: accountID,
		Amount:    body.Amount,
	}
	jsonPayload, _ := json.Marshal(payload)

	event := model.Event{
		ID:          uuid.New().String(),
		AggregateID: accountID,
		Type:        "Withdrawn",
		Payload:     string(jsonPayload),
		Timestamp:   time.Now(),
	}

	if err := kafka.PublishEvent(event); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": "withdrawal event published"})
}
