package dto

import (
	"errors"
	"rest-api/src/core/domain"
)

type CreateUserDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func (c *CreateUserDTO) Validate() error {
	if c.Name == "" {
		return errors.New("name is required")
	}
	if c.Email == "" {
		return errors.New("email is required")
	}
	if c.Age <= 0 {
		return errors.New("age is required")
	}
	return nil
}

type UserResponseDTO struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func FromDomain(user domain.User) UserResponseDTO {
	return UserResponseDTO{
		ID:        user.Id(),
		Name:      user.Name(),
		Email:     user.Email(),
		Age:       user.Age(),
		CreatedAt: user.CreatedAt().Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt: user.UpdatedAt().Format("2006-01-02T15:04:05Z07:00"),
	}
}
