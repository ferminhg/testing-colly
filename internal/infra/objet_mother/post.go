package objet_mother

import (
	"github.com/ferminhg/testing-colly/internal/domain"
	"github.com/go-faker/faker/v4"
)

type PostObjetMother struct {
}

func (p *PostObjetMother) Random() domain.Post {
	post, _ := domain.NewPost(
		faker.Sentence(),
		faker.URL(),
		faker.DomainName(),
		"faker")
	post.SetText(faker.Paragraph())
	return post
}
