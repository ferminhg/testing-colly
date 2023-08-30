package repository

import "github.com/ferminhg/testing-colly/internal/domain"

type VectorRepository interface {
	Search() ([]domain.Vector, error)
	Save(vectors []domain.Vector) error
}
