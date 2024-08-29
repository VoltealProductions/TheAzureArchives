package user

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/VoltealProductions/TheAzureArcchives/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByUsername(username string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE username = ?", username)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.Public,
		&user.Banned,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Badges,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Store) GetUserById(id int) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE id = ? LIMIT 1", id)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func (s *Store) CreateUser(user types.User) error {
	_, err := s.db.Exec("INSERT INTO users (username, email, password, public) VALUES (?, ?, ?, ?)", strings.ToLower(user.Username), user.Email, user.Password, user.Public)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdateUser(userId int, user types.User) error {

	_, err := s.db.Exec(
		"UPDATE users SET username = ?, password = ?, email = ?, public = ? WHERE id = ?",
		strings.ToLower(user.Username), user.Password, user.Email, user.Public, userId,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) DeleteUser(userId int) error {
	_, err := s.db.Exec(
		"DELETE FROM users WHERE id = ?",
		userId,
	)
	if err != nil {
		return err
	}

	return nil
}
