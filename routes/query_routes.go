package routes

import (
	"github.com/Thedarkmatter10/ledger-service/middleware"
	"github.com/Thedarkmatter10/ledger-service/query"

	"github.com/gin-gonic/gin"
)

func QueryRoutes(router *gin.RouterGroup) {

	accounts := router.Group("/accounts", middleware.AccountExistsMiddleware())
	{
		accounts.GET("/:id/balance", query.GetBalance)
		accounts.GET("/:id/transactions", query.GetTransactionHistory)

	}

}
