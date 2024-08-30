package guild

import (
	"database/sql"
	"fmt"

	"github.com/VoltealProductions/TheAzureArcchives/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetGuildBySlug(slug string) (*types.Guild, error) {
	rows, err := s.db.Query("SELECT * FROM guilds WHERE slug = ?", slug)
	if err != nil {
		return nil, err
	}

	g := new(types.Guild)
	for rows.Next() {
		g, err = scanRowIntoGuild(rows)
		if err != nil {
			return nil, err
		}
	}

	if g.ID == 0 {
		return nil, fmt.Errorf("guild not found")
	}

	return g, nil
}

func scanRowIntoGuild(rows *sql.Rows) (*types.Guild, error) {
	guild := new(types.Guild)

	err := rows.Scan(
		&guild.ID,
		&guild.OwnerId,
		&guild.Slug,
		&guild.Name,
		&guild.Faction,
		&guild.Realm,
		&guild.Ranks,
		&guild.Recruiting,
		&guild.Description,
		&guild.CreatedAt,
		&guild.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return guild, nil
}

func (s *Store) CreateGuild(guild types.Guild) error {
	_, err := s.db.Exec(
		"INSERT INTO guilds (owner_id, slug, name, faction, realm, recruiting, description) VALUES (?, ?, ?, ?, ?, ?, ?)",
		guild.OwnerId, guild.Slug, guild.Name, guild.Faction, guild.Realm, guild.Recruiting, guild.Description,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdateGuild(slug string, guild types.Guild) error {

	_, err := s.db.Exec(
		"UPDATE guilds SET slug = ?, name = ?, faction = ?, realm = ?, ranks = ?, recruiting = ?, description = ?, updated_at = ? WHERE slug = ?",
		guild.Slug, guild.Name, guild.Faction, guild.Realm, guild.Ranks, guild.Recruiting, guild.Description, guild.UpdatedAt, slug,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) DeleteGuild(slug string) error {
	_, err := s.db.Exec(
		"DELETE FROM guilds WHERE slug = ?",
		slug,
	)
	if err != nil {
		return err
	}

	return nil
}
