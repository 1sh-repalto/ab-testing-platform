package models

import "time"

type Experiment struct {
	ID				int 		`db:"id" json:"id"`
	Name			string		`db:"name" json:"name"`
	Description		string		`db:"description" json:"description"`
	Status			string		`db:"status" json:"status"`
	CreatedBy		int			`db:"created_by" json:"created_by"`
	CreatedAt		time.Time	`db:"created_at" json:"created_at"`
}