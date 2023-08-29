package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "scraper",
	Short: "Scraper is a CLI for scraping tech blogs",
	Run:   scrapperCmd,
}

var DryRun bool

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.PersistentFlags().BoolVarP(
		&DryRun,
		"dry-run",
		"d",
		false,
		`dry run mode: no data will be saved and only 10 posts will be scraped`,
	)

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
