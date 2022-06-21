package main

import (
	"embed"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var salaries []Salary  // 明細リスト
var isConvertMode bool // クロールモード
var dataPath string    // データディレクトリ
var password string    // PDFパスワード
var port string        // 起動ポート
//go:embed static/*
var static embed.FS

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "給与明細からデータを抽出し一覧表示するWebサーバ\n\n")
		fmt.Fprintf(flag.CommandLine.Output(), "sail [Source] [options...]\n\n")
		fmt.Fprintf(flag.CommandLine.Output(), `Source  PDFなどのデータが格納されているディレクトリを指定します (default "./data")`+"\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Options\n")
		flag.PrintDefaults()
		os.Exit(0)
	}
	flag.BoolVar(&isConvertMode, "c", false, "給与明細PDFからのデータ抽出のみを行い、サーバは起動しません")
	flag.StringVar(&password, "w", "", "PDF処理時に指定されたパスワードを使用してPDFを開きます")
	flag.StringVar(&port, "p", ":3000", "Webサーバが使用するポートを指定します")
	flag.StringVar(&dataPath, "d", "./data", "")
	flag.Parse()

	dataPath, _ = filepath.Abs(dataPath)
}

func main() {
	// クロールモード
	if isConvertMode {
		err := convert()
		if err != nil {
			fmt.Println(err)
		}
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
	router.Run(port)
}

func validateArgs() error {
	if i, err := os.Stat(dataPath); err != nil || !i.IsDir() {
		err = os.Mkdir(dataPath, os.ModePerm)
		if err != nil {
			return errors.New(fmt.Sprintf("データディレクトリが作成できません。(%s)", err))
		}
	}

	if password == "" {
		println("[warning]パスワードが設定されていないため、明細の取り込みに失敗する可能性があります。")
	}

	return nil
}
