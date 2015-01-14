package imgur

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type ImgurImage struct {
	Title string
	Url   string
}

func GetExampleImages() ([]*ImgurImage, error) {
	return nil, nil
}

func GetUserImages(albumId string) ([]*ImgurImage, error) {
	return nil, nil
}

func SaveUserImage(albumId string, title string, progData string) error {
	return nil
}

func GetUserAlbumId(username, password string) (string, error) {
	return "", nil
}

func CreateUser(username, password, email, anonAlbum string) (string, error) {
	return "", nil
}

func CreateAnonymousAlbum() (string, error) {
	client := http.Client{}
	req, err := http.NewRequest("POST", "https://api.imgur.com/3/album", strings.NewReader(`{"title":"anon"}`))
	if err != nil {
		return "", err
	}
	addAuth(req)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	fmt.Println(resp, err)
	if resp.StatusCode != 200 {
		return "", errors.New("Bad response code from imgur: " + resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	data := CreateAlbum{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}
	return data.Data.Id, nil
}

func addAuth(req *http.Request) {
	req.Header.Add("Authorization", "Bearer "+credentials.AccessToken)
}

type CreateAlbum struct {
	Success bool `json:"success"`
	Data    struct {
		Id string `json:"id"`
	} `json:"data"`
}

type ImgurCredentials struct {
	AccessToken     string      `json:"access_token"`
	ExpiresIn       int         `json:"expires_in"`
	TokenType       string      `json:"token_type"`
	Scope           interface{} `json:"scope"`
	RefreshToken    string      `json:"refresh_token"`
	AccountUsername string      `json:"account_username"`
	expiration      time.Time   `json:"-"`
}

const tokenUrl string = "https://api.imgur.com/oauth2/token"

var client_id = ""
var client_secret = ""
var credentials *ImgurCredentials

func init() {
	client_id = os.Getenv("imgur_client_id")
	client_secret = os.Getenv("imgur_client_secret")
	start_token := os.Getenv("imgur_refresh_token")
	fmt.Println(client_id, client_secret, start_token)

	var err error
	credentials, err = refreshToken(start_token)
	if err != nil {
		fmt.Println("Problem with initial imgur setup. Will not try.")
	}
}

func refreshToken(refreshToken string) (*ImgurCredentials, error) {
	resp, err := http.PostForm(tokenUrl,
		url.Values{"client_id": {client_id}, "client_secret": {client_secret},
			"refresh_token": {refreshToken}, "grant_type": {"refresh_token"}})
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
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))
	credentials := &ImgurCredentials{}
	err = json.Unmarshal(body, credentials)
	if err != nil {
		credentials.expiration = time.Now().Add(time.Duration(credentials.ExpiresIn) * time.Second)
	}
	return credentials, err
}
