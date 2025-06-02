package notification

import (
	"github.com/google/uuid"
)

type Notification struct {
	ID      string `gorm:"type:uuid;primaryKey" json:"id"`
	Message string `gorm:"type:text" json:"messgae"`
	UserId  string `gorm:"type:text" json:"user_id"`
}

func (Notification) TableName() string {
	return "notification"
}

func NewNotification(messgae string, userId string) *Notification {
	id := uuid.New().String()
	return &Notification{ID: id, Message: messgae, UserId: userId}
}
