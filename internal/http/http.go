package http

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Init() {
	e := gin.Default()
	getScoreData(e)
	e.Run(":8080")
}

func getScoreData(e *gin.Engine) {
	// url/word/i ==> 第i组

	// TODO python grpc 登录接口
	// TODO filter 只允许评委打分
	e.POST("/login", func(c *gin.Context) {
		if result, e := login(c); e != nil {
			log.Printf("can not login : %v", e)
			c.AbortWithStatusJSON(http.StatusBadRequest, result)
		} else {
			c.JSON(http.StatusOK, result)
		}
	})
	// TODO 评分优化
	// 评分上限(迭代）
	e.POST("/:group/:id", func(c *gin.Context) {
		if result, e := giveMarks(c); e != nil {
			log.Printf("can not mark : %v", e)
			c.AbortWithStatusJSON(http.StatusInternalServerError, result)
		} else {
			c.JSON(http.StatusOK, result)
		}
	})
}
