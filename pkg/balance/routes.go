package balance

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/multios12/sail/pkg/balance/converter"
)

var balancePath string // データディレクトリ
var pdfPassword string // PDFパスワード

func Initial(router *gin.Engine, dataPath string, password string) {
	balancePath = path.Join(dataPath, "balance")
	pdfPassword = password
	dbOpen(dataPath)

	// PDFファイルのチェック、追加されていれば再読み込み
	err := readSalaries(balancePath, pdfPassword)
	if err != nil {
		println(err)
		return
	}

	log.Println("info: balance:salaryPath:" + balancePath)
	log.Println("info: balance:dbPath    :" + filepath.Join(balancePath, "db"))
	//	log.Println("info: balance:data[count:" + strconv.Itoa(countSalary()) + "]")

	// バランスシートAPI
	router.GET("/api/balance/:year", getBalanceYear)
	router.GET("/api/balance/:year/:month", getBalanceMonth)
	router.POST("/api/balance/:year/:month", postBalanceMonth)

	// 給与明細API
	router.GET("/api/salary/:year", getSalaryYear)
	router.GET("/api/salary/:year/:month", getSalaryMonth)
	router.PUT("/api/salary/:year/:month", putSalaryMonth)
	router.POST("/api/salary/:year/:month", postSalaryMonth)
	router.GET("/api/salary/:year/:month/images/:file", getSalaryDetailImage)
	router.POST("/api/salary/files", postSalaryFiles)
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
	target.CostHousing = body.CostHousing
	target.CostWater = body.CostWater
	target.CostElectric = body.CostElectric
	target.CostGas = body.CostGas
	target.CostMobile = body.CostMobile
	target.CostLine = body.CostLine
	target.CostTax = body.CostTax
	target.Cost = body.CostHousing + body.CostWater + body.CostElectric + body.CostGas + body.CostMobile + body.CostLine + body.CostTax
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

// ----------------------------------------------------------------------------

// 給与年単位データ GET API
func getSalaryYear(c *gin.Context) {
	y := SalaryYear{Year: c.Param("year")}
	y.EnableYears = findSalaryYears()

	balances, _ := findBalanceByYear(y.Year)
	for _, balance := range balances {
		d := balance.SalaryDetail()
		if len(d.Month) == 0 || d.Month[:4] != y.Year {
			continue
		}
		d.Expense = balance.Expense
		d.Title = d.Month[:4] + "年" + d.Month[4:] + "月 給与"
		y.Details = append(y.Details, d)

		b := balance.BonusDetail()
		if len(b.Month) > 0 && b.Month[:4] == y.Year {
			b.Title = b.Month[:4] + "年" + b.Month[4:6] + "月 賞与"
			y.Details = append(y.Details, b)
		}

		t := d.Totals
		t = append(t, b.Totals...)
		for _, total := range t {
			exist := false
			for i, yearTotal := range y.Totals {
				if total.Name == yearTotal.Name {
					exist = true
					y.Totals[i].Value += total.Value
					break
				}
			}
			if !exist {
				y.Totals = append(y.Totals, converter.DetailItem{Name: total.Name, Value: total.Value})
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
	if strings.Contains(m, "S") {
		m = m[:6]
		if s, _ := findBalanceByMonth(m); len(s.Month) > 0 {
			c.JSON(http.StatusOK, s.BonusDetail())
			return
		}
	} else if s, _ := findBalanceByMonth(m); len(s.Month) > 0 {
		d := s.SalaryDetail()
		d.Expense = s.Expense
		c.JSON(http.StatusOK, s.SalaryDetail())
		return
	}
	c.Status(http.StatusNotFound)
}

// 月単位データ POST API
func postSalaryMonth(c *gin.Context) {
	m := c.Param("year") + c.Param("month")
	if strings.Contains(m, "S") {
		m = m[:6]
	}

	deleteBalanceDetail(m)
	readSalaries(balancePath, pdfPassword)
	c.Status(http.StatusOK)
}

// 給与月単位データ再作成 PUT API
func putSalaryMonth(c *gin.Context) {
	m := c.Param("year") + c.Param("month")
	filename := filepath.Join(balancePath, m)

	if _, err := os.Stat(filename); err == nil {
		err := os.RemoveAll(filename)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
	}

	// PDFファイルの変換とデータ再読み込み
	readSalary(balancePath, pdfPassword)
	if s, _ := findBalanceByMonth(m); len(s.Month) > 0 {
		c.JSON(http.StatusOK, s.SalaryDetail())
		return
	}
	c.Status(http.StatusNotFound)
}

// 月単位画像 GET API
func getSalaryDetailImage(c *gin.Context) {
	m := c.Param("year") + c.Param("month")
	if b, _ := findBalanceByMonth(m); len(b.Month) > 0 {
		filename := c.Param("file")
		i, _ := strconv.Atoi(filename[:1])
		c.Data(http.StatusOK, "image/png", b.Image(DetailType(i)))
		return
	}
	c.Status(http.StatusNotFound)
}

// 給与ファイル保存 POST API
func postSalaryFiles(c *gin.Context) {
	inFile, header, err := c.Request.FormFile("file")
	if err != nil {
		return
	}

	filename := filepath.Join(balancePath, "_pdf", header.Filename)
	outFile, err := os.Create(filename)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	_, err = io.Copy(outFile, inFile)
	if err != nil {
		err = fmt.Errorf("ファイルが保存できません: %w", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	outFile.Close()

	readSalary(filename, pdfPassword)
	c.Status(http.StatusOK)
}
