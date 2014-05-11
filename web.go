package gojistatic

import (
	"flag"
	"log"
	"net/http"

	"github.com/zenazn/goji/bind"
	"github.com/zenazn/goji/graceful"
	"github.com/zenazn/goji/web"

	"github.com/ed0wolf/gojistatic/controllers"
	"github.com/ed0wolf/gojistatic/middleware/assets"
)

var app *web.Mux
var assetRoot string

func init() {
	app = web.New()

	defaultAssetRoot := "./assets"
	flag.StringVar(&assetRoot, "assets", defaultAssetRoot,
		"Path to the top level assets folder")

	//Add middleware
	retreiver := &assets.FileAssetsRetriever{assetRoot}
	handler := &assets.AssetsHandler{retreiver}
	app.Use(handler.HandleAssets)

	//Add routes
	app.Get("/ping", controllers.Ping)
	app.Get("/", controllers.Index(assetRoot+"/index.html"))
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
