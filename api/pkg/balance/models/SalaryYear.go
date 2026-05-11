package models

import "github.com/multios12/sail/pkg/balance/converter"

// 給与支給明細書（年単位集計）モデル
type SalaryYear struct {
	Year        string                   // 年
	EnableYears []string                 // 利用可能な年のリスト
	Totals      []converter.DetailItem   // 合計リスト
	Details     []converter.SalaryDetail // 月ごとの給与支給明細リスト
}
