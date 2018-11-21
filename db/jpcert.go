package db

import (
	"fmt"
	"github.com/inconshreveable/log15"
	"github.com/jinzhu/gorm"
	"github.com/tomoyamachi/gocarts/models"
	pb "gopkg.in/cheggaaa/pb.v1"
)

func (r *RDBDriver) deleteAndInsertJpcert(conn *gorm.DB, article models.JpcertArticle) (err error) {

	tx := conn.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	// Delete old records if found
	old := models.JpcertArticle{}
	result := tx.Where("article_id = ?", article.ArticleID).First(&old)
	if !result.RecordNotFound() {
		// Delete old records
		var errs gorm.Errors
		errs = errs.Add(
			tx.Where("article_id = ?", article.ArticleID).Delete(models.JpcertCve{}).Error,
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
				article.ArticleID, errs.Error())
		}
	}
	if err = tx.Create(&article).Error; err != nil {
		return err
	}
	return nil
}

func (r *RDBDriver) InsertJpcert(articles []models.JpcertArticle) (err error) {
	bar := pb.StartNew(len(articles))
	log15.Info(fmt.Sprintf("Insert %d articles", len(articles)))
	for _, article := range articles {
		if err := r.deleteAndInsertJpcert(r.conn, article); err != nil {
			return fmt.Errorf("Failed to insert. article: %s, err: %s",
				article.ArticleID, err)
		}
		bar.Increment()
	}
	bar.Finish()
	return nil
}
