package main

import (
	"embed"
	"flag"
	"os"
	"path/filepath"
)

// 明細リスト
var details []DetailModel

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
	flag.StringVar(&dataPath, "path", "./data", "data directory")
	port := flag.String("port", ":3000", "server port")
	flag.Parse()

	dataPath, _ = filepath.Abs(dataPath)
	if _, err := os.Stat(dataPath); err != nil {
		os.Mkdir(dataPath, 0777)
	}

	// クロールモード
	if isConvertMode {
		convert()
		return
	}

	details = readAllData(dataPath)

	router := routerInitial(static)
	router.Run(*port)
}
