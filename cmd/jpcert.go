package cmd

import (
	"github.com/inconshreveable/log15"
	"github.com/spf13/cobra"
	"github.com/tomoyamachi/gocarts/fetcher"
)

// debianCmd represents the debian command
var jpcertCmd = &cobra.Command{
	Use:   "jpcert",
	Short: "Fetch alerts from JPCERT",
	Long:  `Fetch alerts from JPCERT`,
	RunE:  fetchJP,
}

func init() {
	fetchCmd.AddCommand(jpcertCmd)
}

func fetchJP(cmd *cobra.Command, args []string) (err error) {
	log15.Info("Fetched alerts from JPCERT")
	alerts, err := fetcher.RetrieveJPCERT()

	log15.Info("Fetched", "Alerts", len(alerts))
	log15.Info("data", alerts)

	return nil
}
