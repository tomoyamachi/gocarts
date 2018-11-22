package models

// アプリとしては利用しない予定
type JpcertAlert struct {
	AlertID     string `gorm:"index"`
	Title       string
	Body        string      `gorm:"type:text;"`
	PublishDate string      //time.Time
	JpcertCves  []JpcertCve `gorm:"foreignkey:AlertID;association_foreignkey:AlertID"`
}

type JpcertCve struct {
	ID      uint   `gorm:"primary_key"`
	CveID   string `gorm:"index"`
	AlertID string `gorm:"index"`
}
