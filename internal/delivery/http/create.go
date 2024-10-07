package http

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.ru/noskov-sergey/what_to_watch_golang/internal/model"
)

func (r *router) createOpinionHandler(w http.ResponseWriter, req *http.Request) {
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
		return
	}

	http.Redirect(w, req, fmt.Sprintf("%s%s%s/%s", "http://", req.Host, "/opinions", strconv.Itoa(id)), http.StatusSeeOther)
}

func (r *router) NewOpinionHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "templates/add.html")
}
