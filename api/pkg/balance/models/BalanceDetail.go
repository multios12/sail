package models

import "gorm.io/datatypes"

// 明細データ
type BalanceDetail struct {
	ID    uint           `gorm:"primaryKey" ` // id
	Month string         // 年月(yyyyMM)
	Type  BalanceType    `gorm:"default:1;"` // 種別
	Json  datatypes.JSON // JSONデータ
	Image []byte         // 画像データ
}

// 給与支給明細のキーデータ
type KeySalary struct {
	Month string // 年月(yyyyMM)
}
