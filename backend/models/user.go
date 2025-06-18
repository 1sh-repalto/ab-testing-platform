package models

import (
	"time"
	"backend/db"
)

type User struct {
	ID        string    `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"password,omitempty" json:"-"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

// Helper Functions
func GetUserByID(id int) (*User, error) {
	var user User
	query := `SELECT id, name, email, created_at FROM users WHERE id = $1`
	err := db.DB.Get(&user, query, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	query := `SELECT * FROM users WHERE email = $1`
	err := db.DB.Get(&user, query, email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}