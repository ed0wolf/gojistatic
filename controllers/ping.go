package controllers

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji/web"
)

func AddRoutes(m *web.Mux) {
	m.Get("/ping", ping)
}

func ping(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "1.0.0")
}
