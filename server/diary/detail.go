package diary

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type detail struct {
	Day  string   // 日付
	Tags []string // タグ
	Memo string   // 内容
}

func readDetail(day string) detail {
	p := filepath.Join(diaryPath, day+".txt")
	d := detail{Day: day, Tags: []string{}, Memo: ""}
	if _, err := os.Stat(p); err != nil {
		return d
	}
	b, err := ioutil.ReadFile(p)
	if err != nil {
		return d
	}
	value := string(b)
	i := strings.Index(value, "\r\n")
	if i == -1 {
		return d
	}
	value = value[i:]
	i = strings.Index(value, "\r\n")
	if i == -1 {
		return d
	}
	d.Tags = strings.Split(value[:i], ",")
	d.Memo = value[i:]

	return d
}
func writeDetail(d detail) error {
	p := filepath.Join(diaryPath, d.Day+".txt")
	tags := strings.Join(d.Tags, ",")
	value := d.Day + "\r\n" + tags + "\r\n" + d.Memo
	return ioutil.WriteFile(p, []byte(value), os.ModePerm)
}
