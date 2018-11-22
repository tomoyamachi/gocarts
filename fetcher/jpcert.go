package fetcher

import (
	"errors"
	"fmt"
	"github.com/inconshreveable/log15"
	"github.com/tomoyamachi/gocarts/models"
	"github.com/tomoyamachi/gocarts/util"
	"regexp"
	"time"
)

// https://security-tracker.debian.org/tracker/data/json
func RetrieveJPCERT(after int) ([]models.JpcertAlert, error) {
	articles := []models.JpcertAlert{}
	thisYear := time.Now().Year()
	// up to current year
	for year := after; year <= thisYear; year++ {

		// fetch alart pages
		alerts, _ := retrieveYearJPCERT(year)
		for articleID, txt := range alerts {
			cveIDs := findCveIDs(txt)
			if date, title, err := detectEachPart(txt); err == nil {

				articles = append(
					articles,
					models.JpcertAlert{
						AlertID:     articleID,
						Title:       title,
						Body:        txt,
						PublishDate: date,
						JpcertCves:  convertCveIDsToCve(articleID, cveIDs),
					},
				)
			}

		}
	}

	return articles, nil
}

func convertCveIDsToCve(articleID string, cveIDs []string) []models.JpcertCve {
	cves := []models.JpcertCve{}
	for _, cveID := range cveIDs {
		cves = append(
			cves,
			models.JpcertCve{
				CveID:   cveID,
				AlertID: articleID,
			},
		)
	}
	return cves
}

var datePattern = regexp.MustCompile(`JPCERT/CC Alert (?P<date>\d{4}-\d{2}-\d{2})\s*>>>\s*(?P<title>.*)`)

func detectEachPart(txt string) (date string, title string, err error) {
	if matches := datePattern.FindStringSubmatch(txt); matches != nil {
		if len(matches) > 2 {
			return matches[1], matches[2], nil
		}
	}
	return "", "", errors.New("invalid text")
}

// return CVE slice mathed from alert's body
var cvePattern = regexp.MustCompile(`CVE-[0-9]+-[0-9]+`)

func findCveIDs(body string) []string {
	cveIDs := []string{}
	rawMatches := cvePattern.FindAllString(body, -1)
	matches := util.RemoveDuplicateFromSlice(rawMatches)
	for _, cveID := range matches {
		cveIDs = append(cveIDs, cveID)
	}
	return cveIDs
}

func retrieveYearJPCERT(year int) (alertBodies map[string]string, err error) {
	alertBodies = map[string]string{}
	// count up when doesn't exist text data
	continueDontExist := 0

	// 連続して10回リンクがなければ、その年は終了
	for seqId := 1; continueDontExist < 10; seqId++ {
		articleID := fmt.Sprintf("%d%04d", year%100, seqId)
		url := fmt.Sprintf("https://www.jpcert.or.jp/at/%d/at%s.txt", year, articleID)
		log15.Info("Fetching", "URL", url)
		text, err := util.FetchURL(url)

		// return error if HTTP status not 404
		if err != nil {
			continueDontExist++
		} else {
			// convert ISO-2022-JP to UTF-8
			if alertBodies[articleID], err = util.FromISO2022JP(string(text)); err != nil {
				log15.Error("something occured ", "ERR", err)
			} else {
				continueDontExist = 0
			}
		}
	}
	return alertBodies, nil
}
