package balance

import (
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/multios12/sail/pkg/balance/converter"
	"github.com/multios12/sail/pkg/balance/models"
)

// 指定されたディレクトリからPDFファイルを読み込み、給与明細モデルリストを返す
func readSalaries(balancePath string, pdfPassword string) error {
	// PDFファイルの取得
	pdfPath := path.Join(balancePath, "_pdf")
	files, err := os.ReadDir(pdfPath)
	if err != nil {
		return err
	}
	// salaryテーブルに存在しない給与明細データを探して登録する
	for _, file := range files {
		filename := path.Join(pdfPath, file.Name())
		readSalary(filename, pdfPassword)
	}
	return err
}

func readSalary(filename string, pdfPassword string) {
	r := regexp.MustCompile(`(\d+)年(\d+)月(給与|.*賞与|経費)_.+\.pdf`)
	if r.MatchString(filename) {
		month := filepath.Base(filename)
		month = r.ReplaceAllString(month, "$1$2")
		if len(month) == 5 {
			month = month[:4] + "0" + month[4:]
		}

		sv := r.ReplaceAllString(filename, "$3")
		s, err := findBalanceByMonth(month)
		s.Month = month
		if err != nil {
			s = Balance{Month: month}
		}

		var detailType models.BalanceType
		if strings.Contains(filename, "経費") {
			detailType = models.DetailTypeExpense
		} else if strings.Contains(sv, "賞与") {
			detailType = models.DetailTypeBonus
		} else {
			detailType = models.DetailTypeSalary
		}

		d, _ := findBalanceDetailByMonthType(month, detailType)
		if len(d.Json) > 0 {
			return
		}

		if detailType == models.DetailTypeExpense {
			json, image, _ := converter.ConvertExpenseDetail(filename, pdfPassword)
			upsertBalanceDetail(month, detailType, json, image)
		} else {
			json, image, _ := converter.ConvertSalaryDetail(filename, pdfPassword)
			upsertBalanceDetail(month, detailType, json, image)
		}
		if err != nil {
			return
		}
	}
}
