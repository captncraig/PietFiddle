package main

import (
	_ "github.com/boltdb/bolt"
)

type Image struct {
	Id   string
	Name string
	Data string
}
type Database interface {
	GetExampleImages() ([]*Image, error)
}

type boltDb struct{}

func NewBoltDb() Database {
	return &boltDb{}
}

func (b *boltDb) GetExampleImages() ([]*Image, error) {
	return nil, nil
}
