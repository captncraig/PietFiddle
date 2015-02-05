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

type userId string

func main() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{Extensions: []string{".tmpl", ".html"}, Delims: render.Delims{"{[{", "}]}"}}))
	m.Get("/", serveIndex)
	m.Get("/examples", serveExamples)
	m.Post("/save", saveImg)
	m.Get("/import", importImg)
	m.Get("/:id", serveImg)

	m.Get("/img/(?P<id>~?[a-zA-Z0-9]+).png", renderPng)
	m.Get("/img/(?P<id>~?[a-zA-Z0-9]+).gif", renderGif)

	//decode user token and make it available to any handler that asks for a userId.
	m.Use(func(r *http.Request, c martini.Context) {
		a := userId("")
		if cookie, err := r.Cookie("pf-auth"); err == nil {
			if v := cookie.Value; v != "" {
				id := database.GetUserId(v)
				if id != "" {
					a = userId(id)
				}
			}
		}
		c.Map(a)
		c.Next()
	})

	m.Run()
}

type BrowserData struct {
	Img *Image
	Uid userId
}

func serveIndex(w http.ResponseWriter, r *http.Request, ren render.Render, uid userId) {
	fmt.Println("u", uid)
	img := Image{Width: 10, Height: 10, Data: ""}
	ren.HTML(200, "editor", BrowserData{&img, uid})
}
func serveImg(w http.ResponseWriter, params martini.Params, ren render.Render, uid userId) {
	id := params["id"]
	img, err := database.GetImage(id)
	if err != nil {
		w.WriteHeader(404)
		fmt.Println(err)
		return
	}
	ren.HTML(200, "editor", BrowserData{img, uid})
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
func importImg(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	resp, err := http.Get(url)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	wid, h, d, err := images.LoadImage(resp.Body, 1)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Println(wid, h, d)
	i := Image{}
	i.Data = d
	i.Width = wid
	i.Height = h
	id, err := database.SaveImage(&i)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	http.Redirect(w, r, "/"+id, 302)
}
func serveExamples(ren render.Render) {
	ren.HTML(200, "examples", database.GetExampleImages())
}
func renderGif(w http.ResponseWriter, r *http.Request, params martini.Params) {
	renderImage(w, r, params, "G")
}
func renderPng(w http.ResponseWriter, r *http.Request, params martini.Params) {
	renderImage(w, r, params, "P")
}
func renderImage(w http.ResponseWriter, r *http.Request, params martini.Params, format string) {
	id := params["id"]
	img, err := database.GetImage(id)
	if err != nil {
		w.WriteHeader(404)
		fmt.Println(err)
		return
	}

	cs := getQuery(r.URL, "cs", 1, 80, 10)
	rotate := getQuery(r.URL, "rot", 0, 17, 0)

	cacheKey := fmt.Sprintf("%s-%d-%d-%s", id, cs, rotate, format)
	data, ok := cache.Get(cacheKey)
	var b []byte
	if ok {
		b = data.([]byte)
	} else {
		buf := bytes.NewBuffer(nil)
		if format == "G" {
			w.Header().Add("Content-Type", "image/gif")
			images.BuildGif(img.Width, img.Height, img.Data, cs, buf)
		} else {
			w.Header().Add("Content-Type", "image/png")
			images.BuildImage(img.Width, img.Height, img.Data, cs, rotate, buf)
		}
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
