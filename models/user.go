package models

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  []byte `json:"-"`
	Phone     string `json:"phone"`
}

func (user *User) SetPassword(password string) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Println("Error hashing password:", err)
		return
	}
	user.Password = hashedPassword
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
