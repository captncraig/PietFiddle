package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/captncraig/pietfiddle/images"
	"github.com/go-martini/martini"
	"github.com/hashicorp/golang-lru"
	"github.com/martini-contrib/render"
	"net/http"
	"strconv"
)

var database Database
var cache *lru.Cache

func init() {
	database = NewBoltDb()
	cache, _ = lru.New(2000)
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{Extensions: []string{".tmpl", ".html"}, Delims: render.Delims{"{[{", "}]}"}}))
	m.Get("/", serveIndex)
	m.Get("/examples", serveExamples)
	m.Post("/save", saveImg)
	m.Get("/:id", serveImg)
	m.Get("/img/(?P<id>~?[a-zA-Z0-9]+).png", renderImage)
	m.Run()
}

func serveIndex(w http.ResponseWriter, r *http.Request, ren render.Render) {
	dat := Image{Width: 10, Height: 10, Data: ""}
	ren.HTML(200, "editor", dat)
}
func saveImg(w http.ResponseWriter, r *http.Request) {
	img := Image{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&img)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	id, err := database.SaveImage(&img)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(id))
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
	cacheKey := fmt.Sprintf("%s-%d", id, cs)
	fmt.Println(cacheKey)
	data, ok := cache.Get(cacheKey)
	fmt.Println(data, ok)
	var b []byte
	if ok {
		b = data.([]byte)
	} else {
		buf := bytes.NewBuffer(nil)
		images.BuildImage(img.Width, img.Height, img.Data, cs, buf)
		b = buf.Bytes()
		cache.Add(cacheKey, b)
	}
	w.Write(b)
}
