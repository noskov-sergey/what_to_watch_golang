package http

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.ru/noskov-sergey/what_to_watch_golang/internal/metrics"
)

func (r *router) getRandomHandler(w http.ResponseWriter, req *http.Request) {
	var mtr = metrics.Met{Handler: metrics.GetRandomHandler}

	opinion, err := r.usecase.Get(context.Background())
	if err != nil {
		mtr.Err = err
		r.met.Add(mtr)
		http.Redirect(w, req, fmt.Sprintf("%s%s%s", "http://", req.Host, "/500"), http.StatusSeeOther)
		return
	}

	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		mtr.Err = err
		r.met.Add(mtr)
		http.Redirect(w, req, fmt.Sprintf("%s%s%s", "http://", req.Host, "/500"), http.StatusSeeOther)
		log.Fatalf("failed to create template: %s", err)
	}
	err = t.Execute(w, *opinion)
	if err != nil {
		mtr.Err = err
		r.met.Add(mtr)
		http.Redirect(w, req, fmt.Sprintf("%s%s%s", "http://", req.Host, "/500"), http.StatusSeeOther)
		return
	}
	r.met.Add(mtr)
}
