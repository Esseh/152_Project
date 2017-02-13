package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Multiplexer Function for CORE
func Handle_CORE(r *httprouter.Router) {
	r.GET("/", index)
}

// Serves the index page.
func index(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	ctx := NewContext(res,req)
	ServeTemplateWithParams(res, "index", struct {
		HeaderData
	}{
		*MakeHeader(ctx),
	})
}
