package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("!!!", r.URL.Query().Get("img"))
	})
	http.ListenAndServe(":4987", nil)
}
