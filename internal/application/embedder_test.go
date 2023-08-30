package application

import (
	"testing"

	"github.com/ferminhg/testing-colly/internal/infra/ai_embedder"
	"github.com/ferminhg/testing-colly/internal/infra/objet_mother"
	"github.com/ferminhg/testing-colly/internal/infra/storage/inmemory"
	"github.com/stretchr/testify/assert"
)

func Test_EmbedderHappyPath(t *testing.T) {
	postObjetMother := objet_mother.PostObjetMother{}

	t.Run("should embed all posts", func(t *testing.T) {
		pr := inmemory.NewPostRepository()
		pr.Save(postObjetMother.Random())

		vr := inmemory.NewVectorRepository()
		emb := ai_embedder.NewFakerEmbedder()

		//Given a post from a post repository
		embedder := NewEmbedder(pr, emb, vr)

		//When the embedder is executed
		assert.NoError(t, embedder.Ingest())

		//Then the vector repository should have the vector
		vectors, _ := vr.Search()
		assert.Equal(t, 2, len(vectors))
	})
}
