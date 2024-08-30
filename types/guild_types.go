package types

import (
	"time"
)

type GuildStore interface {
	GetGuildBySlug(slug string) (*Guild, error)
	GetGuildsByUserId(id int) ([]Guild, error)
	CreateGuild(Guild) error
	ConfirmThatGuildExists(slug string) (bool, error)
	UpdateGuild(slug string, guild Guild) error
	DeleteGuild(slug string) error
}

type Guild struct {
	ID          uint      `json:"id"`
	OwnerId     uint      `json:"owner_id"`
	Slug        string    `json:"slug"`
	Name        string    `json:"name"`
	Faction     string    `json:"faction"`
	Realm       string    `json:"realm"`
	Ranks       string    `json:"ranks"`
	Recruiting  bool      `json:"recruiting"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateGuildPayload struct {
	OwnerId     uint   `json:"owner_id"`
	Slug        string `json:"slug"`
	Name        string `json:"name" validate:"required"`
	Faction     string `json:"faction" validate:"required"`
	Realm       string `json:"realm" validate:"required"`
	Recruiting  bool   `json:"recruiting" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type UpdateGuildPayload struct {
	Slug        string    `json:"slug"`
	Name        string    `json:"name"`
	Faction     string    `json:"faction"`
	Realm       string    `json:"realm"`
	Ranks       string    `json:"ranks"`
	Recruiting  bool      `json:"recruiting"`
	Description string    `json:"description"`
	UpdatedAt   time.Time `json:"updated_at"`
}
