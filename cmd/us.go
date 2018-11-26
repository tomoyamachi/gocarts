package cmd

import (
	"github.com/inconshreveable/log15"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tomoyamachi/gocarts/db"
	"github.com/tomoyamachi/gocarts/fetcher"
)

// debianCmd represents the debian command
var uscertCmd = &cobra.Command{
	Use:   "us",
	Short: "Fetch alerts from US-CERT",
	Long:  `Fetch alerts from US-CERT`,
	RunE:  fetchUscert,
}

func init() {
	FetchCmd.AddCommand(uscertCmd)
}

func fetchUscert(cmd *cobra.Command, args []string) (err error) {
	log15.Info("Initialize Database")
	driver, locked, err := db.NewDB(viper.GetString("dbtype"), viper.GetString("dbpath"), viper.GetBool("debug-sql"))

	if err != nil {
		if locked {
			log15.Error("Failed to initialize DB. Close DB connection before fetching", "err", err)
		}
		return err
	}

	log15.Info("Fetched alerts from US-CERT")
	alerts, err := fetcher.RetrieveUscert(viper.GetInt("after"))

	log15.Info("Insert article into DB", "db", driver.Name())
	if err := driver.InsertAlert(alerts); err != nil {
		log15.Error("Failed to insert.", "dbpath", viper.GetString("dbpath"), "err", err)
		return err
	}
	return nil
}
