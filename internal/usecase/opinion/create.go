package opinion

import (
	"context"
	"fmt"

	"github.ru/noskov-sergey/what_to_watch_golang/internal/model"
)

func (u *UseCase) Create(ctx context.Context, opinion model.Opinion) (int, error) {
	id, err := u.oRep.Create(ctx, opinion)
	if err != nil {
		return 0, fmt.Errorf("usecase create: %w", err)
	}

	return id, nil
}
