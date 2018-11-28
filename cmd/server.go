package cmd

import (
	"github.com/inconshreveable/log15"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tomoyamachi/gocarts/db"
	"github.com/tomoyamachi/gocarts/server"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start alert information HTTP server",
	Long:  `Start alert information HTTP server`,
	RunE:  executeServer,
}

func init() {
	RootCmd.AddCommand(serverCmd)

	serverCmd.PersistentFlags().String("bind", "", "HTTP server bind to IP address (default: loop back interface")
	viper.BindPFlag("bind", serverCmd.PersistentFlags().Lookup("bind"))
	viper.SetDefault("bind", "127.0.0.1")

	serverCmd.PersistentFlags().String("port", "", "HTTP server port number (default: 1235")
	viper.BindPFlag("port", serverCmd.PersistentFlags().Lookup("port"))
	viper.SetDefault("port", "1327")
}

func executeServer(cmd *cobra.Command, args []string) (err error) {
	logDir := viper.GetString("log-dir")
	driver, locked, err := db.NewDB(viper.GetString("dbtype"), viper.GetString("dbpath"), viper.GetBool("debug-sql"))
	if err != nil {
		if locked {
			log15.Error("Failed to initialize DB. Close DB connection before fetching", "err", err)
		}
		return err
	}

	log15.Info("Starting HTTP Server...")
	if err = server.Start(logDir, driver); err != nil {
		log15.Error("Failed to start server.", "err", err)
		return err
	}

	return nil
}
