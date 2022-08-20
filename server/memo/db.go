package memo

import (
	"path/filepath"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var db *gorm.DB

func dbOpen(dataPath string) (err error) {
	filename := filepath.Join(dataPath, "hmemo")
	db, err = gorm.Open(sqlite.Open(filename), &gorm.Config{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&Memo{})
	return err
}

func find() (memos []Memo) {
	db.Find(&memos)
	return memos
}

func findById(id string) (memos []Memo) {
	db.Where("id = ?", id).Find(&memos)
	return memos
}

func findByMonth(month string) (memos []Memo) {
	from := month + "-01"
	toDate, _ := time.Parse("2006-01-02", from)
	toDate = toDate.AddDate(0, 1, 0).AddDate(0, 0, -1)
	to := toDate.Format("2006-01-02")
	db.Where("date >= ? and date <= ?", from, to).Find(&memos)
	return memos
}

func upsertMemo(m Memo) {
	db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&m)
}

func deleteMemo(id string) {
	db.Delete(&Memo{}, id)
}
