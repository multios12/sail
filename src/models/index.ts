// 給与支給明細書モデル
export type DetailModel = {
    Month: string;       // 年月
    Title: string;  // タイトル
    IsError: boolean;         // 取得エラー
    Counts: DetailItem[]; // 日数配列
    Times: TimeItem[]   // 時間配列
    Salarys: DetailItem[] // 支給配列
    Costs: DetailItem[] // 控除配列
    Totals: DetailItem[] // 合計配列
    Expense: number            // 経費等支給合計額
    Expenses: ExpenseItem[] // 経費内訳配列
    Images: string[]      // 画像ファイル配列
}

export type DetailItem = {
    Name: string;
    Value: number;
}

export type TimeItem = {
    Name: string;
    Value: string;
}

export type ExpenseItem = {
    Name: string // 項目名
    Amount: number    // 金額
    Memo: string // 備考
}

export type YearModel = {
    Year: string        // 年
    EnableYears: string[]    // 利用可能な年のリスト
    Totals: DetailItem[]  // 合計リスト
    Details: DetailModel[] // 月ごとの給与支給明細リスト
}