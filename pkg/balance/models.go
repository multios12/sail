package balance

import (
	"encoding/json"
	"time"

	"github.com/multios12/sail/pkg/balance/converter"
	"github.com/multios12/sail/pkg/balance/models"
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
	d, _ := findBalanceDetailByMonthType(b.Month, models.DetailTypeSalary)
	if len(d.Json) == 0 {
		return converter.SalaryDetail{}
	} else {
		var detail converter.SalaryDetail
		json.Unmarshal(d.Json, &detail)
		detail.Images = []string{}
		detail.Images = append(detail.Images, "1.png")
		if len(b.Image(models.DetailTypeExpense)) > 0 {
			detail.Images = append(detail.Images, "3.png")
		}
		return detail
	}
}

// 賞与明細
func (b *Balance) BonusDetail() converter.SalaryDetail {
	d, _ := findBalanceDetailByMonthType(b.Month, models.DetailTypeBonus)
	if len(d.Json) == 0 {
		return converter.SalaryDetail{}
	} else {
		var detail converter.SalaryDetail
		json.Unmarshal(d.Json, &detail)
		detail.Images = []string{}
		detail.Images = append(detail.Images, "2.png")
		return detail
	}
}

// 明細画像
func (b *Balance) Image(detailType models.BalanceType) []byte {
	month := b.Month
	if len(b.Month) > 6 {
		month = b.Month[:6]
	}
	detail, _ := findBalanceDetailByMonthType(month, detailType)
	return detail.Image
}
