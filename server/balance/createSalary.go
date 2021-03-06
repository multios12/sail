package balance

import (
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// 給与支給明細書データの作成
func createSalaryData(dataPath string, filename string, src string, pages string) error {

	r := regexp.MustCompile(`(\d+)年(\d+)月(給与|.*賞与)_.+`)
	month := path.Base(filename)
	month = r.ReplaceAllString(month, "$1$2")
	if len(month) == 5 {
		month = month[:4] + "0" + month[4:]
	}
	n, _ := strconv.Atoi(month)

	s := r.ReplaceAllString(filename, "$3")
	if strings.Contains(s, "賞与") {
		month += "S"
	}

	monthPath := filepath.Join(balancePath, month)
	if _, err := os.Stat(monthPath); !os.IsNotExist(err) {
		return err
	}
	os.Mkdir(monthPath, os.ModePerm)

	// 画像
	dist := filepath.Join(monthPath, "salary")
	exec.Command("pdftocairo", src, dist, "-opw", pdfPassword, "-png").Output()

	if n >= 202005 && strings.Contains(s, "賞与") {
		readSalaryFrom202005S(src, monthPath)
	} else if n >= 202005 {
		readSalaryFrom202005(src, monthPath, pages)
	} else if n >= 202003 {
		readSalaryFrom202003(src, monthPath)
	} else {
		readSalaryFrom201901(src, monthPath)
	}

	return nil
}

// 給与支給明細書(2019年01月～)のテキストデータ読み込み
func readSalaryFrom201901(src string, monthPath string) {
	// 出勤/休出/特休/欠勤
	pdftotext(src, filepath.Join(monthPath, "salary01.txt"), "-x 70 -y 100 -W 20 -H 60")
	pdftotext(src, filepath.Join(monthPath, "salary01.txt"), "-x 100 -y 100 -W 20 -H 60")
	pdftotext(src, filepath.Join(monthPath, "salary01.txt"), "-x 130 -y 100 -W 20 -H 60")
	pdftotext(src, filepath.Join(monthPath, "salary01.txt"), "-x 180 -y 100 -W 20 -H 60")

	// 有休
	pdftotext(src, filepath.Join(monthPath, "salary04.txt"), "-x 150 -y 100 -W 20 -H 60")

	// 有休残
	pdftotext(src, filepath.Join(monthPath, "salary05.txt"), "-x 210 -y 100 -W 20 -H 60")

	// 時間
	pdftotext(src, filepath.Join(monthPath, "salary02.txt"), "-x 260 -y 110 -W 190 -H 60")
	pdftotext(src, filepath.Join(monthPath, "salary02.txt"), "-x 450 -y 110 -W 500 -H 60")
	// 支給
	pdftotext(src, filepath.Join(monthPath, "salary10.txt"), "-x 85 -y 170 -W 640 -H 80")
	pdftotext(src, filepath.Join(monthPath, "salary10.txt"), "-x 85 -y 220 -W 640 -H 60")
	// 控除
	pdftotext(src, filepath.Join(monthPath, "salary20.txt"), "-x 75 -y 280 -W 800 -H 100")
	// 合計
	pdftotext(src, filepath.Join(monthPath, "salary30.txt"), "-x 250 -y 400 -W 800 -H 80")
}

// 給与支給明細書(2020年03月～)のテキストデータ読み込み
func readSalaryFrom202003(src string, monthPath string) {
	// 出勤/休出/特休/欠勤
	pdftotext(src, filepath.Join(monthPath, "salary01.txt"), "-x 70 -y 100 -W 30 -H 60")
	pdftotext(src, filepath.Join(monthPath, "salary01.txt"), "-x 100 -y 100 -W 30 -H 60")
	pdftotext(src, filepath.Join(monthPath, "salary01.txt"), "-x 130 -y 100 -W 30 -H 60")
	pdftotext(src, filepath.Join(monthPath, "salary01.txt"), "-x 190 -y 100 -W 30 -H 60")
	// 有休
	pdftotext(src, filepath.Join(monthPath, "salary04.txt"), "-x 160 -y 100 -W 20 -H 60")

	// 時間
	pdftotext(src, filepath.Join(monthPath, "salary02.txt"), "-x 220 -y 100 -W 190 -H 60")
	pdftotext(src, filepath.Join(monthPath, "salary02.txt"), "-x 410 -y 100 -W 300 -H 60")
	// 支給
	pdftotext(src, filepath.Join(monthPath, "salary10.txt"), "-x 75 -y 165 -W 800 -H 60")
	pdftotext(src, filepath.Join(monthPath, "salary10.txt"), "-x 75 -y 220 -W 800 -H 60")
	// 控除
	pdftotext(src, filepath.Join(monthPath, "salary20.txt"), "-x 75 -y 350 -W 800 -H 40")
	pdftotext(src, filepath.Join(monthPath, "salary20.txt"), "-x 75 -y 390 -W 800 -H 40")
	// 合計
	pdftotext(src, filepath.Join(monthPath, "salary30.txt"), "-x 75 -y 440 -W 800 -H 60")
}

// 給与支給明細書(2020年05月～)のテキストデータ読み込み
func readSalaryFrom202005(src string, monthPath string, pages string) {
	// 出勤/休出/特休/欠勤
	pdftotext(src, filepath.Join(monthPath, "salary01.txt"), "-x 70 -y 100 -W 30 -H 60 -f "+pages+" -l "+pages)
	pdftotext(src, filepath.Join(monthPath, "salary01.txt"), "-x 100 -y 100 -W 30 -H 60 -f "+pages+" -l "+pages)
	pdftotext(src, filepath.Join(monthPath, "salary01.txt"), "-x 130 -y 100 -W 30 -H 60 -f "+pages+" -l "+pages)
	pdftotext(src, filepath.Join(monthPath, "salary01.txt"), "-x 190 -y 100 -W 30 -H 60 -f "+pages+" -l "+pages)
	// 有休
	pdftotext(src, filepath.Join(monthPath, "salary04.txt"), "-x 160 -y 100 -W 20 -H 60 -f "+pages+" -l "+pages)

	// 時間
	pdftotext(src, filepath.Join(monthPath, "salary02.txt"), "-x 220 -y 100 -W 190 -H 60 -f "+pages+" -l "+pages)
	pdftotext(src, filepath.Join(monthPath, "salary02.txt"), "-x 410 -y 100 -W 300 -H 60 -f "+pages+" -l "+pages)
	// 支給
	pdftotext(src, filepath.Join(monthPath, "salary10.txt"), "-x 75 -y 160 -W 800 -H 60 -f "+pages+" -l "+pages)
	pdftotext(src, filepath.Join(monthPath, "salary10.txt"), "-x 75 -y 200 -W 800 -H 60 -f "+pages+" -l "+pages)
	// 控除
	pdftotext(src, filepath.Join(monthPath, "salary20.txt"), "-x 75 -y 330 -W 800 -H 60 -f "+pages+" -l "+pages)
	pdftotext(src, filepath.Join(monthPath, "salary20.txt"), "-x 75 -y 380 -W 800 -H 60 -f "+pages+" -l "+pages)
	// 合計
	pdftotext(src, filepath.Join(monthPath, "salary30.txt"), "-x 75 -y 440 -W 800 -H 60 -f "+pages+" -l "+pages)
}

// 給与支給明細書[賞与](2020年05月～)のテキストデータ読み込み
func readSalaryFrom202005S(src string, monthPath string) {
	pdftotext(src, filepath.Join(monthPath, "salary10.txt"), "-x 75 -y 160 -W 800 -H 60")
	pdftotext(src, filepath.Join(monthPath, "salary20.txt"), "-x 75 -y 280 -W 800 -H 60")
	pdftotext(src, filepath.Join(monthPath, "salary30.txt"), "-x 75 -y 380 -W 800 -H 60")
}
