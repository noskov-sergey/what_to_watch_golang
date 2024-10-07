package opinion

import (
	"context"
	"fmt"

	"github.ru/noskov-sergey/what_to_watch_golang/internal/model"
)

func (u *UseCase) GetById(ctx context.Context, id int) (*model.Opinion, error) {
	opinion, err := u.oRep.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("repository get_by_id: %w", err)
	}

	return opinion, nil
}
