package converter

// 給与・賞与支給明細書モデル
type SalaryDetail struct {
	Month    string          // 年月
	Title    string          // タイトル
	IsError  bool            // 取得エラー
	Counts   []DetailItem    // 日数配列
	Times    []TimeItem      // 時間配列
	Salarys  []DetailItem    // 支給配列
	Costs    []DetailItem    // 控除配列
	Totals   []DetailItem    // 合計配列
	Expense  int             // 経費等支給合計額
	Expenses []ExpenseDetail // 経費内訳配列
	Images   []string        // 画像ファイル配列
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

// ----------------------------------------------------------------------------

// 経費明細モデル
type ExpenseDetail struct {
	Month   string       // 年月
	Expense int          // 経費等支給合計額
	Amounts []AmountItem // 経費内訳配列
}

// 経費項目
type AmountItem struct {
	Name   string // 項目名
	Amount int    // 金額
	Memo   string // 備考
}
