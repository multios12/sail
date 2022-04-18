package main

import (
	"embed"
	"flag"
	"os"
	"path/filepath"
	"strings"
)

// 明細リスト
var salaries []Salary

// クロールモード
var isConvertMode bool

// データディレクトリパス
var dataPath string

// PDFパスワード
var password string

//go:embed static/*
var static embed.FS

func main() {
	flag.BoolVar(&isConvertMode, "convert", false, "true is PDF convert mode")
	flag.StringVar(&password, "password", "", "PDF password")
	flag.StringVar(&dataPath, "datapath", "./data", "data directory")
	port := flag.String("port", ":3000", "server port")
	flag.VisitAll(func(f *flag.Flag) {
		if s := os.Getenv(strings.ToUpper(f.Name)); s != "" {
			f.Value.Set(s)
		}
	})
	flag.Parse()

	dataPath, _ = filepath.Abs(dataPath)
	if _, err := os.Stat(dataPath); err != nil {
		os.Mkdir(dataPath, 0777)
	}

	if password == "" {
		println("[warning]パスワードが設定されていないため、明細の取り込みに失敗する可能性があります。")
	}

	// クロールモード
	if isConvertMode {
		convert()
		return
	}

	dbOpen()

	var err error
	salaries, err = readAllData(dataPath)
	if err != nil {
		println(err)
		return
	}

	router := routerInitial(static)
	router.Run(*port)
}
