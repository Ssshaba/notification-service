package routes

import (
	"github.com/Ssshaba/notification-service/internal/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	notificationController := controllers.NewNotificationController()

	r.POST("/notifications", notificationController.Create)
	r.GET("/notifications", notificationController.List)
}
