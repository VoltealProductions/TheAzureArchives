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

func (s *Store) GetCharacterByUniqueId(id string) (*types.Character, error) {
	rows, err := s.db.Query("SELECT * FROM characters WHERE unique_id = ? LIMIT 1", id)
	if err != nil {
		return nil, err
	}

	c := new(types.Character)
	for rows.Next() {
		c, err = scanRowIntoCharacter(rows)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

func scanRowIntoCharacter(rows *sql.Rows) (*types.Character, error) {
	char := new(types.Character)

	err := rows.Scan(
		&char.ID,
		&char.UserID,
		&char.UniqueId,
		&char.Firstname,
		&char.Lastname,
		&char.Faction,
		&char.Class,
		&char.Species,
		&char.ShortTitle,
		&char.FullTitle,
		&char.Age,
		&char.Gender,
		&char.Pronouns,
		&char.Height,
		&char.Weight,
		&char.Birthplace,
		&char.Residence,
		&char.About,
		&char.History,
		&char.CreatedAt,
		&char.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return char, nil
}
