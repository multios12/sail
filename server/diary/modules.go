package diary

import (
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"sort"
	"strings"
	"time"
)

var diaryPath string

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
func readDetail(day string) detailModel {
	m := newListModel(day[0:4] + day[5:7])
	for _, l := range m.Lines {
		if day == l.Day {
			return *newDetailModel(l)
		}
	}
	return detailModel{}
}

func writeDetail(line detailModel) error {
	month := line.Day[0:4] + line.Day[5:7]

	// monthファイルの更新
	m := newListModel(month)
	flag := false
	for i, l := range m.Lines {
		if line.Day == l.Day {
			m.Lines[i].Outline = line.Outline
			m.Lines[i].Tags = line.Tags
			flag = true
		}
		m.Lines[i].IsDetail = len(line.Detail) > 0
		m.Lines[i].Outline = strings.TrimSpace(m.Lines[i].Outline)
	}
	if !flag || len(m.Lines) == 0 {
		m.Lines = append(m.Lines, lineModel{line.Day, line.Outline, []string{}, len(line.Detail) > 0, 0})
	}

	// 詳細ファイルの更新
	line.WriteDetail()
	return m.writeMonthFile(month)
}

func removeDetail(day string) error {
	month := day[0:4] + day[5:7]
	m := newListModel(month)
	for i, l := range m.Lines {
		if day == l.Day {
			m.Lines = append(m.Lines[:i], m.Lines[i+1:]...)
			break
		}
	}

	filename := strings.ReplaceAll(day, "-", "") + ".txt"
	filename = path.Join(diaryPath, filename)
	if _, err := os.Stat(filename); err == nil {
		os.Remove(filename)
	}

	return m.writeMonthFile(month)
}
