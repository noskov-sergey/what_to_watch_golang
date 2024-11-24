package http

import "net/http"

func (r *router) internalServerErrorHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "templates/errors/500.html")
}
