package main

import (
	"math/rand"
	"time"
)

type UserSession struct {
	Username     string
	accessToken  string
	refreshToken string
	expires      time.Time
}

//Session id to username
var UserNames map[string]string = map[string]string{}

//username to tokens
var AllSessions map[string]*UserSession = map[string]*UserSession{}

func generateSessionId() string {
	alphabet := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var bytes = make([]byte, 20, 20)
	for i := 0; i < 20; i++ {
		bytes[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(bytes)
}

func NewSession(img *ImgurCredentials) string {
	sid := generateSessionId()
	session := UserSession{}
	uname := img.AccountUsername
	session.Username = uname
	UserNames[sid] = uname
	session.accessToken = img.AccessToken
	session.refreshToken = img.RefreshToken
	session.expires = time.Now().Add(time.Duration(img.ExpiresIn) * time.Second)
	AllSessions[uname] = &session
	return sid
}
