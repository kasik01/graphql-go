package notification

import (
	"github.com/google/uuid"
)

type Notification struct {
	ID      string `gorm:"type:uuid;primaryKey" json:"id"`
	Message string `gorm:"type:text" json:"message"`
	UserId  uint   `json:"user_id"`
}

func (Notification) TableName() string {
	return "notification"
}

func NewNotification(message string, userId uint) *Notification {
	id := uuid.New().String()
	return &Notification{ID: id, Message: message, UserId: userId}
}
