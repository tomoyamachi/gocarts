package fetcher

import (
	"fmt"
	"github.com/inconshreveable/log15"
	"github.com/tomoyamachi/gocarts/models"
	"github.com/tomoyamachi/gocarts/util"
	"regexp"
)

// https://security-tracker.debian.org/tracker/data/json
func RetrieveJPCERT() ([]models.JpcertArticle, error) {
	articles := []models.JpcertArticle{}

	// TODO : 指定した年のURLをfor分で回して取得
	alerts, _ := retrieveYearJPCERT(2018)
	for articleID, txt := range alerts {
		//cveIDs := findCveIDs(txt)
		date, title := detectEachPart(txt)
		articles = append(
			articles,
			models.JpcertArticle{
				ArticleID:   articleID,
				Title:       title,
				Body:        txt,
				PublishDate: date,
				Cves:        []models.JpcertCve{},
			},
		)

	}
	fmt.Printf("%v", articles)
	return articles, nil
}

var datePattern = regexp.MustCompile(`JPCERT/CC Alert (?P<date>\d{4}-\d{2}-\d{2})\s*>>>\s*(?P<title>.*)`)

func detectEachPart(txt string) (string, string) {
	matches := datePattern.FindStringSubmatch(txt)
	return matches[1], matches[2]
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
	for seqId := 40; seqId < 42; seqId++ {
		//for seqId := 1; continueDontExist < 1; seqId++ {
		articleID := fmt.Sprintf("%d%04d", year%100, seqId)
		url := fmt.Sprintf("https://www.jpcert.or.jp/at/%d/at%s.txt", year, articleID)
		log15.Info("Fetching", "URL", url)
		text, err := util.FetchURL(url)

		// TODO : return error if HTTP status not 404
		if err != nil {
			continueDontExist++
		} else {
			// convert ISO-2022-JP to UTF-8
			alertBodies[articleID], err = util.FromISO2022JP(string(text))
			if err != nil {
				panic(err)
			}
			continueDontExist = 0
		}
	}
	return alertBodies, nil
}
