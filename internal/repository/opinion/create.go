package opinion

import (
	"context"
	"fmt"

	"github.ru/noskov-sergey/what_to_watch_golang/internal/model"
)

func (r *Repository) Create(ctx context.Context, opinion model.Opinion) (int, error) {
	row := r.db.QueryRow(`
		INSERT INTO opinions (title,text,source)
		VALUES ($1, $2, $3) RETURNING id`,
		opinion.Title, opinion.Text, opinion.Source)

	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("repository create: %w", err)
	}

	return id, nil
}
