package models

import "time"

type Team struct {
	ID      *int       `db:"id" json:"id"`
	Name    *string    `db:"name" json:"name"`
	Avatar  *string    `db:"avatar" json:"avatar"`
	Created *time.Time `db:"created" json:"created"`
	Updated *time.Time `db:"updated" json:"updated"`
}
