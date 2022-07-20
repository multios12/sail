package diary

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

// gin.Engineインスタンスにルーティングを設定して返す
func Initial(router *gin.Engine, dataPath string) {
	diaryPath, _ = filepath.Abs(dataPath)
	diaryPath = filepath.Join(diaryPath, "diary")
	if _, err := os.Stat(diaryPath); err != nil {
		os.Mkdir(diaryPath, 0777)
	}
	router.GET("/api/diary/:year/:month", getMonth)
	router.GET("/api/diary/:year/:month/:day", getDay)
	router.POST("/api/diary/:year/:month/:day", postDay)
	router.POST("/api/diary/detail", postDetail)
}

func getMonth(c *gin.Context) {
	m := c.Param("year") + c.Param("month")
	var l = listModel{Lines: readMonthFile(m), WritedMonths: getWritedMonths()}
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

func postDetail(c *gin.Context) {
	var d detail
	err := c.ShouldBindJSON(&d)
	if err == nil {
		err = writeDetail(d)
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.Status(http.StatusOK)
	}
}
