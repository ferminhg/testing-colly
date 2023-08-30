package website_scraper

import (
	"bytes"
	"encoding/json"
	"github.com/ferminhg/testing-colly/internal/infra/storage/filesystem"
	"log"
	"strings"

	"github.com/ferminhg/testing-colly/internal/domain"
	"github.com/ferminhg/testing-colly/internal/domain/repository"
	"github.com/gocolly/colly"
)

type MartinFowlerStrategy struct {
	url                   string
	author                string
	postsQuerySelector    string
	postTextQuerySelector string
	domain                string
	repository            repository.PostRepository
	dryRun                bool
	linkFounds            int
}

func NewMartinFowlerStrategy(repository repository.PostRepository, dryRun bool) *MartinFowlerStrategy {
	return &MartinFowlerStrategy{
		domain:                "martinfowler.com",
		author:                "Martin Fowler",
		url:                   "https://martinfowler.com/tags/index.html",
		postsQuerySelector:    "div.title-list p a",
		postTextQuerySelector: "div.paperBody",
		repository:            repository,
		dryRun:                dryRun,
		linkFounds:            0,
	}
}

const MaxPosts = 10

func (m *MartinFowlerStrategy) Execute() error {
	tagCollector := m.newTagCollector(m.newPostCollector())

	if err := tagCollector.Visit(m.url); err != nil {
		log.Println("🚨 Error scraping tag page", err)
		return err
	}

	if err := m.marshalPosts(); err != nil {
		log.Println("🚨 Error marshalling posts", err)
		return err
	}
	return nil
}

func (m *MartinFowlerStrategy) newTagCollector(pc *colly.Collector) *colly.Collector {
	c := colly.NewCollector()
	c.OnRequest(logVisitingPage())

	c.OnHTML(m.postsQuerySelector, func(e *colly.HTMLElement) {
		if m.checkDryRun() {
			return
		}

		post, err := m.buildAndSavePostFromHTML(e)
		if err != nil {
			log.Println(err)
			return
		}

		if err := pc.Visit(post.Link.String()); err != nil {
			log.Println("🚨 Error scraping post", err)
			return
		}
	})

	return c
}

func (m *MartinFowlerStrategy) checkDryRun() bool {
	if m.dryRun && m.linkFounds > MaxPosts {
		return true
	}
	return false
}

func (m *MartinFowlerStrategy) buildAndSavePostFromHTML(e *colly.HTMLElement) (*domain.Post, error) {
	post, err := domain.NewPost(e.Text, e.Attr("href"), m.domain, m.author)
	if err != nil {
		return nil, err
	}

	if err := m.repository.Save(post); err != nil {
		log.Println("🚨 Error saving post", err)
	}
	m.linkFounds++
	return &post, nil
}

func (m *MartinFowlerStrategy) newPostCollector() *colly.Collector {
	c := colly.NewCollector()
	c.OnRequest(logVisitingPage())

	c.OnHTML(m.postTextQuerySelector, func(e *colly.HTMLElement) {
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
	})
	return c
}

func (m *MartinFowlerStrategy) marshalPosts() error {
	buf := new(bytes.Buffer)
	posts, err := m.repository.Search()
	if err != nil {
		return err
	}

	if err := json.NewEncoder(buf).Encode(posts); err != nil {
		return err
	}

	outputFile := filesystem.OutputDir +
		strings.ReplaceAll(m.author, " ", "") +
		"-scraped-posts.json"

	if err := filesystem.WriteBuffer2File(outputFile, buf); err != nil {
		return err
	}

	return nil
}

func logVisitingPage() colly.RequestCallback {
	return func(r *colly.Request) {
		log.Println("🌍 Visiting Page: ", r.URL)
	}
}
