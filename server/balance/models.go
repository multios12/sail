package balance

import (
	"time"
)

// バランスシート(年単位集計)モデル
type BalanceYear struct {
	Year        string    // 年
	EnableYears []string  // 利用可能な年のリスト
	Balances    []Balance // バランス配列
}

// バランスシート(Balanceレコード)
type Balance struct {
	Month   string `gorm:"primaryKey"` // 年月(yyyyMM)
	Salary  int    // 総支給額
	Paid    int    // 差引支給額
	Expense int    // 経費支給額

	Cost         int  `gorm:"default:0;"` // 固定支出総額
	IsNotCost    bool `gorm:"default:0;"` // 固定支出額未入力あり
	CostHousing  int  `gorm:"default:0;"` // 費用：住宅
	CostWater    int  `gorm:"default:0;"` // 費用：水道
	CostElectric int  `gorm:"default:0;"` // 費用：電気
	CostGas      int  `gorm:"default:0;"` // 費用：ガス
	CostMobile   int  `gorm:"default:0;"` // 費用：携帯
	CostLine     int  `gorm:"default:0;"` // 費用：通信
	CostTax      int  `gorm:"default:0;"` // 費用：納税

	Saving    int       `gorm:"default:0;"` // 貯蓄額
	Memo      string    // メモ
	CreatedAt time.Time // 作成時に値がゼロ値の場合、現在時間がセットされる
	UpdatedAt time.Time // 更新時、または作成時の値がゼロ値の場合、現在のUNIX秒がセットされる
}

// ----------------------------------------------------------------------------

// 給与支給明細書（年単位集計）モデル
type SalaryYear struct {
	Year        string       // 年
	EnableYears []string     // 利用可能な年のリスト
	Totals      []DetailItem // 合計リスト
	Details     []Salary     // 月ごとの給与支給明細リスト
}

// 給与支給明細書モデル
type Salary struct {
	Month    string        // 年月
	Title    string        // タイトル
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
