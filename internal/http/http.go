package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init() {
	e := gin.Default()
//	e.GET("/getDate/:i", func(c *gin.Context) {
//		c.JSONP(200, gin.H{
//			"status":  1,
//			"message": "ok",
//			"content": c.Param("i"),
//		})
//	})
	getScoreData(e)
	e.Run(":8080")
}

func getScoreData(e *gin.Engine) {
	// url/word/i ==> 第i组

	// TODO python grpc 接口
	e.GET("/login", func(c *gin.Context) {

	})
	// TODO 评分
	e.POST("/:group/:id", func(c *gin.Context) {
		if result, e:= GiveMarks(c); e!= nil{
			c.AbortWithStatusJSON(http.StatusInternalServerError, result)
		} else {
			c.JSON(http.StatusOK, result)
		}
	})
}
