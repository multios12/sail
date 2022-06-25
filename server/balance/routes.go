package balance

import (
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var salaryPath string  // データディレクトリ
var pdfPassword string // PDFパスワード

func Initial(router *gin.Engine, dataPath string, password string) {
	salaryPath = path.Join(dataPath, "salary")
	pdfPassword = password
	dbOpen(dataPath)

	var err error
	Salaries, err = readAllData()
	if err != nil {
		println(err)
		return
	}

	router.GET("/api/balance/:year", getBalanceYear)
	router.GET("/api/balance/:year/:month", getBalanceMonth)
	router.POST("/api/balance/:year/:month", postBalanceMonth)

	router.GET("/api/salary/:year", getSalaryYear)
	router.GET("/api/salary/:year/:month", getSalaryMonth)
	router.PUT("/api/salary/:year/:month", putSalaryMonth)
	router.GET("/api/salary/:year/:month/images/:file", getSalaryDetailImage)

	router.POST("/api/salary/files", postFiles)
}

// バランスシート年単位データ GET API
func getBalanceYear(c *gin.Context) {
	balances, err := findBalanceByYear(c.Param("year"))
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	model := BalanceYear{Year: c.Param("year"), Balances: balances}
	var now = time.Now().AddDate(0, 1, 0)
	for i := now.Year(); i >= 2019; i-- {
		model.EnableYears = append(model.EnableYears, strconv.Itoa(i))
	}
	c.JSON(http.StatusOK, model)
}

// バランスシート月単位データ GET API
func getBalanceMonth(c *gin.Context) {
	m := c.Param("year") + c.Param("month")

	balance, err := findBalanceByMonth(m)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, balance)
}

// バランスシート月単位データ POST API
func postBalanceMonth(c *gin.Context) {
	value := c.Param("year") + c.Param("month")
	var regmonth = regexp.MustCompile(`^[0-9]{6}$`)
	if !regmonth.MatchString(value) {
		c.Status(http.StatusNotFound)
		return
	}

	var body Balance
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	target, err := findBalanceByMonth(value)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	target.CostWater = body.CostWater
	target.CostElectric = body.CostElectric
	target.CostGas = body.CostGas
	target.CostMobile = body.CostMobile
	target.CostLine = body.CostLine
	target.CostTax = body.CostTax
	target.Cost = body.CostWater + body.CostElectric + body.CostGas + body.CostMobile + body.CostLine + body.CostTax
	target.Saving = body.Saving
	target.Memo = body.Memo

	var nowMonth = strconv.Itoa(int(time.Now().Month()))
	if len(nowMonth) == 1 {
		nowMonth = "0" + nowMonth
	}
	nowMonth = strconv.Itoa(time.Now().Year()) + nowMonth

	target.IsNotCost = body.Month < nowMonth && (body.CostElectric == 0 || body.CostGas == 0 || body.CostMobile == 0 || body.CostLine == 0)
	upsertBalance(target)
	c.Status(http.StatusOK)
}

// 給与年単位データ GET API
func getSalaryYear(c *gin.Context) {
	y := SalaryYear{Year: c.Param("year")}

	for _, detail := range Salaries {
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

	for _, detail := range Salaries {
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
	for _, v := range Salaries {
		if v.Month == m {
			c.JSON(http.StatusOK, v)
			return
		}
	}
	c.Status(http.StatusNotFound)
}

// 給与月単位データ再作成 PUT API
func putSalaryMonth(c *gin.Context) {
	m := c.Param("year") + c.Param("month")
	filename := filepath.Join(salaryPath, m)

	if _, err := os.Stat(filename); err == nil {
		err := os.RemoveAll(filename)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
	}

	// PDFファイルの変換とデータ再読み込み
	Convert(salaryPath)
	var err error
	Salaries, err = readAllData()
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	updateBalanceFromSalaries(m[:6], Salaries)

	for _, v := range Salaries {
		if v.Month == m {
			c.JSON(http.StatusOK, v)
			return
		}
	}
	c.Status(http.StatusNotFound)
}

// 月単位画像 GET API
func getSalaryDetailImage(c *gin.Context) {
	m := c.Param("year") + c.Param("month")
	filename := filepath.Join(salaryPath, m, c.Param("file"))
	c.File(filename)
}

// ファイル保存 POST API
func postFiles(c *gin.Context) {
	inFile, header, err := c.Request.FormFile("file")
	if err != nil {
		return
	}

	filename := filepath.Join(salaryPath, header.Filename)
	outFile, err := os.Create(filename)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, inFile)
	if err != nil {
		c.String(http.StatusInternalServerError, "ファイルを保存できません")
		return
	}

	f, _ := os.Stat(filename)
	err = createData(salaryPath, f)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	Salaries, err = readAllData()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}
