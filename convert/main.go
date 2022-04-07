package main

import (
	"flag"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var password string
var dataPath string

// コンテキスト
func main() {
	flag.StringVar(&password, "password", "", "PDF password")
	flag.StringVar(&dataPath, "path", "./data", "data directory")
	flag.Parse()

	files, err := ioutil.ReadDir(dataPath)
	if err != nil {
		log.Print(err)
		return
	}

	for _, file := range files {
		createSalaryData(file)
	}
	for _, file := range files {
		createExpenseData(file)
	}
}

// 給与支給明細書データの作成
func createSalaryData(file fs.FileInfo) {
	ext := path.Ext(file.Name())
	if file.IsDir() || ext != ".pdf" {
		return
	}
	r := regexp.MustCompile(`(\d+)年(\d+)月(給与|.*賞与)_.+`)
	if !r.MatchString(file.Name()) {
		return
	}

	month := r.ReplaceAllString(file.Name(), "$1$2")
	if len(month) == 5 {
		month = month[:4] + "0" + month[4:]
	}
	n, _ := strconv.Atoi(month)

	s := r.ReplaceAllString(file.Name(), "$3")
	if strings.Contains(s, "賞与") {
		month += "S"
	}

	monthPath := filepath.Join(dataPath, month)
	if _, err := os.Stat(monthPath); !os.IsNotExist(err) {
		return
	}
	os.Mkdir(monthPath, os.ModePerm)

	src := filepath.Join(dataPath, file.Name())
	println(src)

	// 画像
	dist := filepath.Join(monthPath, "salary")
	exec.Command("pdftocairo", src, dist, "-opw", password, "-png").Output()

	pages := pdfinfo(src)

	if n >= 202005 && strings.Contains(s, "賞与") {
		readSalaryFrom202005S(src, monthPath)
	} else if n >= 202005 {
		readSalaryFrom202005(src, monthPath, pages)
	} else if n >= 202003 {
		readSalaryFrom202003(src, monthPath)
	} else {
		readSalaryFrom201901(src, monthPath)
	}
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
	pdftotext(src, filepath.Join(monthPath, "salary10.txt"), "-x 85 -y 170 -W 650 -H 80")
	pdftotext(src, filepath.Join(monthPath, "salary10.txt"), "-x 85 -y 240 -W 650 -H 60")
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

// ----------------------------------------------------------------------------
// 経費等支給明細書データ作成
func createExpenseData(file fs.FileInfo) {
	ext := path.Ext(file.Name())
	if file.IsDir() || ext != ".pdf" {
		return
	}
	r := regexp.MustCompile(`(\d+)年(\d+)月(経費)_.+`)
	if !r.MatchString(file.Name()) {
		return
	}

	month := r.ReplaceAllString(file.Name(), "$1$2")
	if len(month) == 5 {
		month = month[:4] + "0" + month[4:]
	}

	monthPath := filepath.Join(dataPath, month)
	dist := filepath.Join(monthPath, "expense01.txt")
	if _, err := os.Stat(dist); !os.IsNotExist(err) {
		return
	}
	os.Mkdir(monthPath, os.ModePerm)

	src := filepath.Join(dataPath, file.Name())
	println(src)

	// 画像
	dist = filepath.Join(monthPath, "expense")
	exec.Command("pdftocairo", src, dist, "-opw", password, "-png").Output()

	pdftotext(src, filepath.Join(monthPath, "expense01.txt"), "-x 300 -y 140 -W 300 -H 40")
	pdftotext(src, filepath.Join(monthPath, "expense01.txt"), "-x 100 -y 210 -W 800 -H 40")
	pdftotext(src, filepath.Join(monthPath, "expense01.txt"), "-x 100 -y 250 -W 800 -H 40")
}

// ----------------------------------------------------------------------------

// pdftotextコマンドの実行し、テキストデータを出力する
func pdftotext(src string, dist string, opt string) {
	opt = src + " " + dist + " -opw " + password + " " + opt
	args := strings.Split(opt, " ")

	text := ""
	if _, e := os.Stat(dist); e == nil {
		text = readTextFileToString(dist)
	}

	exec.Command("pdftotext", args...).Output()
	text += readTextFileToString(dist)
	ioutil.WriteFile(dist, []byte(text), fs.ModePerm)
}

// pdfinfoコマンドを実行し、ページ数を返す
func pdfinfo(filename string) string {
	b, err := exec.Command("pdfinfo", filename, "-opw", password).Output()
	if err != nil {
		return "1"
	}

	text := string(b)
	lines := strings.Split(text, "\n")
	for _, v := range lines {
		if strings.Contains(v, "Pages:") {
			v = strings.ReplaceAll(v, "Pages:", "")
			v = strings.TrimSpace(v)
			return v
		}
	}
	return "1"
}

// テキストファイルの読み込み、文字列を返す
func readTextFileToString(filename string) string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}

	text := string(bytes)
	text = strings.ReplaceAll(text, "\f", "")

	return text
}
