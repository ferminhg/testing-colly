package inmemory

import (
	"github.com/ferminhg/testing-colly/internal/domain"
)

type PostRepository struct {
	posts map[string]domain.Post
}

func NewPostRepository() *PostRepository {
	return &PostRepository{
		posts: make(map[string]domain.Post),
	}
}

func (p *PostRepository) Save(post domain.Post) error {
	p.posts[post.Link.String()] = post
	return nil
}

func (p *PostRepository) Search() ([]domain.Post, error) {
	var values []domain.Post
	for _, value := range p.posts {
		values = append(values, value)
	}
	return values, nil
}

func (p *PostRepository) FindByLink(link string) (domain.Post, bool) {
	val, ok := p.posts[link]
	return val, ok
}

func (p *PostRepository) Update(post domain.Post) error {
	p.posts[post.Link.String()] = post
	return nil
}
