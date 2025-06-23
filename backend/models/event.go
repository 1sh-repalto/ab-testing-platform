package models

import "time"

type Event struct {
	ID				int			`db:"id" json:"id"`
	ExperimentID	int			`db:"experiment_id" json:"experiment_id"`
	VariantID		int			`db:"variant_id" json:"variant_id"`
	UserIdentifier	string		`db:"user_identifier" json:"user_identifier"`
	EventType		string 		`db:"event_type" json:"event_type"`	
	CreatedAt		time.Time	`db:"created_at" json:"created_at"`
}
