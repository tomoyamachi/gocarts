package fetcher

import (
	"fmt"
	"github.com/inconshreveable/log15"
	"github.com/tomoyamachi/gocarts/models"
	"github.com/tomoyamachi/gocarts/util"
	"regexp"
)

var re = regexp.MustCompile(`CVE-[0-9]+-[0-9]+`)

// https://security-tracker.debian.org/tracker/data/json
func RetrieveJPCERT() (cves models.JpcertCveKeyMap, err error) {
	cves = models.JpcertCveKeyMap{} //map[models.CveID][]models.ArticleID{}
	// TODO : 指定した年のURLをfor分で回して取得
	alerts, _ := retrieveYearJPCERT(2017)

	for articleID, body := range alerts {
		addArticleIDtoCveMap(cves, models.ArticleID(articleID), body)
	}
	fmt.Printf("%q", cves)
	return cves, nil
}

// TODO : ポインタ渡しのほうがいいのでは
func addArticleIDtoCveMap(cves models.JpcertCveKeyMap, articleID models.ArticleID, body string) models.JpcertCveKeyMap {
	rawMatches := re.FindAllString(body, -1)
	matches := util.RemoveDuplicateFromSlice(rawMatches)
	for _, cveID := range matches {
		cves[models.CveID(cveID)] = append(
			returnOrCreateCveSlice(cves, models.CveID(cveID)),
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
