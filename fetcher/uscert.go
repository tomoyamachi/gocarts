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

func RetrieveUscert(after int) (articles []models.Alert, err error) {
	thisYear := time.Now().Year()
	// up to current year
	for year := after; year <= thisYear; year++ {
		seqId := 0
		// fetch alart pages
		alerts, _ := retrieveYearUscert(year)
		for url, txt := range alerts {
			cveIDs := findCveIDs(txt)
			if dateString, title, err := detectUsEachPart(txt); err == nil {
				seqId++
				articleID := util.ReturnArticleID(year, seqId, models.TEAM_USCERT)
				date, _ := time.Parse("January 02, 2006", dateString)
				articles = append(
					articles,
					models.Alert{
						AlertID:     articleID,
						Team:        models.TEAM_USCERT,
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

var usDatePattern = regexp.MustCompile(`<footer class="submitted meta-text">Original release date: (?P<date>.*)</footer>`)
var usTitlePattern = regexp.MustCompile(`<h2 id="page-sub-title">(?P<title>.*)</h2>`)

func detectUsEachPart(txt string) (date string, title string, err error) {
	if dateMatch := usDatePattern.FindStringSubmatch(txt); dateMatch != nil {
		date = dateMatch[1]
	} else {
		log15.Error("Failed to parse", "date", "date")
		return "", "", errors.New("Cant detect date format")
	}
	if titleMatch := usTitlePattern.FindStringSubmatch(txt); titleMatch != nil {
		title = titleMatch[1]
	} else {
		log15.Error("Failed to parse", "title", "title")
		return "", "", errors.New("Cant detect title format")
	}
	return date, title, nil
}

func retrieveYearUscert(year int) (alertBodies map[string]string, err error) {
	alertBodies = map[string]string{}

	urls := findUrlsUscert(year)
	for _, url := range urls {
		log15.Info("Fetching", "URL", url)
		body, err := util.FetchURL(url)
		if err != nil {
			log15.Error("Failed to fetch", "URL", url)
			continue
		}
		alertBodies[url] = string(body)
	}
	return alertBodies, nil
}

var uscertUrlPattern = regexp.MustCompile(`/ncas/alerts/[A-Z]A\d{2}-\d{3}(A|B|C|D|E)`)

func findUrlsUscert(year int) (urls []string) {
	url := fmt.Sprintf("https://www.us-cert.gov/ncas/alerts/%d", year)
	body, err := util.FetchURL(url)
	if err != nil {
		log15.Error("Failed to fetch", "URL", url)
		return urls
	}
	rawMatches := uscertUrlPattern.FindAllString(string(body), -1)
	matches := util.UniqueStrings(rawMatches)
	for _, url := range matches {
		urls = append(urls, "https://www.us-cert.gov"+url)
	}
	return urls
}
