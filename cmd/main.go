package main

import (
	"fmt"
	"github.com/ferminhg/testing-colly/internal/application"
	"github.com/ferminhg/testing-colly/internal/domain"
	"github.com/ferminhg/testing-colly/internal/infra/website_scrapper"
	"log"
)

func main() {
	fmt.Println("🦀 scraping tech blogs 🦀")
	strategies := []domain.SiteStrategy{
		website_scrapper.NewMartinFowlerStrategy(),
	}
	scrapper := application.NewScrapper(strategies)
	if err := scrapper.Run(); err != nil {
		log.Fatal(err)
	}
}
