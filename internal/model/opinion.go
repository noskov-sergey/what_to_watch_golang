package model

import (
	"time"
)

type Opinion struct {
	ID        *int       `db:"id"`
	Title     *string    `db:"title"`
	Text      *string    `db:"text"`
	Source    *string    `db:"source"`
	CreatedAt *time.Time `db:"created_at"`
}
