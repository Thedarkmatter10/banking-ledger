package middleware

import (
	"net/http"

	"github.com/Thedarkmatter10/ledger-service/repository"

	"github.com/gin-gonic/gin"
)

func AccountExistsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		accountID := c.Param("id")
		exists, err := repository.AccountExists(accountID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not validate account"})
			c.Abort()
			return
		}
		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": "account does not exist"})
			c.Abort()
			return
		}
		c.Next()
	}
}
