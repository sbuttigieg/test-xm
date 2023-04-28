package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

const cost = 14

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name" binding:"required,max=30"`
	Username  string    `json:"username" binding:"max=30"`
	Email     string    `json:"email"  binding:"required,max=50"`
	Password  string    `json:"password" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return err
	}

	user.Password = string(bytes)

	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}

	return nil
}
