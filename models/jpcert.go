package models

type ArticleID string // {year}2{id}6 の形式

// CVE IDをキーにしたデータ, 関連する警戒情報を保存する
type CveID string
type JpcertCveKeyMap map[CveID][]ArticleID

// ArticleIDをキーにしたデータ, CVEが存在しないArticleを特定し対策を取る
type JpcertArticleKeyMap map[ArticleID][]CveID

// アプリとしては利用しない予定
type JpcertArticle struct {
	ArticleID   string
	Title       string
	Body        string `gorm:"type:text;"`
	PublishDate string //time.Time
	Cves        []JpcertCve
}

type JpcertCve struct {
	ID        int64
	CveID     string `gorm:"index"`
	ArticleID string `gorm:"index"`
}
