package http

import (
	"HackAssess/internal/dao"
	"github.com/gin-gonic/gin"
)

func GiveMarks(c *gin.Context) (string ,error){
	if e := dao.Mark(c); e!=nil{
		return "err", e
	}
	return "ok",nil
}