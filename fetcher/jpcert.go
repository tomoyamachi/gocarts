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

func RetrieveJpcert(after int) (articles []models.Alert, err error) {
	thisYear := time.Now().Year()
	// up to current year
	for year := after; year <= thisYear; year++ {
		// fetch alart pages
		alerts, _ := retrieveYearJpcert(year)
		for seqId, txt := range alerts {
			cveIDs := findCveIDs(txt)
			if dateString, title, err := detectEachPart(txt); err == nil {
				date, _ := time.Parse("2006-01-02", dateString)
				articleID := util.ReturnArticleID(year, seqId, models.TEAM_JPCERT)
				url := generateUrl(year, seqId)
				articles = append(
					articles,
					models.Alert{
						AlertID:     articleID,
						Team:        models.TEAM_JPCERT,
						Title:       title,
						URL:         url,
						PublishDate: date,
						Cves:        convertCveIDsToCve(articleID, cveIDs),
					},
				)
			}

		}
	}

	return articles, nil
}

func convertCveIDsToCve(articleID uint, cveIDs []string) (cves []models.Cve) {
	for _, cveID := range cveIDs {
		cves = append(
			cves,
			models.Cve{
				CveID:   cveID,
				AlertID: articleID,
			},
		)
	}
	return cves
}

var datePattern = regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
var titlePattern = regexp.MustCompile(`<h3>(?P<title>.*)</h3>`)

func detectEachPart(txt string) (date string, title string, err error) {
	if dateMatch := datePattern.FindStringSubmatch(txt); dateMatch != nil {
		date = dateMatch[0]
	} else {
		return "", "", errors.New("Cant detect date format")
	}
	if titleMatch := titlePattern.FindStringSubmatch(txt); titleMatch != nil {
		title = titleMatch[1]
	} else {
		return "", "", errors.New("Cant detect title format")
	}
	return date, title, nil
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

func generateUrl(year int, id int) string {
	// https://www.jpcert.or.jp/at/199x/99-0002-02.txt
	if year < 2000 {
		articleID := fmt.Sprintf("%02d-%04d-01", year%100, id)
		return fmt.Sprintf("https://www.jpcert.or.jp/at/199x/%s.html", year, articleID)
	}
	// https://www.jpcert.or.jp/at/2010/at100033.txt
	articleID := fmt.Sprintf("%02d%04d", year%100, id)
	return fmt.Sprintf("https://www.jpcert.or.jp/at/%d/at%s.html", year, articleID)
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
			alertBodies[seqId] = string(text)
			continueDontExist = 0
		}
	}
	return alertBodies, nil
}
