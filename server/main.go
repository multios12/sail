package main

import (
	"embed"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/multios12/sail/balance"
	"github.com/multios12/sail/diary"
	"github.com/multios12/sail/memo"
)

var dataPath string // データディレクトリ
var password string // PDFパスワード
var port string     // 起動ポート
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
	flag.StringVar(&password, "w", "", "PDF処理時に指定されたパスワードを使用してPDFを開きます")
	flag.StringVar(&port, "p", ":3000", "Webサーバが使用するポートを指定します")
	flag.StringVar(&dataPath, "d", "./data", "")
	flag.Parse()

	dataPath, _ = filepath.Abs(dataPath)
}

func main() {
	router := gin.Default()
	router.GET("/", getStatic)
	router.GET("/index.html", getStatic)
	router.GET("/favicon.ico", getStatic)
	balance.Initial(router, dataPath, password)
	memo.Initial(router, dataPath)
	diary.Initial(router, dataPath)

	router.Run(port)
}

// スタティックリソース GET API
func getStatic(c *gin.Context) {
	p := "static" + c.Request.URL.Path
	c.FileFromFS(p, http.FS(static))
}

func validateArgs() error {
	if i, err := os.Stat(dataPath); err != nil || !i.IsDir() {
		err = os.Mkdir(dataPath, os.ModePerm)
		if err != nil {
			message := fmt.Sprintf("データディレクトリが作成できません。(%s)", err)
			return errors.New(message)
		}
	}

	if password == "" {
		log.Println("warning: パスワードが設定されていないため、明細の取り込みに失敗する可能性があります。")
	}

	return nil
}
