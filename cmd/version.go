package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Scraper",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Scrapper v.0.0.1 -- HEAD")
	},
}
