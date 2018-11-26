package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// fetchCmd represents the fetch command
var FetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch X-CERT alerts",
	Long:  `Fetch X-CERT alerts`,
}

func init() {
	RootCmd.AddCommand(FetchCmd)
	FetchCmd.PersistentFlags().Int("after", 0, "Fetch articles after the specified year (e.g. 2017) (default: 1900)")
	viper.BindPFlag("after", FetchCmd.PersistentFlags().Lookup("after"))
	viper.SetDefault("after", "1990")

	FetchCmd.PersistentFlags().Int("wait", 0, "Interval between fetch (seconds)")
	viper.BindPFlag("wait", FetchCmd.PersistentFlags().Lookup("wait"))
	viper.SetDefault("wait", 0)

	FetchCmd.PersistentFlags().Int("threads", 5, "The number of threads to be used")
	viper.BindPFlag("threads", FetchCmd.PersistentFlags().Lookup("threads"))
	viper.SetDefault("threads", 5)
}
