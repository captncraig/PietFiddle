package main

import (
	"fmt"
	"github.com/captncraig/pietfiddle/images"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"strconv"
)

var database Database

func init() {
	database = NewBoltDb()
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{Extensions: []string{".tmpl", ".html"}, Delims: render.Delims{"{[{", "}]}"}}))
	m.Get("/", serveIndex)
	m.Get("/examples", serveExamples)
	m.Get("/:id", serveImg)
	m.Get("/img/(?P<id>~?[a-zA-Z0-9]+).png", renderImage)
	m.Run()
}

func serveIndex(w http.ResponseWriter, r *http.Request, ren render.Render) {
	dat := Image{Width: 10, Height: 10, Data: ""}
	ren.HTML(200, "editor", dat)
}

func serveImg(w http.ResponseWriter, params martini.Params, ren render.Render) {
	id := params["id"]
	img, err := database.GetImage(id)
	if err != nil {
		w.WriteHeader(404)
		fmt.Println(err)
		return
	}
	ren.HTML(200, "editor", img)
}

func serveExamples(ren render.Render) {
	ren.HTML(200, "examples", database.GetExampleImages())
}

func renderImage(w http.ResponseWriter, r *http.Request, params martini.Params) {
	id := params["id"]
	cs := 10
	csQuery := r.URL.Query().Get("cs")
	if csQuery != "" {
		cs64, err := strconv.ParseInt(csQuery, 10, 32)
		if err == nil {
			cs = int(cs64)
			if cs > 80 {
				cs = 80
			}
			if cs < 1 {
				cs = 1
			}
		}
	}
	fmt.Println(id)
	img, err := database.GetImage(id)
	if err != nil {
		w.WriteHeader(404)
		fmt.Println(err)
		return
	}
	w.Header().Add("Content-Type", "image/png")
	images.BuildImage(img.Width, img.Height, img.Data, cs, w)
}
