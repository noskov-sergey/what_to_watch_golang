package opinion

import (
	"context"

	"github.ru/noskov-sergey/what_to_watch_golang/internal/model"
)

type OpinionRepository interface {
	GetRandom(ctx context.Context) (*model.Opinion, error)
	Create(ctx context.Context, opinion model.Opinion) (int, error)
	GetById(ctx context.Context, id int) (*model.Opinion, error)
}

type UseCase struct {
	oRep OpinionRepository
}

func New(oRep OpinionRepository) *UseCase {
	return &UseCase{
		oRep: oRep,
	}
}
