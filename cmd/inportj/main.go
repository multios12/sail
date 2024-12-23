package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/multios12/sail/pkg/balance/models"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

var dataPath string // データディレクトリ

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "給与明細からデータを抽出し一覧表示するWebサーバ\n\n")
		fmt.Fprintf(flag.CommandLine.Output(), "sail [Source] [options...]\n\n")
		fmt.Fprintf(flag.CommandLine.Output(), `Source  PDFなどのデータが格納されているディレクトリを指定します (default "./data")`+"\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Options\n")
		flag.PrintDefaults()
		os.Exit(0)
	}
	flag.StringVar(&dataPath, "d", "./data", "")
	flag.Parse()

	dataPath, _ = filepath.Abs(dataPath)
}

func main() {
	jPath := path.Join(dataPath, "j")
	files, _ := os.ReadDir(jPath)

	for _, j := range files {
		if path.Ext(j.Name()) != ".csv" {
			continue
		}
		file, err := os.Open(j.Name())
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		month := j.Name()[:6]

		useDetails := []models.UseDetail{}

		// ShiftJISのデコーダーを噛ませたReaderを作成
		reader := bufio.NewReader(transform.NewReader(file, japanese.ShiftJIS.NewDecoder()))
		for {
			line, _, err := reader.ReadLine()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}

			model := models.CreateUseDetail("J", month, string(line))
			useDetails = append(useDetails, model)
		}
	}
}

func validateArgs() error {
	if i, err := os.Stat(dataPath); err != nil || !i.IsDir() {
		err = os.Mkdir(dataPath, os.ModePerm)
		if err != nil {
			message := fmt.Sprintf("データディレクトリが作成できません。(%s)", err)
			return errors.New(message)
		}
	}

	return nil
}
