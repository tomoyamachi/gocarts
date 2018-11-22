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
func RetrieveJpcert(after int) (articles []models.JpcertAlert, err error) {
	thisYear := time.Now().Year()
	// up to current year
	for year := after; year <= thisYear; year++ {
		// fetch alart pages
		alerts, _ := retrieveYearJpcert(year)
		for seqId, txt := range alerts {
			cveIDs := findCveIDs(txt)
			if dateString, title, err := detectEachPart(txt); err == nil {
				date, _ := time.Parse("2006-01-02", dateString)
				articleID := uint(year*10000 + seqId)
				url := generateUrl(year, seqId)
				articles = append(
					articles,
					models.JpcertAlert{
						AlertID:     articleID,
						Title:       title,
						URL:         url,
						PublishDate: date,
						JpcertCves:  convertCveIDsToCve(articleID, cveIDs),
					},
				)
			}

		}
	}

	return articles, nil
}

func convertCveIDsToCve(articleID uint, cveIDs []string) (cves []models.JpcertCve) {
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

func findCveIDs(body string) (cveIDs []string) {
	rawMatches := cvePattern.FindAllString(body, -1)
	matches := util.UniqueStrings(rawMatches)
	for _, cveID := range matches {
		cveIDs = append(cveIDs, cveID)
	}
	return cveIDs
}

func generateUrl(year int, id int) (url string) {
	articleID := fmt.Sprintf("%02d%04d", year%100, id)
	url = fmt.Sprintf("https://www.jpcert.or.jp/at/%d/at%s.txt", year, articleID)
	return url
}

func retrieveYearJpcert(year int) (alertBodies map[int]string, err error) {
	// count up when doesn't exist text data
	alertBodies = map[int]string{}
	continueDontExist := 0

	// 連続して10回リンクがなければ、その年は終了
	for seqId := 1; continueDontExist < 10; seqId++ {
		url := generateUrl(year, seqId)
		log15.Info("Fetching", "URL", url)
		text, err := util.FetchURL(url)

		// return error if HTTP status not 404
		if err != nil {
			continueDontExist++
		} else {
			// convert ISO-2022-JP to UTF-8
			if alertBodies[seqId], err = util.FromISO2022JP(string(text)); err != nil {
				log15.Error("something occured ", "ERR", err)
			} else {
				continueDontExist = 0
			}
		}
	}
	return alertBodies, nil
}
