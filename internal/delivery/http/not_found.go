package http

import "net/http"

func (r *router) notFoundHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "templates/errors/404.html")
}
