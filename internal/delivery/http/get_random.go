package http

import (
	"context"
	"fmt"
	"html/template"
	"net/http"

	"go.uber.org/zap"

	"github.ru/noskov-sergey/what_to_watch_golang/internal/metrics"
)

func (r *router) getRandomHandler(w http.ResponseWriter, req *http.Request) {
	var mtr = metrics.Met{Handler: metrics.GetRandomHandler}
	log := r.log.With(
		zap.String("method", req.Method),
	)

	log.Info("start get_random method")

	opinion, err := r.usecase.GetRandom(context.Background())
	if err != nil {
		log.Error("failed to get random:", zap.Error(err))

		mtr.Err = err
		r.met.Add(mtr)

		http.Redirect(w, req, fmt.Sprintf("%s%s%s", "http://", req.Host, "/500"), http.StatusSeeOther)
		return
	}

	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Error("failed to create template", zap.Error(err))

		mtr.Err = err
		r.met.Add(mtr)

		http.Redirect(w, req, fmt.Sprintf("%s%s%s", "http://", req.Host, "/500"), http.StatusSeeOther)
	}

	err = t.Execute(w, *opinion)
	if err != nil {
		log.Error("failed to parse template", zap.Error(err))

		mtr.Err = err
		r.met.Add(mtr)

		http.Redirect(w, req, fmt.Sprintf("%s%s%s", "http://", req.Host, "/500"), http.StatusSeeOther)
		return
	}

	log.Info("success get_random method")
	r.met.Add(mtr)
}
