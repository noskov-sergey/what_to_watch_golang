package http

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"github.ru/noskov-sergey/what_to_watch_golang/internal/metrics"
	"github.ru/noskov-sergey/what_to_watch_golang/internal/model"
)

type Usecase interface {
	GetRandom(ctx context.Context) (*model.Opinion, error)
	Create(ctx context.Context, opinion model.Opinion) (int, error)
	GetById(ctx context.Context, id int) (*model.Opinion, error)
}

type router struct {
	met *metrics.Metrics

	log *zap.Logger

	chi.Router
	usecase Usecase
}

func New(usecase Usecase, met *metrics.Metrics, log *zap.Logger) *router {
	r := &router{
		Router:  chi.NewRouter(),
		usecase: usecase,
		met:     met,
		log:     log,
	}

	r.Get("/", r.getRandomHandler)
	r.Get("/add", r.NewOpinionHandler)
	r.Post("/add", r.createOpinionHandler)
	r.Route("/opinions", func(oR chi.Router) {
		oR.Get("/{opinionID}", r.getOpinionHandler)
	})
	r.Get("/404", r.notFoundHandler)
	r.Get("/500", r.internalServerErrorHandler)

	r.Handle("/static/*", http.FileServer(http.Dir("templates/")))

	return r
}
