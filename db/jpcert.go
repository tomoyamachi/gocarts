package db

import (
	"fmt"
	"github.com/inconshreveable/log15"
	"github.com/jinzhu/gorm"
	"github.com/tomoyamachi/gocarts/models"
	pb "gopkg.in/cheggaaa/pb.v1"
)

func (r *RDBDriver) deleteAndInsertJpcert(conn *gorm.DB, alert models.JpcertAlert) (err error) {

	tx := conn.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	// Delete old records if found
	old := models.JpcertAlert{}
	result := tx.Where("alert_id = ?", alert.AlertID).First(&old)
	if !result.RecordNotFound() {
		// Delete old records
		var errs gorm.Errors
		errs = errs.Add(
			tx.Where("alert_id = ?", alert.AlertID).Delete(models.JpcertCve{}).Error,
		)
		errs = errs.Add(tx.Unscoped().Delete(&old).Error)

		// Delete nil in errs
		var validErrs []error
		for _, err := range errs {
			if err != nil {
				validErrs = append(validErrs, err)
			}
		}
		errs = validErrs

		if len(errs.GetErrors()) > 0 {
			return fmt.Errorf("Failed to delete old records. id: %s, err: %s",
				alert.AlertID, errs.Error())
		}
	}
	if err = tx.Create(&alert).Error; err != nil {
		return err
	}
	return nil
}

func (r *RDBDriver) InsertJpcert(alerts []models.JpcertAlert) (err error) {
	bar := pb.StartNew(len(alerts))
	log15.Info(fmt.Sprintf("Insert %d alerts", len(alerts)))
	for _, alert := range alerts {
		if err := r.deleteAndInsertJpcert(r.conn, alert); err != nil {
			return fmt.Errorf("Failed to insert. alert: %s, err: %s",
				alert.AlertID, err)
		}
		bar.Increment()
	}
	bar.Finish()
	return nil
}
