package models

// 明細タイプ
type BalanceType int

const (
	DetailTypeSalary  BalanceType = 1 // 明細タイプ：給与明細
	DetailTypeBonus   BalanceType = 2 // 明細タイプ：賞与明細
	DetailTypeExpense BalanceType = 3 // 明細タイプ：経費明細
)
