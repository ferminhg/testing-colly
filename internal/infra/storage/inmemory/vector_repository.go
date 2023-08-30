package inmemory

import (
	"github.com/ferminhg/testing-colly/internal/domain"
)

type VectorRepository struct {
	vectors []domain.Vector
}

func NewVectorRepository() *VectorRepository {
	return &VectorRepository{
		vectors: []domain.Vector{},
	}
}

func (v *VectorRepository) Search() ([]domain.Vector, error) {
	return v.vectors, nil
}

func (v *VectorRepository) Save(vectors []domain.Vector) error {
	v.vectors = append(v.vectors, vectors...)
	return nil
}
