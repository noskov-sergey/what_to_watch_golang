package http

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.ru/noskov-sergey/what_to_watch_golang/internal/metrics"
	"github.ru/noskov-sergey/what_to_watch_golang/internal/model"
)

func (r *router) createOpinionHandler(w http.ResponseWriter, req *http.Request) {
	var mtr = metrics.Met{Handler: metrics.CreateHandler}
	ctx := context.Context(context.Background())

	title := req.FormValue("title")
	text := req.FormValue("text")
	source := req.FormValue("source")

	opinion := model.Opinion{
		Title:  &title,
		Text:   &text,
		Source: &source,
	}

	id, err := r.usecase.Create(ctx, opinion)
	if err != nil {
		http.Redirect(w, req, fmt.Sprintf("%s%s%s", "http://", req.Host, "/500"), http.StatusSeeOther)
		mtr.Err = err
		r.met.Add(mtr)
		return
	}

	http.Redirect(w, req, fmt.Sprintf("%s%s%s/%s", "http://", req.Host, "/opinions", strconv.Itoa(id)), http.StatusSeeOther)
	r.met.Add(mtr)
}

func (r *router) NewOpinionHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "templates/add.html")
}
