package dto

import "time"

type CreateUserDTO struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type UpdateUserDTO struct {
	ID       int    `json:"id"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type UserResponseDTO struct {
	ID        int       `json:"id"`
	UserName  string    `json:"userName"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
