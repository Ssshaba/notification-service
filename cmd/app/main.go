package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/Ssshaba/notification-service/internal/db"
	"github.com/Ssshaba/notification-service/internal/models"
	"github.com/Ssshaba/notification-service/internal/rabbit"
	"github.com/Ssshaba/notification-service/internal/routes"
	"github.com/Ssshaba/notification-service/internal/service"
)

func main() {
	_ = godotenv.Load()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8081"
	}

	db.Connect()

	rabbit.InitRabbit(os.Getenv("RABBITMQ_URL"))

	rabbit.Consumer("notifications", func(msg []byte) error {
		var n models.Notification
		if err := json.Unmarshal(msg, &n); err != nil {
			return err
		}
		return service.SaveNotificationDirect(&n)
	})

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong!"})
	})

	routes.SetupRoutes(r)

	log.Printf("Starting server on :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
