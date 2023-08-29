package inmemory

import "github.com/ferminhg/testing-colly/internal/domain"

type PostRepository struct {
	posts []domain.Post
}

func NewPostRepository() *PostRepository {
	return &PostRepository{}
}

func (p *PostRepository) Save(post domain.Post) error {
	p.posts = append(p.posts, post)
	return nil
}

func (p *PostRepository) FindAll() ([]domain.Post, error) {
	return p.posts, nil
}
