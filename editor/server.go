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
	cookie, _ := r.Cookie("SSID")
	fmt.Println(cookie.Value)
	session := AllSessions[cookie.Value]
	fmt.Println(session)
	fmt.Println(AllSessions)

	if session != nil {
		fmt.Println(session.Username)
		data.Username = session.Username
	}
	ren.HTML(200, "editor", NewBrowserData())
}

func handlePinAuth(w http.ResponseWriter, params martini.Params) {
	pin := params["pin"]
	sessionId, err := ExchangePinForToken(pin)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	} else {
		cookie := &http.Cookie{}
		cookie.Name = "SSID"
		cookie.Secure = false
		cookie.Expires = time.Now().AddDate(5, 0, 0)
		cookie.Value = sessionId
		cookie.Domain = "127.0.0.1"
		cookie.Path = "/"
		http.SetCookie(w, cookie)
	}

}
