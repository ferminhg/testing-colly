package domain

import (
	"fmt"
	"github.com/google/uuid"
	"net/url"
)

type Post struct {
	id     uuid.UUID
	title  string
	link   url.URL
	text   string
	author string
}

func (p *Post) SetText(text string) {
	// TODO normalizar texto

	p.text = text
}

func (p *Post) Title() string {
	return p.title
}

func (p *Post) Link() string {
	return p.link.String()
}

func (p *Post) Text() string {
	return p.text
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
		id:     id,
		title:  title,
		link:   *u,
		author: author,
	}, nil
}

func (p *Post) String() string {
	return fmt.Sprintf("#id: %s \t title: %s, [%s]", p.id, p.Title(), p.Link())
}
