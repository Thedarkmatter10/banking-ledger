package routes

import (
	"github.com/Thedarkmatter10/ledger-service/command"
	"github.com/Thedarkmatter10/ledger-service/middleware"

	"github.com/gin-gonic/gin"
)

// CommandRoutes sets up all command-related routes.
func CommandRoutes(router *gin.RouterGroup) {
	// Register account creation route
	router.POST("/accounts", command.CreateAccount)

	// Group routes that require account existence check
	registerAccountTransactionRoutes(router)
}

// registerAccountTransactionRoutes encapsulates all routes that need account existence validation
func registerAccountTransactionRoutes(router *gin.RouterGroup) {
	accounts := router.Group("/accounts", middleware.AccountExistsMiddleware())
	{
		accounts.POST("/:id/deposit", command.HandleDeposit)
		accounts.POST("/:id/withdraw", command.HandleWithdraw)
	}
}
