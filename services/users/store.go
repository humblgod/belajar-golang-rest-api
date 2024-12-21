package users

import (
	"database/sql"
	"fmt"

	"github.com/humblgod/belajar-golang-rest-api/types"
)

type Store struct {
	db *sql.DB
}

// injection
func NewUserStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	user := new(types.User)

	// passing parameter with $1..$n
	if err := s.db.QueryRow("SELECT * FROM users WHERE email = $1", email).Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("query database error")
	}

	return user, nil 
} 


func (s *Store) CreateUser(user types.User) (error) {
	query := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3)"
	_, err := s.db.Exec(query, user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}