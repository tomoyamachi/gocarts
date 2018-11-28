package db

import (
	"fmt"
	"github.com/inconshreveable/log15"
	"github.com/jinzhu/gorm"
	"github.com/tomoyamachi/gocarts/models"
	"gopkg.in/cheggaaa/pb.v1"
	"time"
)

func (r *RDBDriver) deleteAndInsertJpcert(conn *gorm.DB, alert models.Alert) (err error) {

	tx := conn.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	// Delete old records if found
	old := models.Alert{}
	result := tx.Where("alert_id = ?", alert.AlertID).First(&old)
	if !result.RecordNotFound() {
		// Delete old records
		var errs gorm.Errors
		errs = errs.Add(
			tx.Where("alert_id = ?", alert.AlertID).Delete(models.Cve{}).Error,
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

// Fecth alerts by CVE-ID and Team
func (r *RDBDriver) GetAlertsByCveId(cveId string) (allAlerts []models.Alert, err error) {
	// Fetch target id's reference data
	refs := []models.Cve{}
	if err = r.conn.Where("cve_id = ?", cveId).Find(&refs).Error; err != nil {
		return nil, err
	}
	if len(refs) == 0 {
		return allAlerts, nil
	}

	// Fetch from reference ids

	// aggregate alert ids
	alertIds := []uint{}
	for _, ref := range refs {
		alertIds = append(alertIds, ref.AlertID)
	}

	// get alerts from alert id
	all := []models.Alert{}
	if err = r.conn.Where("alert_id in (?)", alertIds).Find(&all).Error; err != nil {
		return nil, err
	}

	if allAlerts, err = appendRelatedCves(r, all); err != nil {
		return nil, err
	}
	return allAlerts, nil
}

// Fetch alerts by published date
func (r *RDBDriver) GetAfterTimeJpcert(after time.Time) (allAlerts []models.Alert, err error) {
	all := []models.Alert{}
	if err = r.conn.Where("publish_date >= ?", after.Format("2006-01-02")).Find(&all).Error; err != nil {
		return nil, err
	}

	if allAlerts, err = appendRelatedCves(r, all); err != nil {
		return nil, err
	}
	return allAlerts, nil
}

func appendRelatedCves(r *RDBDriver, alerts []models.Alert) (allAlerts []models.Alert, err error) {
	for _, a := range alerts {
		cves := []models.Cve{}
		//r.conn.Model(&a).Related(&a.Cves)  TODO : it didn't work.
		if err = r.conn.Where("alert_id = ?", a.AlertID).Find(&cves).Error; err != nil {
			return nil, err
		}
		a.Cves = cves
		allAlerts = append(allAlerts, a)
	}
	return allAlerts, nil

}

func (r *RDBDriver) InsertAlert(alerts []models.Alert) (err error) {
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
