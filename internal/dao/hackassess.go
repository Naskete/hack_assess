package dao

import (
	"HackAssess/internal/model"
	"github.com/gin-gonic/gin"
)

// 数据库操作

// Mark 评分
// TODO 优化
// TODO 拿到UserID 做唯一约束或删除UserID 评价后直接关闭 或再换一个标志做唯一约束，没时间就迭代
func Mark(c *gin.Context) error {
	// 得拿到UserID
	var err error
	group := c.Param("group")
	switch group {
	case "director":
		var score model.Director
		err = c.ShouldBind(&score)
		if e := d.DB.Create(&score).Error; e != nil {
			return e
		}
	case "product":
		var score model.Product
		err = c.ShouldBind(&score)
		if e := d.DB.Create(&score).Error; e != nil {
			return e
		}
	case "design":
		var score model.Design
		err = c.ShouldBind(&score)
		if e := d.DB.Create(&score).Error; e != nil {
			return e
		}
	case "front":
		var score model.Front
		err = c.ShouldBind(&score)
		if e := d.DB.Create(&score).Error; e != nil {
			return e
		}
	case "back":
		var score model.Back
		err = c.ShouldBind(&score)
		if e := d.DB.Create(&score).Error; e != nil {
			return e
		}
	case "show":
		var score model.Director
		err = c.ShouldBind(&score)
		if e := d.DB.Create(&score).Error; e != nil {
			return e
		}
	}
	if err != nil {
		return err
	}

	//var score model.Score
	//var err error
	//err = c.ShouldBind(&score)
	//if err != nil {
	//	return err
	//}
	//if err := d.DB.Create(&score).Error; err != nil {
	//	return err
	//}
	return nil
}
