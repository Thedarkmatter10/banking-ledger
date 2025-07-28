package main

import (
	"github.com/Thedarkmatter10/ledger-service/kafka"
	"github.com/Thedarkmatter10/ledger-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	go kafka.StartConsumer() // Run Kafka consumer in background

	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
