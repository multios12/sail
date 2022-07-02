package memo

import (
	"path/filepath"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var db *gorm.DB

func dbOpen(dataPath string) (err error) {
	filename := filepath.Join(dataPath, "memo")
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

func upsertMemo(m Memo) {
	db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&m)
}

func deleteMemo(id string) {

}
