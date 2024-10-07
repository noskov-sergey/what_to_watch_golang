package opinion

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"github.ru/noskov-sergey/what_to_watch_golang/internal/model"
)

func (r *Repository) Get(ctx context.Context) ([]*model.Opinion, error) {
	builder := sq.Select(
		idColumn,
		titleColumn,
		textColumn,
		sourceColumn,
		addedByColumn,
		createdAtColumn,
	).From(tableName)

	sqlQuery, _, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("to sql: %w", err)
	}

	rows, err := r.db.QueryContext(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var opinions []*model.Opinion

	for rows.Next() {
		var o model.Opinion
		err = rows.Scan(&o.ID, &o.Title, &o.Text, &o.Source, &o.AddedBy, &o.CreatedAt)
		if err != nil {
			return nil, err
		}

		opinions = append(opinions, &o)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return opinions, nil
}
