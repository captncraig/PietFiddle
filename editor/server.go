package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{Extensions: []string{".tmpl", ".html"}, Delims: render.Delims{"{[{", "}]}"}}))
	m.Get("/", serveIndex)
	m.Get("/pin/:pin", handlePinAuth)
	m.Run()
}

func serveIndex(w http.ResponseWriter, r *http.Request, ren render.Render) {

	ren.HTML(200, "editor", NewBrowserData())
}

func handlePinAuth(w http.ResponseWriter, params martini.Params) {
	pin := params["pin"]
	ExchangePinForToken(pin)
}
