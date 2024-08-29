package types

import (
	"database/sql"
	"time"
)

type UserStore interface {
	GetUserByUsername(username string) (*User, error)
	GetUserById(id int) (*User, error)
	CreateUser(User) error
	UpdateUser(userId int, user User) error
	DeleteUser(userId int) error
}

type User struct {
	ID        int64          `json:"id"`
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	Public    bool           `json:"public"`
	Banned    bool           `json:"banned"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	Badges    sql.NullString `json:"badges"`
}

type RegisterPayload struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=5,max=75"`
	Public   bool   `json:"public"`
	Badges   string `json:"badges"`
}

type UpdatePayload struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
	Public   bool   `json:"public"`
	Badges   string `json:"badges"`
}

type LoginPayload struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
