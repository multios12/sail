package balance

import (
	"path/filepath"
	"sort"
	"strconv"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// データベース
var db *gorm.DB

// データベースオープン
func dbOpen(dataPath string) (err error) {
	filename := filepath.Join(dataPath, "db")
	db, err = gorm.Open(sqlite.Open(filename), &gorm.Config{})
	db.AutoMigrate(&Balance{})
	return
}

/** 年をキーとしてバランスシートの検索、テーブルに存在しない月は空データを返す */
func findBalanceByYear(year string) (costs []Balance, err error) {
	if year, e := strconv.Atoi(year); e != nil {
		return nil, e
	} else {

		now := time.Now()
		lastMonth := 12
		if now.Year() == year && now.Month() < 12 {
			lastMonth = int(now.Month()) + 1
		}

		db.Where("month >= ? and month <= ?", year*100+1, year*100+lastMonth).Find(&costs)
		for i := year*100 + 1; i <= year*100+lastMonth; i++ {
			var exist = false
			for _, v := range costs {
				if v.Month == strconv.Itoa(i) {
					exist = true
					break
				}
			}
			if !exist {
				var c = Balance{Month: strconv.Itoa(i)}
				costs = append(costs, c)
			}
		}
		sort.Slice(costs, func(i, j int) bool { return costs[i].Month > costs[j].Month })
		return costs, e
	}
}

/** 年月をキーとしてバランスシートの検索 */
func findBalanceByMonth(month string) (b Balance, err error) {
	result := db.Find(&b, month)
	b.Month = month
	return b, result.Error

}

/**  バランスシートの追加・更新 */
func upsertBalance(b Balance) {
	db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&b)
}
