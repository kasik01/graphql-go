package user

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             string `gorm:"type:uuid;primaryKey" json:"id"`
	Username       string `gorm:"type:text" json:"username"`
	HashedPassword string `gorm:"type:text"`
}

func NewUser(request RegisterUserRequest) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.PassWord), 10)
	if err != nil {
		return nil, err
	}

	id := uuid.New().String()
	return &User{ID: id, Username: request.UserName, HashedPassword: string(hashedPassword)}, nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(providedPassword))
	return err
}
