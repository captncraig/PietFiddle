package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/captncraig/pietfiddle/images"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type Image struct {
	Id            string
	Name          string
	Width, Height int
	Data          string
}

var examples []*Image
var bdb *bolt.DB

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
	rand.Seed(time.Now().UTC().UnixNano())
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
			Name:   img.Title,
			Height: h,
			Data:   d,
		}
		examples = append(examples, &image)
	}
	dbPath := "pietfiddle.db"
	if p := os.Getenv("BOLT_PATH"); p != "" {
		dbPath = filepath.Join(p, dbPath)
	}
	bdb, err = bolt.Open(dbPath, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	err = bdb.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("images"))
		return err
	})
	if err != nil {
		log.Fatal(err)
	}
}

type Database interface {
	GetExampleImages() []*Image
	GetImage(id string) (*Image, error)
	SaveImage(i *Image) (string, error)
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

type boltDb struct{}

func NewBoltDb() Database {
	return &boltDb{}
}

func (b *boltDb) GetExampleImages() []*Image {
	return examples
}

func (b *boltDb) SaveImage(i *Image) (string, error) {
	id := randSeq(10)
	err := bdb.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("images"))
		j, _ := json.Marshal(i)
		err := b.Put([]byte(id), j)
		return err
	})
	return id, err
}

func (b *boltDb) GetImage(id string) (*Image, error) {
	if id[0] == '~' {
		for _, ex := range examples {
			if ex.Id == id {
				return ex, nil
			}
		}
	}
	i := Image{}
	err := bdb.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("images"))
		v := b.Get([]byte(id))
		if v == nil {
			return errors.New("Not found")
		}
		return json.Unmarshal(v, &i)
	})
	return &i, err

}
