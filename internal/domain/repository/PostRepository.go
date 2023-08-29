package repository

import "github.com/ferminhg/testing-colly/internal/domain"

type PostRepository interface {
	Save(post domain.Post) error
	Search() ([]domain.Post, error)
	Update(post domain.Post) error
	FindByLink(link string) (domain.Post, bool)
}
