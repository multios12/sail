package main

import (
	"time"
)

// 給与支給明細書（集計）モデル
type SalaryYearModel struct {
	Year        string             // 年
	EnableYears []string           // 利用可能な年のリスト
	Totals      []DetailItem       // 合計リスト
	Details     []SalaryMonthModel // 月ごとの給与支給明細リスト
}

// 給与支給明細書モデル
type SalaryMonthModel struct {
	Title    string        // タイトル
	Month    string        // 年月
	IsError  bool          // 取得エラー
	Counts   []DetailItem  // 日数配列
	Times    []TimeItem    // 時間配列
	Salarys  []DetailItem  // 支給配列
	Costs    []DetailItem  // 控除配列
	Totals   []DetailItem  // 合計配列
	Expense  int           // 経費等支給合計額
	Expenses []ExpenseItem // 経費内訳配列
	Images   []string      // 画像ファイル配列
}

// 詳細項目
type DetailItem struct {
	Name  string // 項目名
	Value int    // 金額
}

// 時間項目
type TimeItem struct {
	Name  string // 項目名
	Value string // 値
}

// 経費項目
type ExpenseItem struct {
	Name   string // 項目名
	Amount int    // 金額
	Memo   string // 備考
}

// 支出集計モデル
type SumCost struct {
	Year        string   // 年
	EnableYears []string // 利用可能な年のリスト
	Costs       []Cost   // 月ごとの支出配列
}

// 支出モデル
type Cost struct {
	Month     int       `gorm:"primaryKey"` // 年月(yyyyMM)
	Water     int       // 水道費
	Electric  int       // 電気費
	Gas       int       // ガス費
	Mobile    int       // 通信費(携帯)
	Line      int       // 通信費(固定)
	CreatedAt time.Time // 作成時に値がゼロ値の場合、現在時間がセットされる
	UpdatedAt time.Time // 更新時、または作成時の値がゼロ値の場合、現在のUNIX秒がセットされる
}
