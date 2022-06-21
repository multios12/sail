package main

import (
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
)

// ----------------------------------------------------------------------------
// 経費等支給明細書データ作成
func createExpenseData(file fs.FileInfo, src string, pages string) error {

	r := regexp.MustCompile(`(\d+)年(\d+)月(経費)_.+`)

	month := r.ReplaceAllString(file.Name(), "$1$2")
	if len(month) == 5 {
		month = month[:4] + "0" + month[4:]
	}

	monthPath := filepath.Join(dataPath, month)
	dist := filepath.Join(monthPath, "expense01.txt")
	if _, err := os.Stat(dist); !os.IsNotExist(err) {
		return err
	}
	os.Mkdir(monthPath, os.ModePerm)

	// 画像
	dist = filepath.Join(monthPath, "expense")
	exec.Command("pdftocairo", src, dist, "-opw", password, "-png").Output()

	pdftotext(src, filepath.Join(monthPath, "expense01.txt"), "-x 300 -y 140 -W 300 -H 40")
	pdftotext(src, filepath.Join(monthPath, "expense01.txt"), "-x 100 -y 210 -W 800 -H 40")
	pdftotext(src, filepath.Join(monthPath, "expense01.txt"), "-x 100 -y 250 -W 800 -H 40")
	return nil
}
