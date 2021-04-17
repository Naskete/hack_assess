package http

import (
	"github.com/gin-gonic/gin"
	"github.com/ncuhome/us-backend/feedback/pkg"
	"net/http"
)

func Init() {
	e := gin.Default()
	e.Use(pkg.Cors())
	getScoreData(e)
	_ = e.Run(":8088")
}

// TODO 导表
func getScoreData(e *gin.Engine) {
	e.POST("/login", func(c *gin.Context) {
		if result, e := login(c); e != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, result)
		} else {
			c.JSON(http.StatusOK, result["data"])
		}
	})
	// 返回值优化
	// 评分只能打一次，如需修改？？？（new demand) 再议!!!
	e.POST("/:group/:id", func(c *gin.Context) {
		if result, e := giveMarks(c); e != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, result)
		} else {
			c.JSON(http.StatusOK, result)
		}
	})
}
