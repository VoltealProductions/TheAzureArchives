package types

import "time"

type UserStore interface {
	GetUserByUsername(username string) (*User, error)
	GetUserById(id int) (*User, error)
	CreateUser(User) error
	UpdateUser(userId int, user User) error
	DeleteUser(userId int) error
}

type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Public    bool      `json:"public"`
	Banned    bool      `json:"banned"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RegisterPayload struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=5,max=75"`
	Public   bool   `json:"public"`
}

type UpdatePayload struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
	Public   bool   `json:"public"`
}

type LoginPayload struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type CharacterStore interface {
	// GetCharacterByName(name string) (*Character, error)
	// GetCharacterById(id int) (*Character, error)
	CreateCharacter(Character) error
	// UpdateCharacter(id int, user Character) error
	// DeleteCharacter(id int) error
}

type Character struct {
	ID             int64     `json:"id"`
	UniqueIdNumber int64     `json:"unique_id_number"`
	Firstname      string    `json:"firstname"`
	Lastname       string    `json:"lastname"`
	Slug           string    `json:"slug"`
	Faction        string    `json:"faction"`
	Class          string    `json:"class"`
	ShortTitle     string    `json:"short_title"`
	FullTitle      string    `json:"full_title"`
	Age            int32     `json:"age"`
	Gender         string    `json:"gender"`
	Pronouns       string    `json:"pronouns"`
	Height         float32   `json:"height"`
	Weight         float32   `json:"weight"`
	Species        string    `json:"species"`
	Birthplace     string    `json:"birthplace"`
	Residence      string    `json:"residence"`
	About          string    `json:"about"`
	History        string    `json:"history"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type CreateCharacterPayload struct {
}

type UpdateCharacterPayload struct {
}
