package balance

import (
	"errors"
	"io/fs"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

// コンテキスト
func Convert(dataPath string) error {
	pdfPath := path.Join(dataPath, "_pdf")
	files, err := ioutil.ReadDir(pdfPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		createData(dataPath, file)
	}

	return nil
}

func createData(dataPath string, file fs.FileInfo) error {
	ext := path.Ext(file.Name())
	if file.IsDir() || ext != ".pdf" {
		return errors.New("給与明細・経費明細のPDFファイルを選択して下さい")
	}

	regexp := regexp.MustCompile(`(\d+)年(\d+)月(給与|.*賞与|経費)_.+`)
	if regexp.MatchString(file.Name()) {
		src := filepath.Join(dataPath, "_pdf", file.Name())
		println(src)
		pages, err := pdfinfo(src)
		if err != nil {
			return errors.New("給与明細PDFファイルが読み込めません。パスワードを確認してください。")
		}

		s := regexp.ReplaceAllString(file.Name(), "$3")
		if !strings.Contains(s, "経費") {
			return createSalaryData(dataPath, file, src, pages)
		} else {
			return createExpenseData(dataPath, file, src, pages)
		}
	}
	return errors.New("給与明細・経費明細のPDFファイルを選択して下さい")
}

// ----------------------------------------------------------------------------

// pdftotextコマンドを実行し、テキストデータを出力する
func pdftotext(src string, dist string, opt string) {
	opt = src + " " + dist + " -opw " + pdfPassword + " " + opt
	args := strings.Split(opt, " ")

	text := ""
	if _, e := os.Stat(dist); e == nil {
		text = readTextFileToString(dist)
	}

	command := "pdftotext"

	exec.Command(command, args...).Output()
	text += readTextFileToString(dist)
	ioutil.WriteFile(dist, []byte(text), fs.ModePerm)
}

// pdfinfoコマンドを実行し、ページ数を返す
func pdfinfo(filename string) (string, error) {
	command := "pdfinfo"

	b, err := exec.Command(command, filename, "-opw", pdfPassword).Output()
	if err != nil {
		return "1", err
	}

	text := string(b)
	lines := strings.Split(text, "\n")
	for _, v := range lines {
		if strings.Contains(v, "Pages:") {
			v = strings.ReplaceAll(v, "Pages:", "")
			v = strings.TrimSpace(v)
			return v, nil
		}
	}
	return "1", nil
}

// テキストファイルを読み込み、文字列を返す
func readTextFileToString(filename string) string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}

	text := string(bytes)
	text = strings.ReplaceAll(text, "\f", "")
	text = strings.ReplaceAll(text, "\r\n", "\n")

	return text
}
