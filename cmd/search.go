package cmd

import (
	"github.com/spf13/cobra"
)

// fetchCmd represents the fetch command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search X-CERT alerts",
	Long:  `Search X-CERT alerts`,
}

func init() {
	RootCmd.AddCommand(searchCmd)

}
