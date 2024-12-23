package balance

import (
	"encoding/json"
	"path/filepath"
	"sort"
	"strconv"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/multios12/sail/pkg/balance/converter"
	"github.com/multios12/sail/pkg/balance/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// データベース
var db *gorm.DB

// データベースオープン
func dbOpen(salaryPath string) (err error) {
	filename := filepath.Join(salaryPath, "balance.db")
	db, err = gorm.Open(sqlite.Open(filename), &gorm.Config{})
	if err != nil {
		return
	}

	// オートマイグレーション
	err = db.AutoMigrate(&Balance{})
	if err != nil {
		return
	}
	err = db.AutoMigrate(&models.BalanceDetail{})
	return
}

/** 年をキーとしてバランスシートデータを返す、テーブルに存在しない月は空データを返す */
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

/** 年月をキーとしてバランスシートデータを返す */
func findBalanceByMonth(month string) (b Balance, err error) {
	result := db.Find(&b, month)
	b.Month = month
	return b, result.Error
}

/**  バランスシートの追加・更新 */
func upsertBalance(b Balance) {
	db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&b)
}

// 支給額の再計算
func calculateSalary(month string) {
	details, _ := findBalanceDetailByMonth(month)
	salaryItem := converter.SalaryDetail{}
	bonusItem := converter.SalaryDetail{}
	expenseItem := converter.ExpenseDetail{}

	for _, detail := range details {
		if detail.Type == models.DetailTypeSalary {
			json.Unmarshal(detail.Json, &salaryItem)
		} else if detail.Type == models.DetailTypeBonus {
			json.Unmarshal(detail.Json, &bonusItem)
		} else if detail.Type == models.DetailTypeExpense {
			json.Unmarshal(detail.Json, &expenseItem)
		}
	}

	balance, _ := findBalanceByMonth(month)

	balance.Salary = 0
	balance.Paid = 0
	balance.Expense = expenseItem.Expense
	if len(salaryItem.Totals) >= 3 {
		balance.Salary = salaryItem.Totals[0].Value
		balance.Paid = salaryItem.Totals[2].Value
	}
	if len(bonusItem.Totals) >= 3 {
		balance.Salary += bonusItem.Totals[0].Value
		balance.Paid += bonusItem.Totals[2].Value
	}

	upsertBalance(balance)
}

func findSalaryYears() []string {
	rows, err := db.Table("balances").Select("month").Order("month desc").Rows()
	if err != nil {
		return []string{}
	}
	defer rows.Close()
	enableYears := []string{}
	for rows.Next() {
		var month string
		rows.Scan(&month)
		isEnabled := false
		for _, e := range enableYears {
			if e == month[:4] {
				isEnabled = true
			}
		}
		if !isEnabled {
			enableYears = append(enableYears, month[:4])
		}
	}
	return enableYears
}

// ----------------------------------------------------------------------------
func findBalanceDetailByMonth(month string) (details []models.BalanceDetail, err error) {
	result := db.Where("month = ?", month).Order("type").Find(&details)
	return details, result.Error
}

func findBalanceDetailByMonthType(month string, detailType models.BalanceType) (detail models.BalanceDetail, err error) {
	result := db.Where("month = ? and type = ?", month, detailType).Find(&detail)
	detail.Month = month
	detail.Type = detailType
	return detail, result.Error
}

// 給与明細データをもとにバランスシートを更新する
func upsertBalanceDetail(month string, detailType models.BalanceType, item []byte, image []byte) {
	b, _ := findBalanceDetailByMonthType(month, detailType)
	b.Month = month
	b.Json = item
	b.Image = image

	db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&b)
	calculateSalary(month)
}

func deleteBalanceDetail(month string) {
	db.Where("month = ?", month).Delete(&models.BalanceDetail{})
}

// ----------------------------------------------------------------------------

func FindUseById(id uint) (model models.UseDetail, err error) {
	result := db.Find(&model, id)
	return model, result.Error
}

func FindUseByPayMonth(payMonth string) (models []models.UseDetail, err error) {
	result := db.Find(&models, "payMonth = ?", payMonth)
	return models, result.Error
}

func UpsertUse(u []models.UseDetail) {
	db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&u)
}
