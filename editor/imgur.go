package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type ImgurBrowserData struct {
	RegisterUrl string
	Username    string
}

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
var accountUrl string = "https://api.imgur.com/3/account/me"

func NewBrowserData() *ImgurBrowserData {
	return &ImgurBrowserData{registerUrl, ""}
}

func ExchangePinForToken(pin string) (*ImgurCredentials, error) {
	resp, err := http.PostForm(tokenUrl,
		url.Values{"client_id": {client_id}, "client_secret": {client_secret},
			"pin": {pin}, "grant_type": {"pin"}})
	if err != nil {
		return nil, err
	}
	return parseTokenResponse(resp)
}

func RefreshToken(tok *ImgurCredentials) (*ImgurCredentials, error) {
	resp, err := http.PostForm(tokenUrl,
		url.Values{"client_id": {client_id}, "client_secret": {client_secret},
			"refresh_token": {tok.RefreshToken}, "grant_type": {"refresh_token"}})
	if err != nil {
		return nil, err
	}
	return parseTokenResponse(resp)
}

func parseTokenResponse(resp *http.Response) (*ImgurCredentials, error) {
	if resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode)
		return nil, errors.New("Bad pin response from imgur")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	credentials := &ImgurCredentials{}
	err := json.Unmarshal(body, credentials)
	return credentials, err
}
