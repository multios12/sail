package main

// 給与支給明細書モデル
type DetailModel struct {
	Month   string       // 年月
	IsError bool         // 取得エラー
	Counts  []DetailItem // 日数配列
	Times   []TimeItem   // 時間配列
	Salarys []DetailItem // 支給配列
	Costs   []DetailItem // 控除配列
	Totals  []DetailItem // 合計配列
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

type YearModel struct {
	Year    string        // 年
	Totals  []DetailItem  // 合計リスト
	Details []DetailModel // 月ごとの給与支給明細リスト
}

// 経費等支給明細書モデル
type ExpenseModel struct {
	Code  string        //社員番号
	Name  string        // 指名
	Month string        // 年月
	Total int           // 経費等支給合計額
	Items []ExpenseItem // 項目(項目名,金額,備考)
}

// 経費項目
type ExpenseItem struct {
	Name   string // 項目名
	Amount int    // 金額
	Memo   string // 備考
}
