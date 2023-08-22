package domain

import (
	"fmt"
	"net/url"
)

type PostLink struct {
	Title string
	Link  url.URL
}

func NewPostLink(title string, link string, host string) (PostLink, error) {
	u, err := url.ParseRequestURI(link)
	if err != nil {
		return PostLink{}, err
	}

	if u.Host == "" {
		u.Scheme = "https"
		u.Host = host
	}
	return PostLink{
		Title: title,
		Link:  *u,
	}, nil
}

func (pl *PostLink) String() string {
	return fmt.Sprintf("[Title: %s, Link: %s]\n", pl.Title, pl.Link.String())
}
