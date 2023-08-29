package repository

import "github.com/ferminhg/testing-colly/internal/domain"

type PostRepository interface {
	Save(post domain.Post) error
	FindAll() ([]domain.Post, error)
}
