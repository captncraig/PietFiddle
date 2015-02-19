package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

var ppm = flag.Bool("ppm", false, "Download ppms instead of pngs.")

func main() {
	flag.Parse()
	fmt := "png"
	if *ppm {
		fmt := "p6"
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers",
				"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		}
		if id := r.URL.Query().Get("img"); id != "" {
			fmt.Printf("Save detected for %s. Downloading...\n", id)
			url := "http://www.pietfiddle.net/img/" + id + fmt + "?cs=1"
			resp, err := http.Get(url)
			if err != nil {
				fmt.Printf("Error downloading! %v", err.Error())
				return
			}
			if resp.StatusCode != 200 {
				fmt.Printf("Bad status code: %d", resp.StatusCode)
				return
			}
			file, err := os.Create("work.png")
			defer file.Close()
			if err != nil {
				fmt.Printf("Error opening file! %v", err.Error())
				return
			}
			io.Copy(file, resp.Body)
			fmt.Println("Success!")
		}
	})
	http.ListenAndServe(":4987", nil)
}
