package http

import "github.com/gin-gonic/gin"

func Init(){
	e := gin.Default()
	e.GET("/getDate/:i", func(context *gin.Context) {
		context.JSONP(200,gin.H{
			"status": 1,
			"message":"ok",
			"content": context.Param("i"),
		})
	})
	e.Run(":8080")
}

func getScoreData(e gin.Engine) {
	// url/word/i ==> 第i组

	e.POST("/login")
	e.POST("/score")
}