package controllers

import (
	"net/http"

	"github.com/Ssshaba/notification-service/internal/models"
	"github.com/Ssshaba/notification-service/internal/service"
	"github.com/gin-gonic/gin"
)

type NotificationController struct{}

func NewNotificationController() *NotificationController {
	return &NotificationController{}
}

func (nc *NotificationController) Create(c *gin.Context) {
	var n models.Notification
	if err := c.ShouldBindJSON(&n); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.SaveNotificationDirect(&n); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save notification"})
		return
	}

	c.JSON(http.StatusCreated, n)
}

func (nc *NotificationController) List(c *gin.Context) {
	notifications, err := service.GetAllNotifications()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch notifications"})
		return
	}

	c.JSON(http.StatusOK, notifications)
}
