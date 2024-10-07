package http

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (r *router) getOpinionHandler(w http.ResponseWriter, req *http.Request) {
	articleID := chi.URLParam(req, "opinionID")
	id, err := strconv.Atoi(articleID)
	if err != nil {
		http.Redirect(w, req, fmt.Sprintf("%s%s%s", "http://", req.Host, "/404"), http.StatusSeeOther)
		return
	}

	opinion, err := r.usecase.GetById(context.Background(), id)
	if err != nil {
		http.Redirect(w, req, fmt.Sprintf("%s%s%s", "http://", req.Host, "/404"), http.StatusSeeOther)
		return
	}

	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Redirect(w, req, fmt.Sprintf("%s%s%s", "http://", req.Host, "/500"), http.StatusSeeOther)
		log.Fatalf("failed to create template: %s", err)
	}
	err = t.Execute(w, *opinion)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
