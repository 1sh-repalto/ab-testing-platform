package models

import "time"

type Variant struct {
	ID				int			`db:"id" json:"id"`
	Name			string		`db:"id" json:"name"`
	Weight			float64		`db:"id" json:"weight"`
	ExperimentID	int			`db:"experiment_id" json:"experiment_id"`
	CreatedAt		time.Time	`db:"created_at" json:"created_at"`
}