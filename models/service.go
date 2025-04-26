package models

import "time"

type Service struct {
	ID          *int       `db:"id" json:"id"`
	Image       *string    `db:"image" json:"image"`
	Title       *string    `db:"title" json:"title"`
	Description *string    `db:"description" json:"description"`
	Created     *time.Time `db:"created" json:"created"`
	Updated     *time.Time `db:"updated" json:"updated"`
}
