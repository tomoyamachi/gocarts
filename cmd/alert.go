package cmd

import (
	"bytes"
	"fmt"
	"github.com/inconshreveable/log15"
	"github.com/mattn/go-runewidth"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tomoyamachi/gocarts/db"
	"github.com/tomoyamachi/gocarts/models"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

// debianCmd represents the debian command
var searchByAlertCmd = &cobra.Command{
	Use:   "alert",
	Short: "Search alerts by alert information",
	Long:  `Search alerts by alert information`,
	RunE:  searchAlert,
}

func init() {

	searchCmd.AddCommand(searchByAlertCmd)

	searchByAlertCmd.PersistentFlags().String("select-cmd", "", "Select command (default: fzf)")
	viper.BindPFlag("select-cmd", searchByAlertCmd.PersistentFlags().Lookup("select-cmd"))
	viper.SetDefault("select-cmd", "fzf")

	searchByAlertCmd.PersistentFlags().String("select-option", "", "Select command options")
	viper.BindPFlag("select-option", searchByAlertCmd.PersistentFlags().Lookup("select-option"))
	viper.SetDefault("select-option", "--reverse")

	searchByAlertCmd.PersistentFlags().String("select-after", "", "Show CVEs after the specified date (e.g. 2017-01-01) (default:  1 year ago)")
	viper.BindPFlag("select-after", searchByAlertCmd.PersistentFlags().Lookup("select-after"))
	viper.SetDefault("select-after", "")
}

func searchAlert(cmd *cobra.Command, args []string) (err error) {
	afterOption := viper.GetString("select-after")
	var after time.Time
	if afterOption != "" {
		if after, err = time.Parse("2006-01-02", afterOption); err != nil {
			return fmt.Errorf("Failed to parse --select-after. err: %s", err)
		}
	} else {
		now := time.Now()
		after = now.Add(time.Duration(-1) * 24 * 365 * time.Hour)
	}

	log15.Info("Initialize Database")
	driver, locked, err := db.NewDB(viper.GetString("dbtype"), viper.GetString("dbpath"), viper.GetBool("debug-sql"))
	if err != nil {
		if locked {
			log15.Error("Failed to initialize DB. Close DB connection before fetching", "err", err)
		}
		return err
	}

	log15.Info("Select all Alerts")
	allAlerts, err := driver.GetAfterTimeJpcert(after)
	if err != nil {
		return err
	}

	allAlertText := []string{}
	for _, alert := range allAlerts {
		allAlertText = append(
			allAlertText,
			fmt.Sprintf(
				"%s | %s | %-60s | %s",
				alert.PublishDate.Format("2006-01-02"),
				alert.URL,
				runewidth.Truncate(convertCvesToText(alert.JpcertCves), 60, "…"),
				alert.Title,
			),
		)
	}

	_, err = filter(allAlertText)

	return err
}

func convertCvesToText(cves []models.JpcertCve) (cveText string) {
	if len(cves) == 0 {
		return "..."
	}
	cveNames := []string{}
	for _, cve := range cves {
		cveNames = append(cveNames, cve.CveID)
	}
	return strings.Join(cveNames, ", ")
}

func run(command string, r io.Reader, w io.Writer) error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", command)
	} else {
		cmd = exec.Command("sh", "-c", command)
	}
	cmd.Stderr = os.Stderr
	cmd.Stdout = w
	cmd.Stdin = r
	return cmd.Run()
}

func filter(cves []string) (results []string, err error) {
	var buf bytes.Buffer
	selectCmd := fmt.Sprintf("%s %s",
		viper.GetString("select-cmd"), viper.GetString("select-option"))
	err = run(selectCmd, strings.NewReader(strings.Join(cves, "\n")), &buf)
	if err != nil {
		return nil, nil
	}

	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")

	return lines, nil
}
