package fetcher

import (
	"fmt"
	"github.com/inconshreveable/log15"
	"github.com/tomoyamachi/gocarts/models"
	"github.com/tomoyamachi/gocarts/util"
)

// https://security-tracker.debian.org/tracker/data/json
func RetrieveJPCERT() (cves models.JpcertCveKeyMap, err error) {

	// TODO : 指定した年のURLをfor分で回して取得
	retrieveYearJPCERT(2018)

	return cves, nil
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
		// TODO : return error if HTTP status not 404
		if err != nil {
			continueDontExist++
		} else {
			alertBodies[articleID] = string(text)
			fmt.Print(string(text))
			continueDontExist = 0
		}
	}
	return alertBodies, nil
}
