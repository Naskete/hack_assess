package dao

import (
	"HackAssess/internal/model"
	"github.com/gin-gonic/gin"
)

// 数据库操作

//Mark 评分
func Mark(c *gin.Context) error{
	// TODO 存入数据库
	var score model.Score
	err := c.ShouldBind(&score)
	if err != nil {
		return err
	}
	if err := d.DB.Create(&score).Error; err!= nil{
		return err
	}
	return nil
}
