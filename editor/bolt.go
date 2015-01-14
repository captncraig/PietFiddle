package main

import (
	_ "github.com/boltdb/bolt"
)

type boltDb struct{}

func NewBoltDb() Database {
	return &boltDb{}
}

func (b *boltDb) GetExampleImages() ([]*Image, error) {
	return nil, nil
}
func (b *boltDb) GetUserImages(userId string) ([]*Image, error) {
	return nil, nil
}
func (b *boltDb) SaveUserImage(userId string, img *Image) error {
	return nil
}
func (b *boltDb) CreateUser(username, password, email, anonId string) (string, error) {
	return "", nil
}
func (b *boltDb) CreateAnonUser(anonId string) (string, error) {
	return "", nil
}

func (b *boltDb) LookupUser(username, password string) (string, error) {
	return "", nil
}

func (b *boltDb) LookupUserFromSessionId(sessionId string) (userId, username string, err error) {
	return "", "", nil
}
