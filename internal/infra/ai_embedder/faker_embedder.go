package ai_embedder

import (
	"log"

	"github.com/ferminhg/testing-colly/internal/domain"
)

type FakerEmbedder struct {
}

func NewFakerEmbedder() *FakerEmbedder {
	return &FakerEmbedder{}
}

func (e FakerEmbedder) Embed(slices []domain.PostSlice) ([]domain.Vector, error) {
	log.Println("FakerEmbedder: Embedding...", slices)
	vectors := make([]domain.Vector, len(slices))
	return vectors, nil
}
