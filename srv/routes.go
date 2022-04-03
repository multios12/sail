package main

import (
	"embed"
	"net/http"
	"path/filepath"

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
	router.GET("/api/:year/:month/detailImage", getDetailImage)
	return router
}

func getStatic(c *gin.Context) {
	c.FileFromFS("static"+c.Request.URL.Path, http.FS(static))
}

// 年単位データ GET API
func getYear(c *gin.Context) {
	y := YearModel{Year: c.Param("year")}

	for _, detail := range details {
		if detail.Month[:3] != y.Year {
			continue
		}

		y.Details = append(y.Details, detail)
		for _, total := range detail.Totals {
			exist := false
			for _, yearTotal := range y.Totals {
				if total.Name == yearTotal.Name {
					exist = true
					yearTotal.Value += total.Value
					break
				}
			}
			if exist {
				y.Totals = append(y.Totals, DetailItem{Name: total.Name, Value: total.Value})
			}
		}
	}
}

// 月単位データ GET API
func getMonth(c *gin.Context) {
	m := c.Param("year") + c.Param("month")
	var l []DetailModel
	for _, v := range details {
		if v.Month == m {
			l = append(l, v)
		}
	}
	c.JSON(200, l)
}

// 月単位画像 GET API
func getDetailImage(c *gin.Context) {
	m := c.Param("year") + c.Param("month")
	filename := filepath.Join(dataPath, m)
	c.File(filename)
}
