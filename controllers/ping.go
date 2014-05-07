package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/zenazn/goji/web"
)

func AddRoutes(m *web.Mux) {
	m.Get("/ping", ping)
	m.Get("/", index)
	fmt.Print("Added index")
}

func ping(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "1.0.0")
}

func index(c web.C, w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("./assets/index.html")

	if err != nil {
		w.WriteHeader(404)
		return
	}

	w.Write(content)
}
