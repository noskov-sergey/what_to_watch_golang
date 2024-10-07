package opinion

import (
	"database/sql"
)

const (
	tableName       = "opinions"
	idColumn        = "id"
	titleColumn     = "title"
	textColumn      = "text"
	sourceColumn    = "source"
	createdAtColumn = "created_at"
	addedByColumn   = "added_by"
)

type Repository struct {
	db *sql.DB
}

func NewOpinionRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}
