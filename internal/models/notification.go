package models

import "time"

type Notification struct {
    id        uint      `gorm:"primaryKey" json:"id"`
    sender    string    `gorm:"size:255;not null" json:"sender"`
    recipient string    `gorm:"size:255;not null" json:"recipient"`
    message   string    `gorm:"type:text;not null" json:"message"`
    CreatedAt time.Time `json:"created_at"`
}
