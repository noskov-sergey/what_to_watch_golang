package opinion

import (
	"context"
	"fmt"

	"github.ru/noskov-sergey/what_to_watch_golang/internal/model"
)

func (r *Repository) GetRandom(ctx context.Context) (*model.Opinion, error) {
	query, err := r.db.Prepare(`SELECT id, title, text, source FROM opinions ORDER BY RANDOM() LIMIT 1`)
	if err != nil {
		return nil, err
	}

	var o model.Opinion

	err = query.QueryRow().Scan(&o.ID, &o.Title, &o.Text, &o.Source)
	if err != nil {
		return nil, fmt.Errorf("row scan: %w", err)
	}

	return &o, nil
}
