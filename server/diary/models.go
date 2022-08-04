package diary

import (
	"bufio"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
)

type listModel struct {
	WritedMonths []string // 記載された年月(yyyy-mm形式)
	Lines        []lineModel
}

func newListModel(month string) *listModel {
	p := filepath.Join(diaryPath, month+".txt")

	l := new(listModel)
	if _, err := os.Stat(p); err != nil {
		return l
	}

	fp, err := os.Open(p)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	reader := bufio.NewReaderSize(fp, 5000)
	var lines []lineModel
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		s := string(line)
		d := newLineModel(s)
		l.Lines = append(lines, *d)
	}

	return l
}

func (l *listModel) writeMonthFile(month string) error {
	sort.Slice(l.Lines, func(i, j int) bool { return l.Lines[i].Day > l.Lines[j].Day })

	dataFile := ""
	for _, l := range l.Lines {
		if dataFile != "" {
			dataFile += "\n"
		}
		dataFile += l.toString()
	}
	filename := filepath.Join(diaryPath, month+".txt")

	if len(l.Lines) == 0 {
		return os.Remove(filename)
	}
	return ioutil.WriteFile(filename, []byte(dataFile), os.ModePerm)
}

// 月ファイルの行を表すモデル
type lineModel struct {
	Day      string   // 日付(yyyy-mm-dd形式)
	Outline  string   // 概要
	Tags     []string // タグ
	IsDetail bool     // 詳細
	HCount   int      // HMemo数
}

func newLineModel(s string) *lineModel {
	l := new(lineModel)
	l.Day = s[0:10]
	l.IsDetail = s[10:11] == "*"
	s = s[11:]

	tags := strings.Split(s, "#")
	if len(tags) > 1 {
		l.Outline = tags[len(tags)-1]
		l.Tags = tags[:len(tags)-1]
	} else {
		l.Outline = s
		l.Tags = []string{}
	}
	l.HCount = 0
	return l
}

func (l *lineModel) toString() string {
	outline := strings.ReplaceAll(l.Outline, "#", "")
	line := l.Day
	if l.IsDetail {
		line += "*"
	} else {
		line += " "
	}

	if len(l.Tags) > 0 {
		line += strings.Join(l.Tags, "#") + "#"
	}
	line += outline
	return line
}

// 月ファイルの行を表すモデル
type detailModel struct {
	Day     string   // 日付(yyyy-mm-dd形式)
	Outline string   // 概要
	Tags    []string // タグ
	Detail  string   // 詳細
}

func newDetailModel(l lineModel) *detailModel {
	d := new(detailModel)
	// 詳細ファイルの取得
	filename := strings.ReplaceAll(l.Day, "-", "") + ".txt"
	filename = path.Join(diaryPath, filename)
	d.Day = l.Day
	d.Outline = l.Outline
	d.Tags = l.Tags
	if _, err := os.Stat(filename); err == nil {
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			return d
		}
		body := string(b)
		lines := strings.Split(body, "\n")
		if len(lines) > 3 {
			body = strings.Join(lines[3:], "\n")
		}
		d.Detail = body
	}
	return d
}

// 詳細情報を書き込む
func (l *detailModel) WriteDetail() {
	// 詳細ファイルの更新
	filename := strings.ReplaceAll(l.Day, "-", "") + ".txt"
	filename = path.Join(diaryPath, filename)
	if len(l.Detail) == 0 {
		if _, err := os.Stat(filename); err == nil {
			os.Remove(filename)
		}
		return
	}

	data := l.Day + "\n" + l.Outline + "\n" + strings.Join(l.Tags, "#") + "\n" + l.Detail
	ioutil.WriteFile(filename, []byte(data), fs.ModePerm)
}
