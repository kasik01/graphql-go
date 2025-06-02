package user

import (
	"graphql-hasura-demo/internal/database"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func (r *Repository) Save(user *User) (*User, error) {
	err := r.db.Save(&user).Error
	return user, err
}

func (r *Repository) FindByUsername(username string) (*User, error) {
	var user User
	if err := r.db.Where(&User{Username: username}).First(&user).Error; err != nil {
		return nil, &Errors.NotFound
	}
	return &user, nil
}

func NewRepository() *Repository {
	database.Connect()
	db := database.GetDB()
	repository := Repository{db}
	db.AutoMigrate(&User{})

	return &repository
}
