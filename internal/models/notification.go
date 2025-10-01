package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Notification struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Sender    string    `gorm:"size:255;not null" json:"sender" validate:"required,email"`
	Recipient string    `gorm:"size:255;not null" json:"recipient" validate:"required,email"`
	Message   string    `gorm:"type:text;not null" json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

func (n *Notification) Validate() error {
	v := validator.New()
	return v.Struct(n)
}
