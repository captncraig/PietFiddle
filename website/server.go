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
	"net/url"
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
	img, err := database.GetImage(id)
	if err != nil {
		w.WriteHeader(404)
		fmt.Println(err)
		return
	}

	cs := getQuery(r.URL, "cs", 1, 80, 10)
	rotate := getQuery(r.URL, "rot", 0, 17, 0)
	w.Header().Add("Content-Type", "image/png")
	cacheKey := fmt.Sprintf("%s-%d-%d", id, cs, rotate)
	data, ok := cache.Get(cacheKey)
	var b []byte
	if ok {
		b = data.([]byte)
	} else {
		buf := bytes.NewBuffer(nil)
		images.BuildImage(img.Width, img.Height, img.Data, cs, rotate, buf)
		b = buf.Bytes()
		cache.Add(cacheKey, b)
	}
	w.Write(b)
}

func getQuery(u *url.URL, name string, min, max, def int) int {
	if q := u.Query().Get(name); q != "" {
		v64, err := strconv.ParseInt(q, 10, 32)
		if err == nil {
			v := int(v64)
			if v > max {
				v = max
			}
			if v < min {
				v = min
			}
			return v
		}
	}
	return def
}
