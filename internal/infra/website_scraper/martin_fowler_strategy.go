package website_scraper

import (
	"bytes"
	"encoding/json"
	"log"
	"os"

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

	posts, _ := m.repository.Search()
	log.Println("ℹ️ Num posts saved:", len(posts))
	return nil
}

func (m *MartinFowlerStrategy) newTagCollector(pc *colly.Collector) *colly.Collector {
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		log.Println("🌍 Visiting Tag Page", r.URL)
	})

	c.OnHTML(m.postsQuerySelector, func(e *colly.HTMLElement) {
		if m.dryRun && m.linkFounds > MaxPosts {
			return
		}
		postLink, err := domain.NewPost(e.Text, e.Attr("href"), m.domain, m.author)
		if err != nil {
			log.Println(err)
			return
		}

		if err := m.repository.Save(postLink); err != nil {
			log.Println("🚨 Error saving post", err)
		}
		m.linkFounds++
		if err := pc.Visit(postLink.Link.String()); err != nil {
			log.Println("🚨 Error scraping post", err)
			return
		}
		if m.dryRun && m.linkFounds > MaxPosts {
			log.Println("🚨 [DryRun] Max links scrapped reached")
		}
	})

	return c
}

func (m *MartinFowlerStrategy) newPostCollector() *colly.Collector {
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		log.Println("🌍 Visiting Post Page", r.URL)
	})
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

func (m *MartinFowlerStrategy) Marshal(file string) error {
	buf := new(bytes.Buffer)
	posts, err := m.repository.Search()
	if err != nil {
		return err
	}

	if err := json.NewEncoder(buf).Encode(posts); err != nil {
		panic(err)
	}

	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.Write(buf.Bytes()); err != nil {
		return err
	}

	return nil
}
