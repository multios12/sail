package main

import (
	"embed"
	"flag"
	"os"
	"path/filepath"
)

// 明細リスト
var details []DetailModel

// データディレクトリパス
var dataPath string

// //go:embed static/*
var static embed.FS

func main() {
	flag.StringVar(&dataPath, "path", "./data", "data directory")
	port := flag.String("port", ":3000", "server port")
	flag.Parse()

	dataPath, _ = filepath.Abs(dataPath)
	if _, err := os.Stat(dataPath); err != nil {
		os.Mkdir(dataPath, 0777)
	}

	details = readAllData(dataPath)

	router := routerInitial(static)
	router.Run(*port)
}
