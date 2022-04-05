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

	//var texts []string
	for _, file := range files {
		ext := path.Ext(file.Name())
		if file.IsDir() || ext != ".pdf" {
			continue
		}

		create(file)
	}
}

func create(file fs.FileInfo) {
	r := regexp.MustCompile(`(\d+)年(\d+)月給与_.+`)
	if !r.MatchString(file.Name()) {
		return
	}

	month := r.ReplaceAllString(file.Name(), "$1$2")

	if len(month) == 5 {
		month = month[:4] + "0" + month[4:]
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

	// タイトル
	dist = filepath.Join(monthPath, "salary-title.txt")
	exec.Command("pdftotext", src, dist, "-opw", password, "-x", "500", "-y", "30", "-W", "400", "-H", "20").Output()

	// 出勤
	dist = filepath.Join(monthPath, "salary-count1.txt")
	exec.Command("pdftotext", src, dist, "-opw", password, "-x", "70", "-y", "100", "-W", "30", "-H", "60").Output()
	// 休出
	dist = filepath.Join(monthPath, "salary-count2.txt")
	exec.Command("pdftotext", src, dist, "-opw", password, "-x", "100", "-y", "100", "-W", "30", "-H", "60").Output()
	// 特休
	dist = filepath.Join(monthPath, "salary-count3.txt")
	exec.Command("pdftotext", src, dist, "-opw", password, "-x", "130", "-y", "100", "-W", "30", "-H", "60").Output()
	// 有休
	dist = filepath.Join(monthPath, "salary-count4.txt")
	exec.Command("pdftotext", src, dist, "-opw", password, "-x", "160", "-y", "100", "-W", "20", "-H", "60").Output()
	// 欠勤
	dist = filepath.Join(monthPath, "salary-count5.txt")
	exec.Command("pdftotext", src, dist, "-opw", password, "-x", "190", "-y", "100", "-W", "30", "-H", "60").Output()

	// 時間
	dist = filepath.Join(monthPath, "salary-time1.txt")
	exec.Command("pdftotext", src, dist, "-opw", password, "-x", "220", "-y", "90", "-W", "190", "-H", "60").Output()
	dist = filepath.Join(monthPath, "salary-time2.txt")
	exec.Command("pdftotext", src, dist, "-opw", password, "-x", "410", "-y", "90", "-W", "300", "-H", "60").Output()

	// 支給1
	dist = filepath.Join(monthPath, "salary1.txt")
	exec.Command("pdftotext", src, dist, "-opw", password, "-x", "75", "-y", "150", "-W", "800", "-H", "60").Output()
	// 支給2
	dist = filepath.Join(monthPath, "salary2.txt")
	exec.Command("pdftotext", src, dist, "-opw", password, "-x", "75", "-y", "200", "-W", "800", "-H", "60").Output()
	// 控除１
	dist = filepath.Join(monthPath, "salary-cost1.txt")
	exec.Command("pdftotext", src, dist, "-opw", password, "-x", "75", "-y", "330", "-W", "800", "-H", "60").Output()
	// 控除２
	dist = filepath.Join(monthPath, "salary-cost2.txt")
	exec.Command("pdftotext", src, dist, "-opw", password, "-x", "75", "-y", "380", "-W", "800", "-H", "60").Output()

	// 合計
	dist = filepath.Join(monthPath, "salary-total.txt")
	exec.Command("pdftotext", src, dist, "-opw", password, "-x", "75", "-y", "440", "-W", "800", "-H", "60").Output()
}
