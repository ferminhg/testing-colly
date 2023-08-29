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
	pl, err := m.extractPostLinks(true)
	log.Println("ℹ️ Num posts found:", len(pl))
	return err
}

func (m *MartinFowlerStrategy) extractPostLinks(dryRun bool) ([]domain.PostLink, error) {
	postLinks := make([]domain.PostLink, 0)
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
		postLink, err := domain.NewPostLink(e.Text, e.Attr("href"), m.domain)
		if err != nil {
			log.Println(err)
			return
		}

		log.Println("✅ Post found", postLink.String())
		if err := postCollector.Visit(postLink.Link.String()); err != nil {
			log.Println("🚨 Error scraping post", err)
			return
		}
		postLinks = append(postLinks, postLink)
		linksFounds++

		if dryRun && linksFounds > 10 {
			log.Println("🚨 [DryRun] Max links scrapped reached")
		}
	})

	postCollector.OnHTML("div.paperBody", func(e *colly.HTMLElement) {
		log.Println("📄 Post content", e.Text)
		postsScrapped++
	})

	log.Println("ℹ️ Num links found:", linksFounds)
	log.Println("ℹ️ Num posts scrapped:", postsScrapped)

	if err := tagCollector.Visit(m.url); err != nil {
		log.Println("🚨 Error scraping tag page", err)
		return postLinks, err
	}

	return postLinks, nil
}
