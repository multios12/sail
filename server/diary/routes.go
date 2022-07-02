package main

import (
	"bufio"
	"embed"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
)

// gin.Engineインスタンスにルーティングを設定して返す
func Initial(router *gin.Engine, static embed.FS, dataPath string) {
	dataDir, _ = filepath.Abs(dataDir)
	dataDir = filepath.Join(dataDir, "diary")
	if _, err := os.Stat(dataDir); err != nil {
		os.Mkdir(dataDir, 0777)
	}
	router.GET("/api/:year/:month", getMonth)
	router.GET("/api/:year/:month/:day", getDay)
	router.POST("/api/:year/:month/:day", postDay)
}

func getMonth(c *gin.Context) {
	m := c.Param("year") + c.Param("month")
	var l = listModel{Lines: readMonthFile(m), WritedMonths: months}
	c.JSON(200, l)
}

func getDay(c *gin.Context) {
	day := c.Param("year") + "-" + c.Param("month") + "-" + c.Param("day")
	if l := readLine(day); len(l.Day) == 0 {
		c.Status(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, l)
	}
}

func postDay(c *gin.Context) {
	day := c.Param("year") + "-" + c.Param("month") + "-" + c.Param("day")

	var line lineModel
	if err := c.ShouldBindJSON(&line); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else if line.Memo = strings.TrimSpace(line.Memo); line.Memo == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "memo is not found."})
	} else {
		writeLine(day, line.Memo)
		c.Status(http.StatusOK)
	}
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
	if flag {
		lines = append(lines, lineModel{day, memo})
	}
	return writeMonthFile(month, lines)
}

func readMonthFile(month string) []lineModel {
	p := filepath.Join(dataDir, month+".txt")

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
	p := filepath.Join(dataDir, month+".txt")
	return ioutil.WriteFile(p, []byte(dataFile), os.ModePerm)
}
