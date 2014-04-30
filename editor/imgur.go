package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
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

var AllSessions map[string]*ImgurSession = map[string]*ImgurSession{}

type ImgurCredentials struct {
	AccessToken     string      `json:"access_token"`
	ExpiresIn       int         `json:"expires_in"`
	TokenType       string      `json:"token_type"`
	Scope           interface{} `json:"scope"`
	RefreshToken    string      `json:"refresh_token"`
	AccountUsername string      `json:"account_username"`
}

var client_id string = os.Getenv("imgur_client_id")
var client_secret string = os.Getenv("imgur_client_secret")
var registerUrl string = fmt.Sprintf("https://api.imgur.com/oauth2/authorize?client_id=%s&response_type=pin", client_id)
var tokenUrl string = "https://api.imgur.com/oauth2/token"

func NewBrowserData() *ImgurBrowserData {
	return &ImgurBrowserData{registerUrl, ""}
}

func ExchangePinForToken(pin string) (string, error) {
	resp, err := http.PostForm(tokenUrl,
		url.Values{"client_id": {client_id}, "client_secret": {client_secret},
			"pin": {pin}, "grant_type": {"pin"}})
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode)
		return "", nil
	}
	defer resp.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	credentials := &ImgurCredentials{}
	json.Unmarshal(body, credentials)
	id := generateSessionId()
	fmt.Println(credentials.AccountUsername)
	AllSessions[id] = &ImgurSession{credentials.AccountUsername,
		credentials.AccessToken, credentials.RefreshToken,
		time.Now().Add(time.Duration(credentials.ExpiresIn) * time.Second)}
	return id, nil
}

func generateSessionId() string {
	alphabet := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var bytes = make([]byte, 20, 20)
	for i := 0; i < 20; i++ {
		bytes[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(bytes)
}
