package main

import (
	"errors"
	"fmt"
)

func main() {

}

type AuthorId struct {
}

type Author struct {
	Name string
}

func (a AuthorId) Vaild() bool {
	return false
}

func GetAuther(id AuthorId) (*Author, error) {
	if !id.Vaild() {
		return nil, errors.New("GetAuthor: id is invaild")
	}
	return &Author{}, nil
}

type Book struct {
	AuthorId AuthorId
}

func GetAuthorName(b *Book) (string, error) {
	a, err := GetAuther(b.AuthorId)
	if err != nil {
		return "", fmt.Errorf("GetAuthor: %v", err)
	}
	return a.Name, nil
}
