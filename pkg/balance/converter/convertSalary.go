package converter

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"image/png"
	"os/exec"
	"path"
	"regexp"
	"strconv"
	"strings"

	"github.com/harukasan/go-libwebp/webp"
)

// 指定されたPDFファイルから、給与・賞与明細データ・画像を作成する
func ConvertSalaryDetail(filename string, pdfPassword string) ([]byte, []byte, error) {
	r := regexp.MustCompile(`(\d+)年(\d+)月(給与|.*賞与)_.+\.pdf`)
	if !r.MatchString(filename) {
		return nil, nil, errors.New("給与明細・経費明細のPDFファイルを選択して下さい")
	}

	s := SalaryDetail{}
	s.Month = path.Base(filename)
	s.Month = r.ReplaceAllString(s.Month, "$1$2")
	if len(s.Month) == 5 {
		s.Month = s.Month[:4] + "0" + s.Month[4:]
	}

	sv := r.ReplaceAllString(filename, "$3")
	if strings.Contains(sv, "賞与") {
		s.Month += "S"
	}

	pages, err := pdfinfo(filename, pdfPassword)
	if err != nil {
		return nil, nil, fmt.Errorf("給与明細PDFファイルが読み込めません。パスワードを確認してください。: %w", err)
	}

	if data, image, err := readSalaryPdf(filename, s.Month, pages, pdfPassword); err != nil {
		return nil, nil, err
	} else {
		s.Counts, s.IsError = readTextFileToDetailItem(data["salary01.txt"])
		s.Times, s.IsError = readTextFileToTimeItem(data["salary02.txt"])
		s.Salarys, s.IsError = readTextFileToDetailItem(data["salary10.txt"])
		s.Costs, s.IsError = readTextFileToDetailItem(data["salary20.txt"])
		s.Totals, s.IsError = readTextFileToDetailItem(data["salary30.txt"])
		s.Title = fmt.Sprintf("%s年%s月 ", s.Month[:4], s.Month[4:6])
		if strings.Contains(s.Month, "S") {
			s.Title += "賞与"
		} else {
			s.Title += "給与"
		}

		a, b := json.Marshal(s)
		return a, image, b

	}
}

// 給与支給明細書データの作成
func readSalaryPdf(filename string, month string, pages string, pdfPassword string) (map[string]string, []byte, error) {
	n, _ := strconv.Atoi(month[:6])

	// 画像
	imageText, err := readImage(filename, pdfPassword)
	if err != nil {
		return nil, nil, err
	}

	data := map[string]string{}
	if n >= 202005 && strings.Contains(month, "S") {
		data = readSalaryFrom202005S(filename, data, pdfPassword)
	} else if n >= 202005 {
		data = readSalaryFrom202005(filename, data, pages, pdfPassword)
	} else if n >= 202003 {
		data = readSalaryFrom202003(filename, data, pdfPassword)
	} else {
		data = readSalaryFrom201901(filename, data, pdfPassword)
	}
	return data, imageText, nil
}

// 給与支給明細書(2019年01月～)のテキストデータ読み込み
func readSalaryFrom201901(src string, data map[string]string, pdfPassword string) map[string]string {
	// 出勤/休出/特休/欠勤
	data["salary01.txt"] = pdftostring(src, "-x 70 -y 100 -W 20 -H 60", pdfPassword)
	data["salary01.txt"] += pdftostring(src, "-x 100 -y 100 -W 20 -H 60", pdfPassword)
	data["salary01.txt"] += pdftostring(src, "-x 130 -y 100 -W 20 -H 60", pdfPassword)
	data["salary01.txt"] += pdftostring(src, "-x 180 -y 100 -W 20 -H 60", pdfPassword)

	// 有休
	data["salary04.txt"] = pdftostring(src, "-x 150 -y 100 -W 20 -H 60", pdfPassword)

	// 有休残
	data["salary05.txt"] = pdftostring(src, "-x 210 -y 100 -W 20 -H 60", pdfPassword)

	// 時間
	data["salary02.txt"] = pdftostring(src, "-x 260 -y 110 -W 190 -H 60", pdfPassword)
	data["salary02.txt"] += pdftostring(src, "-x 450 -y 110 -W 500 -H 60", pdfPassword)
	// 支給
	data["salary10.txt"] = pdftostring(src, "-x 85 -y 170 -W 640 -H 80", pdfPassword)
	data["salary10.txt"] += pdftostring(src, "-x 85 -y 220 -W 640 -H 60", pdfPassword)
	// 控除
	data["salary20.txt"] = pdftostring(src, "-x 75 -y 280 -W 800 -H 100", pdfPassword)
	// 合計
	data["salary30.txt"] = pdftostring(src, "-x 250 -y 400 -W 800 -H 80", pdfPassword)
	return data
}

