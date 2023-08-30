package application

import (
	"github.com/ferminhg/testing-colly/internal/domain"
	"github.com/ferminhg/testing-colly/internal/domain/repository"
)

type Embedder struct {
	postRepository   repository.PostRepository
	slicer           domain.Slicer
	openIAEmbedder   domain.OpenIAEmbedder
	vectorRepository repository.VectorRepository
}

func NewEmbedder(
	postRepository repository.PostRepository,
	openIAEmbedder domain.OpenIAEmbedder,
	vectorRepository repository.VectorRepository,
) *Embedder {
	return &Embedder{
		postRepository:   postRepository,
		slicer:           *domain.NewSlicer(),
		openIAEmbedder:   openIAEmbedder,
		vectorRepository: vectorRepository,
	}
}

func (e *Embedder) Ingest() error {
	posts, err := e.postRepository.Search()
	if err != nil {
		return err
	}

	var slices []domain.PostSlice
	for _, post := range posts {
		slices = append(slices, e.slicer.Slice(post)...)
	}

	vectors, err := e.openIAEmbedder.Embed(slices)
	if err != nil {
		return err
	}

	err = e.vectorRepository.Save(vectors)
	if err != nil {
		return err
	}
	return nil
}
