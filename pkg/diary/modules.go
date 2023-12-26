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
		months = append(months, time.Now().Format("2006-01"))
	}

	return months
}

// ----------------------------------------------------------------------------
func readDetail(day string) detailModel {
	m := readListFile(day[0:4] + day[5:7])
	for _, l := range m.Lines {
		if day == l.Day {
			return *readDetailFile(l)
		}
	}
	return detailModel{}
}

func removeDetail(day string) error {
	month := day[0:4] + day[5:7]
	m := readListFile(month)
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

	return m.writeListFile(month)
}
