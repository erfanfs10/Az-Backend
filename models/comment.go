package models

import "time"

type Comment struct {
	ID          *int       `db:"id" json:"id"`
	Name        *string    `db:"name" json:"name"`
	Avatar      *string    `db:"avatar" json:"avatar"`
	Title       *string    `db:"title" json:"title"`
	Description *string    `db:"description" json:"description"`
	Created     *time.Time `db:"created" json:"created"`
	Updated     *time.Time `db:"updated" json:"updated"`
}
