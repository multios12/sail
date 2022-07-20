package diary

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

var diaryPath string

type lineModel struct {
	Day  string // 日付(yyyy-mm-dd形式)
	Memo string // 本文
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

	r := regexp.MustCompile(`\d\d\d\d\d\d\.txt`)

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
			return l
		}
	}
	return lineModel{}
}

func writeLine(day string, memo string) error {
	month := day[0:4] + day[5:7]
	lines := readMonthFile(month)
	flag := false
	for i, l := range lines {
		if day == l.Day {
			lines[i].Memo = memo
			flag = true
		}
		lines[i].Memo = strings.TrimSpace(lines[i].Memo)
	}
	if !flag || len(lines) == 0 {
		lines = append(lines, lineModel{day, memo})
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

		d := lineModel{day, s}
		lines = append(lines, d)
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
		dataFile += l.Day + " " + l.Memo
	}
	p := filepath.Join(diaryPath, month+".txt")
	return ioutil.WriteFile(p, []byte(dataFile), os.ModePerm)
}
