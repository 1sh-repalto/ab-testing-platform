package models

type User struct {
	ID			string	`db:"id" json:"id"`
	Name		string	`db:"name" json:"name"`
	Email		string	`db:"email" json:"email"`
	Password	string	`db:"password,omitempty"`
	CreatedAt	string	`db:"created_at" json:"created_at"`
}