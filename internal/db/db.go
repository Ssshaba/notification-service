package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Ssshaba/notification-service/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var err error

	for i := 0; i < 10; i++ {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("Не удалось подключиться к БД... (%d/10)", i+1)
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		log.Fatalf("Ошибка подключения к БД после 10 попыток: %v", err)
	}

	if err := DB.AutoMigrate(&models.Notification{}); err != nil {
		log.Fatalf("❌ Ошибка миграции: %v", err)
	}

	log.Println("БД подключена")
}
