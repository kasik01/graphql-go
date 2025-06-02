package notification

import (
	"graphql-hasura-demo/internal/database"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func (r *repository) save(notification *Notification) (*Notification, error) {
	err := r.db.Save(&notification).Error
	return notification, err
}

func NewRepository() *repository {
	database.Connect()
	db := database.GetDB()
	repository := repository{db}

	return &repository
}
