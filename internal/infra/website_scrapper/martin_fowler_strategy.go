package website_scrapper

import (
	"github.com/ferminhg/testing-colly/internal/domain"
	"github.com/gocolly/colly"
	"log"
)

type MartinFowlerStrategy struct {
	url                string
	postsQuerySelector string
	domain             string
}

func NewMartinFowlerStrategy() *MartinFowlerStrategy {
	return &MartinFowlerStrategy{
		domain:             "martinfowler.com",
		url:                "https://martinfowler.com/tags/index.html",
		postsQuerySelector: "div.title-list p a",
	}
}

func (m *MartinFowlerStrategy) Execute() error {
	pl, err := m.extractPostLinks()
	log.Println("ℹ️ Num posts found:", len(pl))
	return err
}

func (m *MartinFowlerStrategy) extractPostLinks() ([]domain.PostLink, error) {
	postLinks := make([]domain.PostLink, 0)
	c := colly.NewCollector()

	c.OnHTML(m.postsQuerySelector, func(e *colly.HTMLElement) {
		postLink, err := domain.NewPostLink(e.Text, e.Attr("href"), m.domain)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("✅ Post found", postLink.String())
		//TODO: scrappe post content and extract text
		postLinks = append(postLinks, postLink)
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("🌍 Visiting", r.URL)
	})

	if err := c.Visit(m.url); err != nil {
		return postLinks, err
	}
	return postLinks, nil
}
