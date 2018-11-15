package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:           "gocarts",
	Short:         "X-CERT alerts summarizer",
	Long:          `X-CERT alerts summarizer`,
	SilenceErrors: true,
	SilenceUsage:  true,
}

func init() {
	RootCmd.PersistentFlags().Bool("debug", false, "debug mode (default: false)")
	viper.BindPFlag("debug", RootCmd.PersistentFlags().Lookup("debug"))
	viper.SetDefault("debug", false)

	RootCmd.PersistentFlags().Bool("debug-sql", false, "SQL debug mode")
	viper.BindPFlag("debug-sql", RootCmd.PersistentFlags().Lookup("debug-sql"))
	viper.SetDefault("debug-sql", false)

	RootCmd.PersistentFlags().String("dbpath", "", "/path/to/sqlite3 or SQL connection string")
	viper.BindPFlag("dbpath", RootCmd.PersistentFlags().Lookup("dbpath"))
	pwd := os.Getenv("PWD")
	viper.SetDefault("dbpath", filepath.Join(pwd, "gost.sqlite3"))

	RootCmd.PersistentFlags().String("dbtype", "", "Database type to store data in (sqlite3, mysql or postgres supported)")
	viper.BindPFlag("dbtype", RootCmd.PersistentFlags().Lookup("dbtype"))
	viper.SetDefault("dbtype", "sqlite3")

	RootCmd.PersistentFlags().String("http-proxy", "", "http://proxy-url:port (default: empty)")
	viper.BindPFlag("http-proxy", RootCmd.PersistentFlags().Lookup("http-proxy"))
	viper.SetDefault("http-proxy", "")
}
