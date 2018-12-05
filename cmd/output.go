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

// outputCmd represents output golang file for future-architect/vuls
var outputCmd = &cobra.Command{
	Use:   "output",
	Short: "Output alerts to stdout",
	Long:  `Output alerts to stdout`,
	RunE:  outputData,
}

func init() {
	RootCmd.AddCommand(outputCmd)
	outputCmd.PersistentFlags().String("team", "jp", "Output data from XX-CERT(e.g. jp, us)")
	viper.BindPFlag("team", outputCmd.PersistentFlags().Lookup("team"))
	viper.SetDefault("team", "jp")

	outputCmd.PersistentFlags().String("output-type", "cve", "Output data type(e.g. alert, cve)")
	viper.BindPFlag("output-type", outputCmd.PersistentFlags().Lookup("output-type"))
	viper.SetDefault("output-type", "cve")
}

func outputData(cmd *cobra.Command, args []string) (err error) {
	driver, locked, err := db.NewDB(viper.GetString("dbtype"), viper.GetString("dbpath"), viper.GetBool("debug-sql"))
	if err != nil {
		if locked {
			log15.Error("Failed to initialize DB. Close DB connection before fetching", "err", err)
		}
		return err
	}
	outputType := viper.GetString("output-type")
	team := viper.GetString("team")

	var code string
	if outputType == "cve" {
		code, err = outputCveDict(driver, team)
	} else if outputType == "alert" {
		code, err = outputAlertDict(driver, team)
	} else {
		err = fmt.Errorf("output-type error : %s", outputType)
	}
	if err != nil {
		log15.Error("Failed to output file.", "err", err)
		return err
	}

	fmt.Println(code)
	return nil
}

func outputCveDict(driver db.DB, team string) (code string, err error) {
	alerts, _ := driver.GetAllAlertsCveIdKeyByTeam(team)
	if team == models.TEAM_JPCERT {
		code, err = output.GenerateCveDictJP(alerts)
	} else if team == models.TEAM_USCERT {
		code, err = output.GenerateCveDictUS(alerts)
	}
	return code, err
}

func outputAlertDict(driver db.DB, team string) (code string, err error) {
	alerts, _ := driver.GetTargetTeamAlerts(team)
	if team == models.TEAM_JPCERT {
		code, err = output.GenerateAlertDictJP(alerts)
	} else if team == models.TEAM_USCERT {
		code, err = output.GenerateAlertDictUS(alerts)
	}
	return code, err
}
