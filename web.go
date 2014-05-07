package gojistatic

import (
	"log"
	"net/http"

	"github.com/zenazn/goji/bind"
	"github.com/zenazn/goji/graceful"
	"github.com/zenazn/goji/web"

	"github.com/ed0wolf/gojistatic/controllers"
	"github.com/ed0wolf/gojistatic/middleware/assets"
)

var app *web.Mux

func init() {
	app = web.New()

	//Add middleware
	retreiver := &assets.FileAssetsRetriever{"./assets"}
	handler := &assets.AssetsHandler{retreiver}
	app.Use(handler.HandleAssets)

	//Add routes
	controllers.AddRoutes(app)
}

func Start() {
	http.Handle("/", app)

	listener := bind.Default()

	log.Println("Starting Goji on", listener.Addr())

	bind.Ready()

	err := graceful.Serve(listener, http.DefaultServeMux)

	if err != nil {
		log.Fatal(err)
	}
}
