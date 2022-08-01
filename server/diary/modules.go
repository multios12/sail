package diary

import (
	"bufio"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

var diaryPath string

type lineModel struct {
	Day        string   // 日付(yyyy-mm-dd形式)
	Outline    string   // 概要
	Tags       []string // タグ
	Detail     string   // 詳細
	HMemoCount int      // HMemo数
}

type listModel struct {
	WritedMonths []string // 記載された年月(yyyy-mm形式)
	Lines        []lineModel
}

func getWritedMonths() []string {
	l, err := ioutil.ReadDir(diaryPath)
	var months []string
	if err != nil {
		return months
	}

	r := regexp.MustCompile(`^\d\d\d\d\d\d\.txt$`)

	for _, f := range l {
		if !f.IsDir() && r.MatchString(f.Name()) {
			v := f.Name()[:4] + "-" + f.Name()[4:6]
			months = append(months, v)
		}
	}

	if len(months) > 0 {
		sort.Slice(months, func(i int, j int) bool { return months[i] > months[j] })
	} else {
		months = append(months, time.Now().Format("200601"))
	}

	return months
}

// ----------------------------------------------------------------------------
func readLine(day string) lineModel {
	lines := readMonthFile(day[0:4] + day[5:7])
	for _, l := range lines {
		if day == l.Day {
			// 詳細ファイルの取得
			filename := strings.ReplaceAll(day, "-", "") + ".txt"
			filename = path.Join(diaryPath, filename)
			if _, err := os.Stat(filename); err == nil {
				b, err := ioutil.ReadFile(filename)
				if err != nil {
					return l
				}
				body := string(b)
				lines := strings.Split(body, "\n")
				if len(lines) > 3 {
					body = strings.Join(lines[3:], "\n")
				}
				l.Detail = body
			}
			return l
		}
	}
	return lineModel{}
}

func writeLine(day string, outline string, tags []string, detail string) error {
	month := day[0:4] + day[5:7]

	// monthファイルの更新
	lines := readMonthFile(month)
	flag := false
	for i, l := range lines {
		if day == l.Day {
			lines[i].Outline = outline
			lines[i].Tags = tags
			flag = true
		}
		lines[i].Outline = strings.TrimSpace(lines[i].Outline)
	}
	if !flag || len(lines) == 0 {
		lines = append(lines, lineModel{day, outline, []string{}, "", 0})
	}

	// 詳細ファイルの更新
	filename := strings.ReplaceAll(day, "-", "") + ".txt"
	filename = path.Join(diaryPath, filename)
	if len(detail) == 0 {
		if _, err := os.Stat(filename); err == nil {
			os.Remove(filename)
		}
		return writeMonthFile(month, lines)
	}

	data := day + "\n" + outline + "\n" + strings.Join(tags, "#") + "\n" + detail
	ioutil.WriteFile(filename, []byte(data), fs.ModePerm)
	return writeMonthFile(month, lines)
}

func deleteLine(day string) error {
	month := day[0:4] + day[5:7]
	lines := readMonthFile(month)
	for i, l := range lines {
		if day == l.Day {
			lines = append(lines[:i], lines[i+1:]...)
			break
		}
	}

	filename := strings.ReplaceAll(day, "-", "") + ".txt"
	filename = path.Join(diaryPath, filename)
	if _, err := os.Stat(filename); err == nil {
		os.Remove(filename)
	}

	return writeMonthFile(month, lines)
}

func readMonthFile(month string) []lineModel {
	p := filepath.Join(diaryPath, month+".txt")

	if _, err := os.Stat(p); err != nil {
		return []lineModel{}
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
		day := s[0:10]
		s = s[11:]

		tags := strings.Split(s, "#")
		if len(tags) > 1 {
			d := lineModel{day, tags[len(tags)-1], tags[:len(tags)-1], "", 0}
			lines = append(lines, d)
		} else {
			d := lineModel{day, s, []string{}, "", 0}
			lines = append(lines, d)
		}
	}

	return lines
}

func writeMonthFile(month string, lines []lineModel) error {
	sort.Slice(lines, func(i, j int) bool { return lines[i].Day > lines[j].Day })

	dataFile := ""
	for _, l := range lines {
		if dataFile != "" {
			dataFile += "\n"
		}
		l.Outline = strings.ReplaceAll(l.Outline, "#", "")
		dataFile += l.Day + " "
		if len(l.Tags) > 0 {
			dataFile += strings.Join(l.Tags, "#") + "#"
		}
		dataFile += l.Outline
	}
	filename := filepath.Join(diaryPath, month+".txt")

	if len(lines) == 0 {
		return os.Remove(filename)
	}
	return ioutil.WriteFile(filename, []byte(dataFile), os.ModePerm)
}
