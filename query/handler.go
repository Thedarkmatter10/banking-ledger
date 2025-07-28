package query

import (
	"context"
	"net/http"

	"github.com/Thedarkmatter10/ledger-service/cache"
	"github.com/Thedarkmatter10/ledger-service/repository"

	"github.com/gin-gonic/gin"
)

func GetBalance(c *gin.Context) {
	accountID := c.Param("id")

	balance, err := cache.RDB.Get(context.Background(), "balance:"+accountID).Int64()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch balance"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"account_id": accountID,
		"balance":    balance,
	})
}

func GetTransactionHistory(c *gin.Context) {
	accountID := c.Param("id")
	events, err := repository.GetEventsByAggregateID(accountID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, events)
}
