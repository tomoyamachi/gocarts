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
	fetchCmd.AddCommand(uscertCmd)
	uscertCmd.PersistentFlags().String("after", "", "Fetch articles after the specified year (e.g. 2017) (default: 2015)")
	viper.BindPFlag("after", uscertCmd.PersistentFlags().Lookup("after"))
	viper.SetDefault("after", "2015")
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
	articles, err := fetcher.RetrieveUscert(viper.GetInt("after"))

	log15.Info("Insert article into DB", "db", driver.Name())
	if err := driver.InsertJpcert(articles); err != nil {
		log15.Error("Failed to insert.", "dbpath", viper.GetString("dbpath"), "err", err)
		return err
	}
	log15.Info("articles : ", articles)

	return nil
}
