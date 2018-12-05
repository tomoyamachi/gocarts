package db

import (
	"fmt"
	"github.com/inconshreveable/log15"
	"github.com/tomoyamachi/gocarts/models"
	"time"
)

// DB is interface for a database driver
type DB interface {
	Name() string
	OpenDB(string, string, bool) (bool, error)
	MigrateDB() error
	InsertAlert([]models.Alert) error
	GetAfterTimeAlerts(time.Time) ([]models.Alert, error)
	GetTargetTeamAlerts(string) ([]models.Alert, error)
	GetAlertsByCveId(string) ([]models.Alert, error)
	GetAllAlertsCveIdKeyByTeam(string) (map[string][]models.Alert, error)
}

// NewDB returns db driver
func NewDB(dbType, dbPath string, debugSQL bool) (driver DB, locked bool, err error) {
	if driver, err = newDB(dbType); err != nil {
		log15.Error("Failed to new db.", "err", err)
		return driver, false, err
	}

	if locked, err := driver.OpenDB(dbType, dbPath, debugSQL); err != nil {
		log15.Error("Failed to open db.", "err", err)
		if locked {
			log15.Error("db locked.", "err", err)
			return nil, true, err
		}
		return nil, false, err
	}

	if err := driver.MigrateDB(); err != nil {
		log15.Error("Failed to migrate db.", "err", err)
		return driver, false, err
	}
	return driver, false, nil
}

func newDB(dbType string) (DB, error) {
	switch dbType {
	case dialectSqlite3, dialectMysql, dialectPostgreSQL:
		return &RDBDriver{name: dbType}, nil
		//case dialectRedis:
		//	return &RedisDriver{name: dbType}, nil
	}
	return nil, fmt.Errorf("Invalid database dialect. err: %s", dbType)
}
