package dao

import (
	"HackAssess/internal/model"
	"HackAssess/pkg"
	"github.com/gin-gonic/gin"
	"strconv"
)

const (
	sql1 = "SELECT AVG(total) AS score FROM directors WHERE group_id = ?"
	sql2 = "SELECT AVG(total) AS score FROM products WHERE group_id = ?"
	sql3 = "SELECT AVG(total) AS score FROM designs WHERE group_id = ?"
	sql4 = "SELECT AVG(total) AS score FROM fronts WHERE group_id = ?"
	sql5 = "SELECT AVG(total) AS score FROM backs WHERE group_id = ?"
	sql6 = "SELECT AVG(total) AS score FROM shows WHERE group_id = ?"
)

type score struct {
	Score float32 `json:"score"`
}

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

func GetExcelData() []float32{
	var sql []string
	var result []float32
	sql = append(sql, sql1,sql2,sql3,sql4,sql5,sql6)
	// i 组
	for i:=1; i <= 7; i++{
		// 第 j 条sql语句
		for j:=0; j < 6; j++{
			var mark score
			d.DB.Raw(sql[j],i).Scan(&mark)
			result = append(result, mark.Score)
		}
	}
	return result
}