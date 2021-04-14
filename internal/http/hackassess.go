package http

import (
	"HackAssess/internal/dao"
	grpcPython "HackAssess/internal/grpc"
	"HackAssess/proto"
	"github.com/gin-gonic/gin"
	"log"
)

func login(c *gin.Context) (*proto.Auth, error) {
	//var user model.User
	//if e := c.ShouldBind(&user); e!=nil{
	//	return nil, e
	//}
	result, err := grpcPython.LoginClientInit()
	if err != nil {
		log.Printf("get result err : %v", err)
		return nil, err
	}
	return result, nil
}

func giveMarks(c *gin.Context) (string, error) {
	if e := dao.Mark(c); e != nil {
		return "err", e
	}
	return "ok", nil
}
