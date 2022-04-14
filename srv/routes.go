package main

import (
	"embed"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"

	"github.com/gin-gonic/gin"
)

// gin.Engineインスタンスにルーティングを設定して返す
func routerInitial(static embed.FS) *gin.Engine {
	router := gin.Default()

	router.GET("/", getStatic)
	router.GET("/favicon.ico", getStatic)
	router.GET("/static/:dir/:file", getStatic)

	router.GET("/api/:year", getYear)
	router.GET("/api/:year/:month", getMonth)
	router.PUT("/api/:year/:month", putMonth)
	router.GET("/api/:year/:month/images/:file", getDetailImage)
	router.POST("/api/files", postFiles)
	return router
}

func getStatic(c *gin.Context) {
	c.FileFromFS("static"+c.Request.URL.Path, http.FS(static))
}

// 年単位データ GET API
func getYear(c *gin.Context) {
	y := YearModel{Year: c.Param("year")}

	for _, detail := range details {
		isEnabled := false
		for _, e := range y.EnableYears {
			if e == detail.Month[:4] {
				isEnabled = true
			}
		}
		if !isEnabled {
			y.EnableYears = append(y.EnableYears, detail.Month[:4])
		}
	}

	for _, detail := range details {
		if detail.Month[:4] != y.Year {
			continue
		}

		y.Details = append(y.Details, detail)
		for _, total := range detail.Totals {
			exist := false
			for i, yearTotal := range y.Totals {
				if total.Name == yearTotal.Name {
					exist = true
					y.Totals[i].Value += total.Value
					break
				}
			}
			if !exist {
				y.Totals = append(y.Totals, DetailItem{Name: total.Name, Value: total.Value})
			}
		}
	}

	sort.Slice(y.EnableYears, func(i, j int) bool { return y.EnableYears[i] > y.EnableYears[j] })
	sort.Slice(y.Details, func(i, j int) bool { return y.Details[i].Month > y.Details[j].Month })
	c.JSON(200, y)
}

// 月単位データ GET API
func getMonth(c *gin.Context) {
	m := c.Param("year") + c.Param("month")
	for _, v := range details {
		if v.Month == m {
			c.JSON(200, v)
			return
		}
	}
	c.Status(404)
}

// 月単位データ再作成 PUT API
func putMonth(c *gin.Context) {
	m := c.Param("year") + c.Param("month")
	filename := filepath.Join(dataPath, m)

	if _, err := os.Stat(filename); err == nil {
		err := os.RemoveAll(filename)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
	}

	// PDFファイルの変換とデータ再読み込み
	convert()
	var err error
	details, err = readAllData(dataPath)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	c.Status(http.StatusOK)
}

// 月単位画像 GET API
func getDetailImage(c *gin.Context) {
	m := c.Param("year") + c.Param("month")
	filename := filepath.Join(dataPath, m, c.Param("file"))
	c.File(filename)
}

func postFiles(c *gin.Context) {
	inFile, header, err := c.Request.FormFile("file")
	if err != nil {
		return
	}

	filename := filepath.Join(dataPath, header.Filename)
	outFile, err := os.Create(filename)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, inFile)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	convert()
	details, err = readAllData(dataPath)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}
