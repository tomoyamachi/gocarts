package fetcher

import (
	"fmt"
	"github.com/inconshreveable/log15"
	"github.com/tomoyamachi/gocarts/models"
	"github.com/tomoyamachi/gocarts/util"
	"regexp"
)

// https://security-tracker.debian.org/tracker/data/json
func RetrieveJPCERT() (cves models.JpcertCveKeyMap, err error) {
	cves = models.JpcertCveKeyMap{} //map[models.CveID][]models.ArticleID{}
	// TODO : 指定した年のURLをfor分で回して取得
	alerts, _ := retrieveYearJPCERT(2018)
	for articleID, body := range alerts {
		cveIDs := findCveIDs(body)
		addArticleIDtoCveMap(cves, models.ArticleID(articleID), cveIDs)
	}
	fmt.Printf("%q", cves)
	return cves, nil
}

// return CVE slice mathed from alert's body
var cvePattern = regexp.MustCompile(`CVE-[0-9]+-[0-9]+`)

func findCveIDs(body string) []models.CveID {
	cveIDs := []models.CveID{}
	rawMatches := cvePattern.FindAllString(body, -1)
	matches := util.RemoveDuplicateFromSlice(rawMatches)
	for _, cveID := range matches {
		cveIDs = append(cveIDs, models.CveID(cveID))
	}
	return cveIDs
}

func addArticleIDtoCveMap(cves models.JpcertCveKeyMap, articleID models.ArticleID, cveIDs []models.CveID) models.JpcertCveKeyMap {
	for _, cveID := range cveIDs {
		cves[cveID] = append(
			returnOrCreateCveSlice(cves, cveID),
			articleID,
		)
	}
	return cves
}

// return empty slice if it doesnt exist
func returnOrCreateCveSlice(cves models.JpcertCveKeyMap, cveID models.CveID) []models.ArticleID {
	if cveData, ok := cves[cveID]; ok {
		return cveData
	}
	return []models.ArticleID{}
}

func retrieveYearJPCERT(year int) (alertBodies map[models.ArticleID]string, err error) {
	alertBodies = map[models.ArticleID]string{}
	// count up when doesn't exist text data
	continueDontExist := 0

	// 連続して10回リンクがなければ、その年は終了
	//for seqId := ; seqId < 42; seqId++ {
	for seqId := 1; continueDontExist < 1; seqId++ {
		articleID := fmt.Sprintf("%d%04d", year%100, seqId)
		url := fmt.Sprintf("https://www.jpcert.or.jp/at/%d/at%s.txt", year, articleID)
		log15.Info("Fetching", "URL", url)
		text, err := util.FetchURL(url)
		// TODO : return error if HTTP status not 404
		if err != nil {
			continueDontExist++
		} else {
			alertBodies[models.ArticleID(articleID)] = string(text)
			continueDontExist = 0
		}
	}
	return alertBodies, nil
}
