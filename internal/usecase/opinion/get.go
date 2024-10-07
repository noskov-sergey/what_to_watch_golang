package opinion

import (
	"context"
	"fmt"
	"math/rand/v2"

	"github.ru/noskov-sergey/what_to_watch_golang/internal/model"
)

func (u *UseCase) Get(ctx context.Context) (*model.Opinion, error) {
	opinions, err := u.oRep.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("repository get: %w", err)
	}

	if len(opinions) == 0 {
		return &model.Opinion{}, fmt.Errorf("repository get: %s", "no opinions")
	}

	randOpinion := opinions[rand.IntN(len(opinions))]

	return randOpinion, nil
}
