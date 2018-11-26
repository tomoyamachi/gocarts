package models

import (
	"time"
)

type Alert struct {
	AlertID     uint `gorm:"primary_key"`
	Team        string
	Title       string
	URL         string
	PublishDate time.Time
	Cves        []Cve `gorm:"foreignkey:AlertID;association_foreignkey:AlertID"`
}

type Cve struct {
	ID      uint   `gorm:"primary_key"`
	CveID   string `gorm:"index"`
	AlertID uint   `gorm:"index"`
}

const TEAM_JPCERT = "jp"
const TEAM_USCERT = "us"

var TEAM_PREFIX_ID = map[string]int{
	TEAM_JPCERT: 1,
	TEAM_USCERT: 2,
}
