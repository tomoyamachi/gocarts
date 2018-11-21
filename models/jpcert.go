package models

// アプリとしては利用しない予定
type JpcertArticle struct {
	ArticleID   string `gorm:"index"`
	Title       string
	Body        string      `gorm:"type:text;"`
	PublishDate string      //time.Time
	JpcertCves  []JpcertCve `gorm:"foreignkey:ArticleID;association_foreignkey:ArticleID"`
}

type JpcertCve struct {
	ID        uint   `gorm:"primary_key"`
	CveID     string `gorm:"index"`
	ArticleID string `gorm:"index"`
}
