package dao

import (
	"HackAssess/internal/model"
	"github.com/gin-gonic/gin"
)

// Mark 评分
func Mark(c *gin.Context) error {
	group := c.Param("group")
	switch group {
	case "director":
		var score model.Director
		_ = c.ShouldBind(&score)
		score.Position = group
		if e := d.DB.Create(&score).Error; e != nil {
			return e
		}
	case "product":
		var score model.Product
		_ = c.ShouldBind(&score)
		score.Position = group
		if e := d.DB.Create(&score).Error; e != nil {
			return e
		}
	case "design":
		var score model.Design
		_ = c.ShouldBind(&score)
		score.Position = group
		if e := d.DB.Create(&score).Error; e != nil {
			return e
		}
	case "front":
		var score model.Front
		_ = c.ShouldBind(&score)
		score.Position = group
		if e := d.DB.Create(&score).Error; e != nil {
			return e
		}
	case "back":
		var score model.Back
		_ = c.ShouldBind(&score)
		score.Position = group
		if e := d.DB.Create(&score).Error; e != nil {
			return e
		}
	case "show":
		var score model.Director
		_ = c.ShouldBind(&score)
		score.Position = group
		if e := d.DB.Create(&score).Error; e != nil {
			return e
		}
	}
	return nil
}
