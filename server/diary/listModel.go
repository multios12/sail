package diary

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
)

type listModel struct {
	WritedMonths []string // 記載された年月(yyyy-mm形式)
	Lines        []lineModel
}

func readListFile(month string) *listModel {
	p := filepath.Join(diaryPath, month+".txt")

	l := new(listModel)
	if _, err := os.Stat(p); err != nil {
		l.Lines = []lineModel{}
		return l
	}

	fp, err := os.Open(p)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	reader := bufio.NewReaderSize(fp, 5000)
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
		l.Lines = append(l.Lines, *d)
	}

	return l
}

func (l *listModel) writeListFile(month string) error {
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

func (l *listModel) updateLine(month string, target lineModel) {
	for i, line := range l.Lines {
		if target.Day == line.Day {
			l.Lines[i] = target
			l.writeListFile(month)
			return
		}
	}
	l.Lines = append(l.Lines, target)
	l.writeListFile(month)
}
