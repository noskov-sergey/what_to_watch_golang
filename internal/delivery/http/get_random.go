package http

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func (r *router) getRandomHandler(w http.ResponseWriter, req *http.Request) {
	opinion, err := r.usecase.Get(context.Background())
	if err != nil {
		http.Redirect(w, req, fmt.Sprintf("%s%s%s", "http://", req.Host, "/500"), http.StatusSeeOther)
		return
	}

	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Redirect(w, req, fmt.Sprintf("%s%s%s", "http://", req.Host, "/500"), http.StatusSeeOther)
		log.Fatalf("failed to create template: %s", err)
	}
	err = t.Execute(w, *opinion)
	if err != nil {
		http.Redirect(w, req, fmt.Sprintf("%s%s%s", "http://", req.Host, "/500"), http.StatusSeeOther)
		return
	}
}
