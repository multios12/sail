package balance

import (
	"encoding/json"
	"time"

	"github.com/multios12/sail/pkg/balance/converter"
	"gorm.io/datatypes"
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

	Saving int    `gorm:"default:0;"` // 貯蓄額
	Memo   string // メモ

	CreatedAt time.Time // 作成時に値がゼロ値の場合、現在時間がセットされる
	UpdatedAt time.Time // 更新時、または作成時の値がゼロ値の場合、現在のUNIX秒がセットされる
}

// 給与明細
func (b *Balance) SalaryDetail() converter.SalaryDetail {
	d, _ := findBalanceDetailByMonthType(b.Month, DetailTypeSalary)
	if len(d.Json) == 0 {
		return converter.SalaryDetail{}
	} else {
		var detail converter.SalaryDetail
		json.Unmarshal(d.Json, &detail)
		detail.Images = []string{}
		detail.Images = append(detail.Images, "1.webp")
		if len(b.Image(DetailTypeExpense)) > 0 {
			detail.Images = append(detail.Images, "3.webp")
		}
		return detail
	}
}

// 賞与明細
func (b *Balance) BonusDetail() converter.SalaryDetail {
	d, _ := findBalanceDetailByMonthType(b.Month, DetailTypeBonus)
	if len(d.Json) == 0 {
		return converter.SalaryDetail{}
	} else {
		var detail converter.SalaryDetail
		json.Unmarshal(d.Json, &detail)
		detail.Images = []string{}
		detail.Images = append(detail.Images, "2.webp")
		return detail
	}
}

// 明細画像
func (b *Balance) Image(detailType DetailType) []byte {
	month := b.Month
	if len(b.Month) > 6 {
		month = b.Month[:6]
	}
	detail, _ := findBalanceDetailByMonthType(month, detailType)
	return detail.Image
}

// 明細タイプ
type DetailType int

const (
	DetailTypeSalary  DetailType = 1 // 明細タイプ：給与明細
	DetailTypeBonus   DetailType = 2 // 明細タイプ：賞与明細
	DetailTypeExpense DetailType = 3 // 明細タイプ：経費明細
)

// 明細データ
type BalanceDetail struct {
	ID    uint           `gorm:"primaryKey" ` // id
	Month string         // 年月(yyyyMM)
	Type  DetailType     `gorm:"default:1;"` // 種別
	Json  datatypes.JSON // JSONデータ
	Image []byte         // 画像データ
}

// 給与支給明細のキーデータ
type KeySalary struct {
	Month string // 年月(yyyyMM)
}

// ----------------------------------------------------------------------------

// 給与支給明細書（年単位集計）モデル
type SalaryYear struct {
	Year        string                   // 年
	EnableYears []string                 // 利用可能な年のリスト
	Totals      []converter.DetailItem   // 合計リスト
	Details     []converter.SalaryDetail // 月ごとの給与支給明細リスト
}