// 給与支給明細書(2020年03月～)のテキストデータ読み込み
func readSalaryFrom202003(src string, data map[string]string, pdfPassword string) map[string]string {
	// 出勤/休出/特休/欠勤
	data["salary01.txt"] = pdftostring(src, "-x 70 -y 100 -W 30 -H 60", pdfPassword)
	data["salary01.txt"] += pdftostring(src, "-x 100 -y 100 -W 30 -H 60", pdfPassword)
	data["salary01.txt"] += pdftostring(src, "-x 130 -y 100 -W 30 -H 60", pdfPassword)
	data["salary01.txt"] += pdftostring(src, "-x 190 -y 100 -W 30 -H 60", pdfPassword)
	// 有休
	data["salary04.txt"] = pdftostring(src, "-x 160 -y 100 -W 20 -H 60", pdfPassword)

	// 時間
	data["salary02.txt"] = pdftostring(src, "-x 220 -y 100 -W 190 -H 60", pdfPassword)
	data["salary02.txt"] += pdftostring(src, "-x 410 -y 100 -W 300 -H 60", pdfPassword)
	// 支給
	data["salary10.txt"] = pdftostring(src, "-x 75 -y 165 -W 800 -H 60", pdfPassword)
	data["salary10.txt"] += pdftostring(src, "-x 75 -y 220 -W 800 -H 60", pdfPassword)
	// 控除
	data["salary20.txt"] = pdftostring(src, "-x 75 -y 350 -W 800 -H 40", pdfPassword)
	data["salary20.txt"] += pdftostring(src, "-x 75 -y 390 -W 800 -H 40", pdfPassword)
	// 合計
	data["salary30.txt"] = pdftostring(src, "-x 75 -y 440 -W 800 -H 60", pdfPassword)
	return data
}

// 給与支給明細書(2020年05月～)のテキストデータ読み込み
func readSalaryFrom202005(src string, data map[string]string, pages string, pdfPassword string) map[string]string {
	// 出勤/休出/特休/欠勤
	data["salary01.txt"] = pdftostring(src, "-x 70 -y 100 -W 30 -H 60 -f "+pages+" -l "+pages, pdfPassword)
	data["salary01.txt"] += pdftostring(src, "-x 100 -y 100 -W 30 -H 60 -f "+pages+" -l "+pages, pdfPassword)
	data["salary01.txt"] += pdftostring(src, "-x 130 -y 100 -W 30 -H 60 -f "+pages+" -l "+pages, pdfPassword)
	data["salary01.txt"] += pdftostring(src, "-x 190 -y 100 -W 30 -H 60 -f "+pages+" -l "+pages, pdfPassword)

	// 有休
	data["salary04.txt"] = pdftostring(src, "-x 160 -y 100 -W 20 -H 60 -f "+pages+" -l "+pages, pdfPassword)

	// 時間
	data["salary02.txt"] = pdftostring(src, "-x 220 -y 100 -W 190 -H 60 -f "+pages+" -l "+pages, pdfPassword)
	data["salary02.txt"] += pdftostring(src, "-x 410 -y 100 -W 300 -H 60 -f "+pages+" -l "+pages, pdfPassword)

	// 支給
	data["salary10.txt"] = pdftostring(src, "-x 75 -y 160 -W 800 -H 60 -f "+pages+" -l "+pages, pdfPassword)
	data["salary10.txt"] += pdftostring(src, "-x 75 -y 200 -W 800 -H 60 -f "+pages+" -l "+pages, pdfPassword)

	// 控除
	data["salary20.txt"] = pdftostring(src, "-x 75 -y 330 -W 800 -H 60 -f "+pages+" -l "+pages, pdfPassword)
	data["salary20.txt"] += pdftostring(src, "-x 75 -y 380 -W 800 -H 60 -f "+pages+" -l "+pages, pdfPassword)

	// 合計
	data["salary30.txt"] = pdftostring(src, "-x 75 -y 440 -W 800 -H 60 -f "+pages+" -l "+pages, pdfPassword)
	return data
}

// 給与支給明細書[賞与](2020年05月～)のテキストデータ読み込み
func readSalaryFrom202005S(src string, data map[string]string, pdfPassword string) map[string]string {
	data["salary10.txt"] = pdftostring(src, "-x 75 -y 160 -W 800 -H 60", pdfPassword)
	data["salary20.txt"] = pdftostring(src, "-x 75 -y 280 -W 800 -H 60", pdfPassword)
	data["salary30.txt"] = pdftostring(src, "-x 75 -y 400 -W 800 -H 80", pdfPassword)
	return data
}

func readImage(filename string, pdfPassword string) ([]byte, error) {
	var webpBuffer bytes.Buffer
	b, err := exec.Command("pdftocairo", filename, "-", "-opw", pdfPassword, "-singlefile", "-png").Output()
	if err != nil {
		return nil, err
	} else if pngImage, err := png.Decode(bytes.NewReader(b)); err != nil {
		return nil, err
	} else if con, err := webp.ConfigPreset(webp.PresetDefault, 45); err != nil {
		return nil, err
	} else if err = webp.EncodeRGBA(&webpBuffer, pngImage, con); err != nil {
		return nil, err
	}

	return webpBuffer.Bytes(), nil
}
