package controllers

import (
	"github.com/zenazn/goji/web"

	"github.com/ed0wolf/gojistatic/controllers/index"
	"github.com/ed0wolf/gojistatic/controllers/ping"
)

func AddDefaultControllers(indexPath string) func(*web.Mux) {
	return func(m *web.Mux) {
		m.Get("/ping", ping.Ping)
		m.Get("/", index.Index(indexPath))
	}
}
