package character

import (
	"database/sql"

	"github.com/VoltealProductions/TheAzureArcchives/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateCharacter(char types.Character) error {
	_, err := s.db.Exec(
		"INSERT INTO characters (user_id, unique_id, firstname, lastname, faction, species, class) VALUES (?, ?, ?, ?, ?, ?, ?)",
		char.UserID, char.UniqueId, char.Firstname, char.Lastname, char.Faction, char.Species, char.Class,
	)
	if err != nil {
		return err
	}

	return nil
}
