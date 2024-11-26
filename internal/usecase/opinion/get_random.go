package opinion

import (
	"context"
	"fmt"

	"github.ru/noskov-sergey/what_to_watch_golang/internal/model"
)

func (u *UseCase) GetRandom(ctx context.Context) (*model.Opinion, error) {
	opinion, err := u.oRep.GetRandom(ctx)
	if err != nil {
		return nil, fmt.Errorf("repository get: %w", err)
	}

	return opinion, nil
}
