package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

type ImgurSession struct {
	Username     string
	accessToken  string
	refreshToken string
	expires      time.Time
}

type ImgurBrowserData struct {
	RegisterUrl string
	Username    string
}

var client_id string = os.Getenv("imgur_client_id")
var client_secret string = os.Getenv("imgur_client_secret")
var registerUrl string = fmt.Sprintf("https://api.imgur.com/oauth2/authorize?client_id=%s&response_type=pin", client_id)
var tokenUrl string = "https://api.imgur.com/oauth2/token"

func NewBrowserData() *ImgurBrowserData {
	return &ImgurBrowserData{registerUrl, ""}
}

func ExchangePinForToken(pin string) error {
	resp, err := http.PostForm(tokenUrl,
		url.Values{"client_id": {client_id}, "client_secret": {client_secret},
			"pin": {pin}, "grant_type": {"pin"}})
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode)
		return nil
	}
	return nil
}

var AllSessions map[string]*ImgurSession = map[string]*ImgurSession{}
