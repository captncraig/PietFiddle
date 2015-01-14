package main

import (
	"fmt"
	"github.com/captncraig/pietfiddle/imgur"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{Extensions: []string{".tmpl", ".html"}, Delims: render.Delims{"{[{", "}]}"}}))
	m.Get("/", serveIndex)
	m.Run()
}

func serveIndex(w http.ResponseWriter, r *http.Request, ren render.Render) {
	c, err := r.Cookie("Uimg")
	fmt.Println(c, err)
	if err != nil {
		//user cookie not found. Not logged in.
		//check anonymous id
		anonCookie, err := r.Cookie("Anon")
		fmt.Println(anonCookie, err)
		if err != nil {
			//nothing. Don't make anonymous album until they want to save something
		}
	}
	ren.HTML(200, "editor", nil)
}
