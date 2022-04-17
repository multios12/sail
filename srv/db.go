package main

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
func dbOpen() (err error) {
	filename := filepath.Join(dataPath, "db")
	db, err = gorm.Open(sqlite.Open(filename), &gorm.Config{})
	db.AutoMigrate(&Cost{})
	return
}

/** 支出の追加・更新 */
func upsertCost(cost Cost) {
	db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&cost)
}

/** 年月をキーとして支出の検索 */
func findCostByMonth(month string) (cost Cost, err error) {
	if month, err := strconv.Atoi(month); err != nil {
		return cost, err
	} else {
		db.Find(&cost, month)
		return cost, err
	}
}

/** 年をキーとして支出の検索、テーブルに存在しない月は空データを返す */
func findCostByYear(year string) (costs []Cost, err error) {
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
				var c = Cost{Month: strconv.Itoa(i)}
				costs = append(costs, c)
			}
		}
		sort.Slice(costs, func(i, j int) bool { return costs[i].Month > costs[j].Month })
		return costs, e
	}
}
