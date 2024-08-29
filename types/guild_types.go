package types

type GuildStore interface {
	// GetGuildByUniqueId(urlId string) (*Guild, error)
	// GetGuildByUserId(id int) ([]Guild, error)
	// CreateGuild(Guild) error
	// UpdateGuild(id int, user Guild) error
	// DeleteCharacter(id int) error
}

type Guild struct {
}

type CreateGuildPayload struct {
}

type UpdateGuildPayload struct {
}
