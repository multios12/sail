package main

import (
	"embed"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// gin.Engineインスタンスにルーティングを設定して返す
func routerInitial(static embed.FS) *gin.Engine {
	router := gin.Default()

	router.GET("/", getStatic)
	router.GET("/favicon.ico", getStatic)
	router.GET("/static/:dir/:file", getStatic)

	router.GET("/api/:year", getBalanceYear)

	router.GET("/api/salary/:year", getSalaryYear)
	router.GET("/api/salary/:year/:month", getSalaryMonth)
	router.PUT("/api/salary/:year/:month", putSalaryMonth)
	router.GET("/api/salary/:year/:month/images/:file", getSalaryDetailImage)

	router.GET("/api/cost/:year", getCostYear)
	router.GET("/api/cost/:year/:month", getCostYearMonth)
	router.POST("/api/cost/:year/:month", postCostYearMonth)

	router.POST("/api/files", postFiles)
	return router
}

// スタティックリソース GET API
func getStatic(c *gin.Context) {
	c.FileFromFS("static"+c.Request.URL.Path, http.FS(static))
}

// 総表示データ GET API
func getBalanceYear(c *gin.Context) {
	m := map[string]BalanceItem{}

	for _, salary := range salaries {
		if c.Param("year") != salary.Month[:4] {
			continue
		}

		item, ok := m[salary.Month[:6]]
		if !ok {
			item = BalanceItem{Month: salary.Month[:6]}
		}

		item.Salary += salary.Totals[0].Value + salary.Expense
		item.Paid += salary.Totals[2].Value + salary.Expense
		if len(salary.Month) == 7 {
			item.Memo = "＋賞与"
		}
		m[salary.Month[:6]] = item
	}

	costs, err := findCostByYear(c.Param("year"))
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	for _, c := range costs {
		item, ok := m[c.Month]
		if !ok {
			item = BalanceItem{Month: c.Month}
		}

		item.Cost = c.Water + c.Electric + c.Gas + c.Mobile + c.Line
		var nowMonth = strconv.Itoa(int(time.Now().Month()))
		if len(nowMonth) == 1 {
			nowMonth = "0" + nowMonth
		}
		nowMonth = strconv.Itoa(time.Now().Year()) + nowMonth
		item.IsNotCost = c.Month < nowMonth && (c.Electric == 0 || c.Gas == 0 || c.Mobile == 0 || c.Line == 0)
		m[c.Month] = item
	}

	balances := []BalanceItem{}
	for _, v := range m {
		balances = append(balances, v)
	}
	sort.Slice(balances, func(i, j int) bool { return balances[i].Month > balances[j].Month })

	model := BalanceYear{Year: c.Param("year"), Balances: balances}
	var now = time.Now().AddDate(0, 1, 0)
	for i := now.Year(); i >= 2019; i-- {
		model.EnableYears = append(model.EnableYears, strconv.Itoa(i))
	}
	c.JSON(http.StatusOK, model)
}

// 年単位データ GET API
func getSalaryYear(c *gin.Context) {
	y := SalaryYearModel{Year: c.Param("year")}

	for _, detail := range salaries {
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

	for _, detail := range salaries {
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
func getSalaryMonth(c *gin.Context) {
	m := c.Param("year") + c.Param("month")
	for _, v := range salaries {
		if v.Month == m {
			c.JSON(200, v)
			return
		}
	}
	c.Status(404)
}

// 月単位データ再作成 PUT API
func putSalaryMonth(c *gin.Context) {
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
	salaries, err = readAllData(dataPath)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	c.Status(http.StatusOK)
}

// 月単位画像 GET API
func getSalaryDetailImage(c *gin.Context) {
	m := c.Param("year") + c.Param("month")
	filename := filepath.Join(dataPath, m, c.Param("file"))
	c.File(filename)
}

// 年単位固定支出データ GET API
func getCostYear(c *gin.Context) {
	if costs, err := findCostByYear(c.Param("year")); err != nil {
		c.Status(http.StatusNotFound)
	} else {
		var costYear = SumCost{Year: c.Param(("year")), Costs: costs}
		var now = time.Now().AddDate(0, 1, 0)
		for i := now.Year(); i >= 2020; i-- {
			costYear.EnableYears = append(costYear.EnableYears, strconv.Itoa(i))
		}
		c.JSON(http.StatusOK, costYear)
	}
}

// 月単位固定支出データ GET API
func getCostYearMonth(c *gin.Context) {
	value := c.Param("year") + c.Param("month")
	if cost, err := findCostByMonth(value); err != nil {
		c.Status(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, cost)
	}
}

// 月単位固定支出データ POST API
func postCostYearMonth(c *gin.Context) {
	value := c.Param("year") + c.Param("month")
	var regmonth = regexp.MustCompile(`^[0-9]{6}$`)
	if !regmonth.MatchString(value) {
		c.Status(http.StatusNotFound)
		return
	}

	var cost Cost
	if err := c.ShouldBindJSON(&cost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	upsertCost(cost)
	c.Status(http.StatusOK)
}

// ファイル保存 POST API
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
	salaries, err = readAllData(dataPath)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}
