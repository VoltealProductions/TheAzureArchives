package guild

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/VoltealProductions/TheAzureArcchives/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) ConfirmThatGuildExists(slug string) (bool, error) {
	var nui sql.Null[string]
	err := s.db.QueryRow("SELECT id FROM guilds WHERE slug = ?", slug).Scan(&nui)
	if err != nil {
		return false, err
	}

	if nui.Valid {
		return true, nil
	}

	return false, nil
}

func (s *Store) ConfirmThatUserOwnsGuild(slug string, id uint) (bool, error) {
	var nui sql.Null[string]
	err := s.db.QueryRow("SELECT id FROM guilds WHERE slug = ? AND owner_id = ?", slug, id).Scan(&nui)
	if err != nil {
		return false, err
	}

	if nui.Valid {
		return true, nil
	}

	return false, nil
}

func (s *Store) GetGuildsByUserId(id int) ([]types.Guild, error) {
	rows, err := s.db.Query("SELECT * FROM guilds WHERE owner_id = ?", id)
	if err != nil {
		return nil, err
	}

	guilds := []types.Guild{}
	for rows.Next() {
		guild, err := scanRowIntoGuild(rows)
		guilds = append(guilds, *guild)
		if err != nil {
			return nil, err
		}
	}

	return guilds, nil
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

func (s *Store) TransferGuild(slug string, id uint, guild types.Guild) error {
	if id == guild.OwnerId {
		return fmt.Errorf("you can not transfer a guild to yourself")
	}

	_, err := s.db.Exec(
		"UPDATE guilds SET owner_id = ? WHERE owner_id = ? AND slug = ?",
		guild.OwnerId, id, slug,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetAllGuildMembers(slug string) ([]types.Character, error) {
	rows, err := s.db.Query("SELECT * FROM guildmembers WHERE guild_slug = ?", slug)
	if err != nil {
		return nil, err
	}

	characters := []types.Character{}
	for rows.Next() {
		char := s.retrieveCharacters(rows)
		characters = append(characters, *char)
	}

	return characters, nil
}

func (s *Store) retrieveCharacters(rows *sql.Rows) *types.Character {
	id := new(types.GuildMember)
	err := rows.Scan(
		&id.ID,
		&id.GuildSlug,
		&id.CharacterID,
		&id.Rank,
		&id.CreatedAt,
		&id.UpdatedAt,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(id.CharacterID)

	characterRows, err := s.db.Query("SELECT * FROM characters WHERE id = ?", id.CharacterID)
	if err != nil {
		log.Fatal(err)
	}

	g := new(types.Character)
	for characterRows.Next() {
		g, err = scanRowIntoMember(characterRows)
		if err != nil {
			log.Fatal(err)
		}
	}

	return g
}

func scanRowIntoMember(rows *sql.Rows) (*types.Character, error) {
	c := new(types.Character)

	err := rows.Scan(
		&c.ID,
		&c.UserID,
		&c.UniqueId,
		&c.Firstname,
		&c.Lastname,
		&c.Faction,
		&c.Species,
		&c.Class,
		&c.ShortTitle,
		&c.FullTitle,
		&c.Age,
		&c.Gender,
		&c.Pronouns,
		&c.Height,
		&c.Weight,
		&c.Birthplace,
		&c.Residence,
		&c.About,
		&c.History,
		&c.CreatedAt,
		&c.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return c, nil
}
