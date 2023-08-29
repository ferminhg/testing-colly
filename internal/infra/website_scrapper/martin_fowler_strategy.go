package website_scrapper

import (
	"github.com/ferminhg/testing-colly/internal/domain"
	"github.com/ferminhg/testing-colly/internal/domain/repository"
	"github.com/gocolly/colly"
	"log"
)

type MartinFowlerStrategy struct {
	url                   string
	postsQuerySelector    string
	postTextQuerySelector string
	domain                string
	repository            repository.PostRepository
}

func NewMartinFowlerStrategy(repository repository.PostRepository) *MartinFowlerStrategy {
	return &MartinFowlerStrategy{
		domain:                "martinfowler.com",
		url:                   "https://martinfowler.com/tags/index.html",
		postsQuerySelector:    "div.title-list p a",
		postTextQuerySelector: "div.paperBody",
		repository:            repository,
	}
}

func (m *MartinFowlerStrategy) Execute() error {
	if err := m.extractPostLinks(true); err != nil {
		return err
	}
	posts, _ := m.repository.Search()
	log.Println("ℹ️ Num posts saved:", len(posts))
	return nil
}

func (m *MartinFowlerStrategy) extractPostLinks(dryRun bool) error {
	tagCollector, postCollector := colly.NewCollector(), colly.NewCollector()
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

		if err := m.repository.Save(postLink); err != nil {
			log.Println("🚨 Error saving post", err)
		}
		linksFounds++
		if err := postCollector.Visit(postLink.Link()); err != nil {
			log.Println("🚨 Error scraping post", err)
			return
		}
		if dryRun && linksFounds > 10 {
			log.Println("🚨 [DryRun] Max links scrapped reached")
		}
	})

	postCollector.OnHTML(m.postTextQuerySelector, func(e *colly.HTMLElement) {
		log.Println("\t \t ✅ Post content found")
		post, found := m.repository.FindByLink(e.Request.URL.String())
		if !found {
			log.Println("💾 🚨 Post not found: ", e.Request.URL.String())
			return
		}

		post.SetText(e.Text)
		if err := m.repository.Update(post); err != nil {
			log.Println("💾 🚨 Error updating post", err)
		}

		postsScrapped++
	})

	if err := tagCollector.Visit(m.url); err != nil {
		log.Println("🚨 Error scraping tag page", err)
		return err
	}

	log.Println("ℹ️ Num posts scrapped:", postsScrapped)
	return nil
}
