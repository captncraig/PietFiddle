package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"time"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{Extensions: []string{".tmpl", ".html"}, Delims: render.Delims{"{[{", "}]}"}}))
	m.Get("/", serveIndex)
	m.Get("/pin/:pin", handlePinAuth)
	m.Run()
}

func serveIndex(w http.ResponseWriter, r *http.Request, ren render.Render) {
	data := NewBrowserData()
	cookie, _ := r.Cookie("Session_Id")
	if cookie != nil {
		fmt.Printf("Cookie value:%s\n", cookie.Value)
		uname := UserNames[cookie.Value]
		fmt.Printf("Username:%s\n", uname)
		if uname != "" {
			data.Username = uname
		}
	}
	ren.HTML(200, "editor", data)
}

func handlePinAuth(w http.ResponseWriter, params martini.Params) {
	pin := params["pin"]
	tok, err := ExchangePinForToken(pin)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	tok, err = RefreshToken(tok)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	sid := NewSession(tok)
	cookie := &http.Cookie{}
	cookie.Name = "Session_Id"
	cookie.Expires = time.Now().AddDate(5, 0, 0)
	cookie.Value = sid
	cookie.Path = "/"
	http.SetCookie(w, cookie)

}
