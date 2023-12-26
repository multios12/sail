package memo

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Initial(router *gin.Engine, dataPath string) {
	// DB初期化
	err := dbOpen(dataPath)
	if err != nil {
		fmt.Sprintln(err)
		return
	}
	memos := find()
	log.Println("info: memo[dataPath=" + filepath.Join(dataPath, "hmemo") + "]")
	log.Println("info: memo[count=" + strconv.Itoa(len(memos)) + "]")

	// ルーティング
	router.GET("/api/memos", getMemos)
	router.GET("/api/memos/:id", getMemosId)
	router.PUT("/api/memos", putMemos)
	router.POST("/api/memos/:id", postMemosId)
	router.DELETE("/api/memos/:id", deleteMemosId)
}
func getMemos(c *gin.Context) {
	var memos []Memo
	if len(c.Query("month")) == 0 {
		memos = find()
	} else {
		memos = findByMonth(c.Query("month"))
	}
	createResponse(c, memos, nil)
}

func getMemosId(c *gin.Context) {
	memos := findById(c.Param("id"))
	createResponse(c, memos, nil)
}

func putMemos(c *gin.Context) {
	var b Memo
	err := c.ShouldBindJSON(&b)
	if err == nil {
		upsertMemo(b)
	}
	createResponse(c, nil, err)
}

func postMemosId(c *gin.Context) {
	var b Memo
	err := c.ShouldBindJSON(&b)
	if err == nil {
		validate := validator.New() //インスタンス生成
		err = validate.Struct(b)
		if err == nil {
			b.Id, _ = strconv.Atoi(c.Param("id"))
			upsertMemo(b)
		}
	}
	createResponse(c, nil, err)
}

func deleteMemosId(c *gin.Context) {
	deleteMemo(c.Param("id"))
	createResponse(c, nil, nil)
}

func createResponse(c *gin.Context, memos []Memo, err error) {
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, memos)
}
