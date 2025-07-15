package entity

import (
	"goexpert-api/pkg/entity"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"Email"`
	Password string    `json:"-"`
}

func NewUser(name, email, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	user := &User{
		ID:       entity.NewID(),
		Name:     name,
		Email:    email,
		Password: string(hash),
	}
	return user, nil
}

func (u *User) ValidatePassword(password string) bool {
	hasErro := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	return hasErro == nil
}
