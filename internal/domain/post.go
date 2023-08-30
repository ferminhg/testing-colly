package domain

import (
	"fmt"
	"net/url"

	"github.com/google/uuid"
)

type Post struct {
	Id     uuid.UUID `json:"id"`
	Title  string    `json:"title"`
	Link   url.URL   `json:"link" custom:"url"`
	Text   string    `json:"text"`
	Author string    `json:"author"`
}

func (p *Post) SetText(text string) {
	// TODO normalizar texto

	p.Text = text
}

func NewPost(title string, link string, host string, author string) (Post, error) {
	u, err := url.ParseRequestURI(link)
	if err != nil {
		return Post{}, err
	}
	if u.Host == "" {
		u.Scheme = "https"
		u.Host = host
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return Post{}, err
	}

	return Post{
		Id:     id,
		Title:  title,
		Link:   *u,
		Author: author,
	}, nil
}

func (p *Post) String() string {
	return fmt.Sprintf("#id: %s \t title: %s, [%s]", p.Id, p.Title, p.Link.String())
}
