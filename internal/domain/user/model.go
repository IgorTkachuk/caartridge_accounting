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

func (u *User) GeneratePasswordHash() error {
	pwd, err := generatePasswordHash(u.PwdHash)
	if err != nil {
		return err
	}
	u.PwdHash = pwd
	return nil
}

func generatePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password due to error %w", err)
	}

	return string(hash), nil
}

type CreateUserDTO struct {
	Name           string `json:"name"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
}

type UpdateUserDTO struct {
	Id          int    `json:"id"`
	Password    string `json:"password"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type SignInUserDTO struct {
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
}

func NewUser(dto CreateUserDTO) User {
	return User{
		Name:    dto.Name,
		PwdHash: dto.Password,
	}
}
