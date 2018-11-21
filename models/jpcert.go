package models

import "time"

type ArticleID string // {year}2{id}6 の形式

// CVE IDをキーにしたデータ, 関連する警戒情報を保存する
type CveID string
type JpcertCveKeyMap map[CveID][]ArticleID

// あとからどの警戒情報にCVE IDが含まれていないかをチェックする用
// アプリとしては利用しない予定
type JpcertAlertMap map[string]JpcertAlert
type JpcertAlert struct {
	ArticleID   string
	Title       string
	Body        string
	PublishDate time.Time
	CveIDs      []string
}
