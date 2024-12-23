package models

import (
	"strconv"
	"strings"
)

// "ご利用者","カテゴリ","ご利用日","ご利用先など","ご利用金額(￥)","支払区分","今回回数","訂正サイン","お支払い金額(￥)","国内／海外","摘要","備考"
// "****-****-****-1234　ｘｘｘ　ｘｘ　ｘｘ様","≪ショッピング取組（国内）≫"," 2022/12/01","決済Ａ","275","１回","","","275","国内","","*"

// 利用明細
type UseDetail struct {
	ID       uint   `gorm:"primaryKey" ` // id
	Source   string // 発生元
	Type     string // 種別
	PayMonth string // 支払年月(yyyyMM)
	Date     string // 利用日
	Title    string // 利用先など
	Amount   int    // 利用金額
	Line     string // 明細
}

func CreateUseDetail(source string, month string, line string) UseDetail {
	u := UseDetail{Source: source, PayMonth: month, Line: line}
	values := strings.Split(line, `","`)
	for i := 0; i < len(values); i++ {
		values[i] = strings.ReplaceAll(values[i], `"`, "")
	}

	u.Date = values[2]
	u.Title = values[3]
	amount := strings.ReplaceAll(values[4], ",", "")
	u.Amount, _ = strconv.Atoi(amount)
	return u
}
