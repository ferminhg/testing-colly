package cmd

import (
	"fmt"
	"log"

	"github.com/ferminhg/testing-colly/internal/application"
	"github.com/ferminhg/testing-colly/internal/domain"
	"github.com/ferminhg/testing-colly/internal/infra/storage/inmemory"
	"github.com/ferminhg/testing-colly/internal/infra/website_scrapper"
	"github.com/spf13/cobra"
)

func scrapperCmd(cmd *cobra.Command, args []string) {
	dryRun, err := cmd.Flags().GetBool("dry-run")
	if err != nil {
		log.Fatal(err)
	}

	scrapper(dryRun)
}

func scrapper(dryRun bool) {
	fmt.Println("🦀 scraping tech blogs 🦀")
	postRepository := inmemory.NewPostRepository()

	strategies := []domain.SiteStrategy{
		website_scrapper.NewMartinFowlerStrategy(postRepository, dryRun),
	}
	scrapper := application.NewScrapper(strategies)

	if err := scrapper.Run(); err != nil {
		log.Fatal(err)
	}
}
