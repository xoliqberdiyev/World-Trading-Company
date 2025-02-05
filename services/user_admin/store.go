package user_admin

import (
	"database/sql"

	types_admin "github.com/XoliqberdiyevBehruz/wtc_backend/types/user_admin"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateUser(user *types_admin.UserCreatePayload) error {
	query := `INSERT INTO users(username, first_name, last_name, password) VALUES($1, $2, $3, $4)`
	_, err := s.db.Exec(query, &user.Username, &user.FirstName, &user.LastName, &user.Password)
	if err != nil {
		return nil
	}
	return nil
}

func (s *Store) GetUserById(id string) (*types_admin.UserDetailPayload, error) {
	var user types_admin.UserDetailPayload
	query := `SELECT id, username, first_name, last_name, created_at FROM users WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(&user.Id, &user.Username, &user.FirstName, &user.LastName, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (s *Store) GetUserByUsername(username string) (*types_admin.User, error) {
	var firstName *sql.NullString
	var lastName *sql.NullString
	var user types_admin.User
	query := `SELECT * FROM users WHERE username = $1`
	err := s.db.QueryRow(query, username).Scan(&user.Id, &user.Username, &firstName, &lastName, &user.Password, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (s *Store) DeleteUserById(id string) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := s.db.Query(query, id)
	if err != nil {
		return nil
	}
	return nil
}

func (s *Store) UpdateUserById(user *types_admin.UserUpdatePayload, id string) error {
	query := `UPDATE users SET username = $2, first_name = $3, last_name = $4 WHERE id = $1`
	_, err := s.db.Query(query, id, &user.Username, &user.FirstName, &user.LastName)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetAllUsers() ([]*types_admin.UserDetailPayload, error) {
	var users []*types_admin.UserDetailPayload
	query := `SELECT id, username, first_name, last_name, created_at FROM users`
	rows, err := s.db.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	for rows.Next() {
		var user types_admin.UserDetailPayload
		if err := rows.Scan(&user.Id, &user.Username, &user.FirstName, &user.LastName, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}
