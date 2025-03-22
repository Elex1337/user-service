package entity

import "time"

type User struct {
	ID        int       `json:"id" db:"id"`
	UserName  string    `json:"userName" db:"user_name"`
	Password  string    `json:"-" db:"password"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}
