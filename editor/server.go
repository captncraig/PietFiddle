package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
)

type Image struct {
	Id   string
	Name string
	Data string
}
type Database interface {
	GetExampleImages() ([]*Image, error)
	GetUserImages(userId string) ([]*Image, error)
	SaveUserImage(userId string, img *Image) error
	CreateUser(username, password, email, anonId string) (string, error)
	CreateAnonUser(anonId string) (string, error)
	LookupUser(username, password string) (string, error)
	LookupUserFromSessionId(sessionId string) (userId, username string, err error)
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{Extensions: []string{".tmpl", ".html"}, Delims: render.Delims{"{[{", "}]}"}}))
	m.Get("/", serveIndex)
	m.Run()
}

func serveIndex(w http.ResponseWriter, r *http.Request, ren render.Render) {

	ren.HTML(200, "editor", nil)
}
