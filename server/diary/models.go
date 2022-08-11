package diary

import (
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

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

func readDetailFile(l lineModel) *detailModel {
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
func (d *detailModel) writeDetailFile() error {
	month := d.Day[0:4] + d.Day[5:7]

	// monthファイルの更新
	m := readListFile(month)
	target := lineModel{d.Day, d.Outline, d.Tags, len(d.Detail) > 0, 0}
	m.updateLine(month, target)

	// 詳細ファイルの更新
	filename := strings.ReplaceAll(d.Day, "-", "") + ".txt"
	filename = path.Join(diaryPath, filename)
	if len(d.Detail) == 0 {
		if _, err := os.Stat(filename); err == nil {
			return os.Remove(filename)
		}
		return nil
	}

	data := d.Day + "\n" + d.Outline + "\n" + strings.Join(d.Tags, "#") + "\n" + d.Detail
	return ioutil.WriteFile(filename, []byte(data), fs.ModePerm)
}
