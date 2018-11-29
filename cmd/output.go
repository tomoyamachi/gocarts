package cmd

import (
	"fmt"
	"github.com/inconshreveable/log15"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tomoyamachi/gocarts/db"
	"github.com/tomoyamachi/gocarts/models"
	"github.com/tomoyamachi/gocarts/output"
)

// outputCmd represents the server command
var outputCmd = &cobra.Command{
	Use:   "output",
	Short: "Output alerts to stdout",
	Long:  `Output alerts to stdout`,
	RunE:  outputData,
}

func init() {
	RootCmd.AddCommand(outputCmd)
	outputCmd.PersistentFlags().String("team", "jp", "Output article from XX-CERT(e.g. jp, us)")
	viper.BindPFlag("team", outputCmd.PersistentFlags().Lookup("team"))
	viper.SetDefault("team", "jp")

}

func outputData(cmd *cobra.Command, args []string) (err error) {
	driver, locked, err := db.NewDB(viper.GetString("dbtype"), viper.GetString("dbpath"), viper.GetBool("debug-sql"))
	if err != nil {
		if locked {
			log15.Error("Failed to initialize DB. Close DB connection before fetching", "err", err)
		}
		return err
	}

	team := viper.GetString("team")
	alerts, _ := driver.GetAllAlertsCveIdKeyByTeam(team)

	var code string
	if team == models.TEAM_JPCERT {
		code, err = output.GenerateJP(alerts)
	} else if team == models.TEAM_USCERT {
		code, err = output.GenerateUS(alerts)
	}
	if err != nil {
		log15.Error("Failed to output file.", "err", err)
		return err
	}
	fmt.Println(code)
	return nil
}
