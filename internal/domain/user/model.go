package user

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	PwdHash string `json:"pwd_hash"`
}

func (u User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.PwdHash), []byte(password))
	if err != nil {
		return fmt.Errorf("password does not match")
	}

	return nil
}

type createUserDTO struct {
	Name           string `json:"name"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeatPassword"`
}

type updateUserDTO struct {
	Id          int    `json:"id"`
	Password    string `json:"password"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
