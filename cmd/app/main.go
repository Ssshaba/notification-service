package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/Ssshaba/notification-service/internal/db"
	"github.com/Ssshaba/notification-service/internal/rabbit"
	"github.com/Ssshaba/notification-service/internal/routes"
)

func main() {
	_ = godotenv.Load()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8081"
	}

	db.Connect()

	rabbit.InitRabbit(os.Getenv("RABBIT_URL"))

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pdddong"})
	})

	routes.SetupRoutes(r)

	log.Printf("Starting server on :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
