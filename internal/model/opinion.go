package model

import (
	"time"
)

type Opinion struct {
	ID        *int       `db:"id"`
	Title     *string    `db:"title"`
	Text      *string    `db:"text"`
	Source    *string    `db:"source"`
	AddedBy   *string    `db:"added_by"`
	CreatedAt *time.Time `db:"created_at"`
}
