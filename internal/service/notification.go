package service

import (
	"encoding/json"

	"github.com/Ssshaba/notification-service/internal/db"
	"github.com/Ssshaba/notification-service/internal/models"
	"github.com/Ssshaba/notification-service/internal/rabbit"
)

func SaveNotificationDirect(n *models.Notification) error {
	if err := n.Validate(); err != nil {
		return err
	}
	return db.DB.Create(n).Error
}

func GetAllNotifications() ([]models.Notification, error) {
	var notifications []models.Notification
	err := db.DB.Find(&notifications).Error
	return notifications, err
}

func SendNotificationToQueue(n *models.Notification) error {
	body, err := json.Marshal(n)
	if err != nil {
		return err
	}

	return rabbit.Publish("notifications", body)
}
