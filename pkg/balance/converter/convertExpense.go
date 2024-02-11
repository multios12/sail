package converter

import (
	"encoding/json"
	"errors"
	"fmt"
	"path"
	"regexp"
)

// 指定されたPDFファイルから、経費明細データ・画像を作成する
func ConvertExpenseDetail(filename string, pdfPassword string) ([]byte, []byte, error) {
	r := regexp.MustCompile(`(\d+)年(\d+)月(経費)_.+\.pdf`)
	if !r.MatchString(filename) {
		return nil, nil, errors.New("給与明細・経費明細のPDFファイルを選択して下さい")
	}

	month := path.Base(filename)
	month = r.ReplaceAllString(month, "$1$2")
	pages, err := pdfinfo(filename, pdfPassword)
	if err != nil {
		return nil, nil, fmt.Errorf("給与明細PDFファイルが読み込めません。パスワードを確認してください。: %w", err)
	}

	if data, image, err := readExpensePdf(filename, month, pages, pdfPassword); err != nil {
		return nil, nil, err
	} else if e, err := readTextFileToExpenseItem(data["expense01"]); err != nil {
		return nil, nil, err
	} else {
		e.Month = month
		a, err := json.Marshal(e)
		return a, image, err
	}
}

// ----------------------------------------------------------------------------
// 経費等支給明細書データ作成
func readExpensePdf(src string, month string, pages string, pdfPassword string) (map[string]string, []byte, error) {
	// 画像
	image, err := readImage(src, pdfPassword)
	if err != nil {
		return nil, nil, err
	}

	var data map[string]string = map[string]string{}
	data["expense01"] = pdftostring(src, "-x 300 -y 140 -W 300 -H 40", pdfPassword)
	data["expense01"] += pdftostring(src, "-x 100 -y 210 -W 800 -H 40", pdfPassword)
	data["expense01"] += pdftostring(src, "-x 100 -y 250 -W 800 -H 40", pdfPassword)
	return data, image, nil
}
