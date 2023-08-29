package website_scrapper

import (
	"github.com/ferminhg/testing-colly/internal/domain"
	"github.com/ferminhg/testing-colly/internal/domain/repository"
	"github.com/gocolly/colly"
	"log"
)

type MartinFowlerStrategy struct {
	url                string
	postsQuerySelector string
	domain             string
	postRepository     repository.PostRepository
}

func NewMartinFowlerStrategy(postRepository repository.PostRepository) *MartinFowlerStrategy {
	return &MartinFowlerStrategy{
		domain:             "martinfowler.com",
		url:                "https://martinfowler.com/tags/index.html",
		postsQuerySelector: "div.title-list p a",
		postRepository:     postRepository,
	}
}

func (m *MartinFowlerStrategy) Execute() error {
	pl, err := m.extractPostLinks(true)
	log.Println("ℹ️ Num posts found:", len(pl))
	return err
}

func (m *MartinFowlerStrategy) extractPostLinks(dryRun bool) ([]domain.Post, error) {
	posts := make([]domain.Post, 0)
	tagCollector := colly.NewCollector()
	postCollector := colly.NewCollector()

	linksFounds, postsScrapped := 0, 0

	tagCollector.OnRequest(func(r *colly.Request) {
		log.Println("🌍 Visiting Tag Page", r.URL)

	})

	postCollector.OnRequest(func(r *colly.Request) {
		log.Println("\t 🌍 Visiting Post Page", r.URL)
	})

	tagCollector.OnHTML(m.postsQuerySelector, func(e *colly.HTMLElement) {
		if dryRun && linksFounds > 10 {
			return
		}
		postLink, err := domain.NewPost(e.Text, e.Attr("href"), m.domain, "Martin Fowler")
		if err != nil {
			log.Println(err)
			return
		}

		log.Println("✅ Post found", postLink.String())
		if err := postCollector.Visit(postLink.Link()); err != nil {
			log.Println("🚨 Error scraping post", err)
			return
		}
		posts = append(posts, postLink)
		linksFounds++

		if dryRun && linksFounds > 10 {
			log.Println("🚨 [DryRun] Max links scrapped reached")
		}
	})

	postCollector.OnHTML("div.paperBody", func(e *colly.HTMLElement) {
		log.Println("📄 Post content:", e.Text[:100])
		log.Println("End of Post")
		postsScrapped++
	})

	log.Println("ℹ️ Num links found:", linksFounds)
	log.Println("ℹ️ Num posts scrapped:", postsScrapped)

	if err := tagCollector.Visit(m.url); err != nil {
		log.Println("🚨 Error scraping tag page", err)
		return posts, err
	}

	return posts, nil
}
