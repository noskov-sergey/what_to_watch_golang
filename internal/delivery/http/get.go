package http

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"github.ru/noskov-sergey/what_to_watch_golang/internal/metrics"
)

func (r *router) getOpinionHandler(w http.ResponseWriter, req *http.Request) {
	var mtr = metrics.Met{Handler: metrics.GetHandler}
	log := r.log.With(
		zap.String("method", req.Method),
	)

	articleID := chi.URLParam(req, "opinionID")
	id, err := strconv.Atoi(articleID)
	if err != nil {
		log.Error("failed to convert:", zap.Error(err))

		mtr.Err = err
		r.met.Add(mtr)

		http.Redirect(w, req, fmt.Sprintf("%s%s%s", "http://", req.Host, "/404"), http.StatusSeeOther)
		return
	}

	opinion, err := r.usecase.GetById(context.Background(), id)
	if err != nil {
		log.Error("failed to get by id:", zap.Error(err))

		mtr.Err = err
		r.met.Add(mtr)

		http.Redirect(w, req, fmt.Sprintf("%s%s%s", "http://", req.Host, "/404"), http.StatusSeeOther)
		return
	}

	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal("failed to create template: %s", zap.Error(err))

		mtr.Err = err
		r.met.Add(mtr)

		http.Redirect(w, req, fmt.Sprintf("%s%s%s", "http://", req.Host, "/500"), http.StatusSeeOther)
		return
	}
	err = t.Execute(w, *opinion)
	if err != nil {
		log.Error("failed to execute:", zap.Error(err))

		mtr.Err = err
		r.met.Add(mtr)

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	r.met.Add(mtr)
}
