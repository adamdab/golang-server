package data

import (
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) Save() error {
	filename := "resources/data/" + p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func LoadPage(title string) (*Page, error) {
	filename := "resources/data/" + title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
