package types

import (
	"time"
)

type CharacterStore interface {
	GetCharacterByUniqueId(urlId string) (*Character, error)
	GetCharacterByUserId(id int) ([]Character, error)
	CreateCharacter(Character) error
	UpdateCharacter(id int, user Character) error
	DeleteCharacter(id int) error
}

type Character struct {
	ID         int64     `json:"id"`
	UserID     int64     `json:"user_id"`
	UniqueId   string    `json:"unique_id_number"`
	Firstname  string    `json:"firstname"`
	Lastname   string    `json:"lastname"`
	Faction    string    `json:"faction"`
	Species    string    `json:"species"`
	Class      string    `json:"class"`
	ShortTitle string    `json:"short_title"`
	FullTitle  string    `json:"full_title"`
	Age        int64     `json:"age"`
	Gender     string    `json:"gender"`
	Pronouns   string    `json:"pronouns"`
	Height     float64   `json:"height"`
	Weight     float64   `json:"weight"`
	Birthplace string    `json:"birthplace"`
	Residence  string    `json:"residence"`
	About      string    `json:"about"`
	History    string    `json:"history"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CreateCharacterPayload struct {
	UserID    int64  `json:"user_id"`
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"required"`
	Faction   string `json:"faction" validate:"required"`
	Species   string `json:"species"`
	Class     string `json:"class" validate:"required"`
}

type UpdateCharacterPayload struct {
	Firstname  string  `json:"firstname"`
	Lastname   string  `json:"lastname"`
	Faction    string  `json:"faction"`
	Species    string  `json:"species"`
	Class      string  `json:"class"`
	ShortTitle string  `json:"short_title"`
	FullTitle  string  `json:"full_title"`
	Age        int64   `json:"age"`
	Gender     string  `json:"gender"`
	Pronouns   string  `json:"pronouns"`
	Height     float64 `json:"height"`
	Weight     float64 `json:"weight"`
	Birthplace string  `json:"birthplace"`
	Residence  string  `json:"residence"`
	About      string  `json:"about"`
	History    string  `json:"history"`
}
