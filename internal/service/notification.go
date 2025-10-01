package service

import (
	"github.com/Ssshaba/notification-service/internal/db"
	"github.com/Ssshaba/notification-service/internal/models"
)

// SaveNotificationDirect сохраняет уведомление напрямую в БД
func SaveNotificationDirect(n *models.Notification) error {
	return db.DB.Create(n).Error
}

// GetAllNotifications возвращает список всех уведомлений
func GetAllNotifications() ([]models.Notification, error) {
	var notifications []models.Notification
	err := db.DB.Find(&notifications).Error
	return notifications, err
}
