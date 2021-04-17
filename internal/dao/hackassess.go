package dao

import (
	"HackAssess/internal/model"
	"HackAssess/pkg"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Mark 评分
func Mark(c *gin.Context) error {
	group := c.Param("group")
	// TODO 优化代码结构
	switch group {
	case "director":
		var score model.Director
		_ = c.ShouldBind(&score)
		score.Position = group
		score.Tag = strconv.Itoa(score.UserId) + "-" + score.Position + "-" + score.GroupID
		total := pkg.Sum(score.Product, score.Idea, score.Integrity, score.Interactive, score.Planning, score.Promotion, score.Slogan, score.Research)
		score.Total = total
		if e := d.DB.Create(&score).Error; e != nil {
			return e
		}
		// TODO research 0
	case "product":
		var score model.Product
		_ = c.ShouldBind(&score)
		score.Position = group
		score.Tag = strconv.Itoa(score.UserId) + "-" + score.Position + "-" + score.GroupID
		total := pkg.Sum(score.Innovation, score.Rationality, score.Conformity, score.Research, score.Requirements, score.Prototype, score.Competing, score.Req, score.Project, score.Plan)
		score.Total = total
		if e := d.DB.Create(&score).Error; e != nil {
			return e
		}
	case "design":
		var score model.Design
		_ = c.ShouldBind(&score)
		score.Position = group
		score.Tag = strconv.Itoa(score.UserId) + "-" + score.Position + "-" + score.GroupID
		total := pkg.Sum(score.Design, score.Logo, score.Beauty, score.Uniformity, score.Friendliness, score.Reliability, score.Font, score.Thinking, score.Bonus)
		score.Total = total
		if e := d.DB.Create(&score).Error; e != nil {
			return e
		}
	case "front":
		var score model.Front
		_ = c.ShouldBind(&score)
		score.Position = group
		score.Tag = strconv.Itoa(score.UserId) + "-" + score.Position + "-" + score.GroupID
		total := pkg.Sum(score.Function, score.Fit, score.Layout, score.Encapsulation, score.Resources, score.Docking)
		score.Total = total
		if e := d.DB.Create(&score).Error; e != nil {
			return e
		}
	case "back":
		var score model.Back
		_ = c.ShouldBind(&score)
		score.Position = group
		score.Tag = strconv.Itoa(score.UserId) + "-" + score.Position + "-" + score.GroupID
		total := pkg.Sum(score.Integrity, score.Property, score.Scalability, score.Cert, score.Configuration, score.Information, score.Code, score.Database, score.Interface, score.Document, score.Coverage, score.Solution)
		score.Total = total
		if e := d.DB.Create(&score).Error; e != nil {
			return e
		}
	case "show":
		var score model.Show
		_ = c.ShouldBind(&score)
		score.Position = group
		score.Tag = strconv.Itoa(score.UserId) + "-" + score.Position + "-" + score.GroupID
		total := pkg.Sum(score.Performance, score.Project, score.Framework, score.Emphasize, score.Idea, score.Style, score.Colour, score.Animation, score.Writing)
		score.Total = total
		if e := d.DB.Create(&score).Error; e != nil {
			return e
		}
	}
	return nil
}
