// routes/routes.go
package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	CommandRoutes(v1)

	QueryRoutes(v1)

}
