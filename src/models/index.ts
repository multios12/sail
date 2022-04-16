/** 給与支給明細書（集計）モデル */
export type SalaryModel = {
  /** 年 */
  Year: string
  /** 利用可能な年のリスト */
  EnableYears: string[]
  /** 合計配列 */
  Totals: DetailItem[]
  /** 月ごとの給与支給明細リスト */
  Details: SalaryMonthModel[]
}

/** 給与支給明細書モデル */
export type SalaryMonthModel = {
  /** 年月 */
  Month: string
  /** タイトル */
  Title?: string
  /** 取得エラー */
  IsError?: false
  /** 勤務日数等配列 */
  Counts: DetailItem[]
  /** 業務時間等配列 */
  Times: TimeItem[]
  /** 支給金額配列 */
  Salarys: DetailItem[]
  /** 控除金額配列 */
  Costs: DetailItem[]
  /** 合計金額配列 */
  Totals: DetailItem[]
  /** 経費等支給合計額 */
  Expense: number
  /** 経費内訳配列 */
  Expenses: ExpenseItem[]
  /** 画像ファイル配列 */
  Images: string[]
}

/** 詳細項目 */
export type DetailItem = {
  /** 項目名 */
  Name: string
  /** 金額 */
  Value: 0
}

/** 時間項目 */
export type TimeItem = {
  /** 項目名 */
  Name: string
  Value: string
}

/** 経費項目 */
export type ExpenseItem = {
  /** 項目名 */
  Name: string
  /** 金額 */
  Amount: number
  /** 備考 */
  Memo: string
}

// 支出集計モデル
export type SumCostModel= {
	Year       : string   // 年
	EnableYears: string[] // 利用可能な年のリスト
	Costs      : CostModel[]   // 月ごとの支出配列
}


/** 支出モデル */
export type CostModel = {
  Month: number      // 年月(yyyyMM)
  Water: number       // 水道費
  Electric: number       // 電気費
  Gas: number       // ガス費
  Mobile: number       // 通信費(携帯)
  Line: number       // 通信費(固定)
}