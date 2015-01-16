package main

import (
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/boltdb/bolt"
	"github.com/captncraig/pietfiddle/images"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Image struct {
	Id            string
	Name          string
	Width, Height int
	Data          string
}

var examples []*Image

type ImgurAlbum struct {
	Data struct {
		Images []struct {
			Title string `json:"title"`
			Id    string `json:"id"`
			Link  string `json:"link"`
		} `json:"images"`
	} `json:"data"`
}

func init() {
	examples = []*Image{}
	client := http.Client{}
	req, _ := http.NewRequest("GET", "https://api.imgur.com/3/album/z4otu", nil)
	if clientId := os.Getenv("imgur_client_id"); clientId != "" {
		req.Header.Add("Authorization", "Client-ID "+clientId)
	}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		log.Println("Error getting example images.", err, resp)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error getting example images.", err)
	}
	album := ImgurAlbum{}
	err = json.Unmarshal(body, &album)
	if err != nil {
		log.Println("Error getting example images.", err)
	}
	for _, img := range album.Data.Images {
		fmt.Println(img.Id, img.Title, img.Link)
		resp, err := http.Get(img.Link)
		if err != nil || resp.StatusCode != 200 {
			log.Println("Error getting example images.", err, resp)
		}
		w, h, d := images.LoadImage(resp.Body, 1)
		image := Image{
			Id:     "~" + img.Id,
			Width:  w,
			Height: h,
			Data:   d,
		}
		examples = append(examples, &image)
	}
}

type Database interface {
	GetExampleImages() []*Image
	GetImage(id string) (*Image, error)
}

type boltDb struct{}

func NewBoltDb() Database {
	return &boltDb{}
}

func (b *boltDb) GetExampleImages() []*Image {
	return examples
}

func (b *boltDb) GetImage(id string) (*Image, error) {
	if id[0] == '~' {
		for _, ex := range examples {
			if ex.Id == id {
				return ex, nil
			}
		}
	}
	return nil, errors.New("Not found")
}
