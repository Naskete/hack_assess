package http

import (
	"HackAssess/internal/dao"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ncuhome/go-common/format"
	"io/ioutil"
	"log"
	"net/http"
)

const target = "https://api-usv2.ncuos.com/api/user/login"

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type response struct {
	UserID    int    `json:"user_id"`    // 拒绝重复评分
	StudentID string `json:"student_id"` // 判断评委
	Token     string `json:"token"`      // 保存登录状态
}

type auth struct {
	Token string `json:"token"`
}

func login(c *gin.Context) (*response, error) {
	// TODO 添加评委判断
	result, err := usLogin(c)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func giveMarks(c *gin.Context) (map[string]interface{}, error) {
	// TODO 添加登录验证
	if e := dao.Mark(c); e != nil {
		return format.FormatErrReturn(e, "Failed to mark")
	}
	return format.FormatNormalReturn(nil,"successful")
}

func usLogin(c *gin.Context) (*response, error) {
	var usr user
	//if e := c.ShouldBind(&usr); e != nil {
	//	log.Println("username or password can not be null")
	//	return nil, e
	//}
	_ = c.ShouldBind(&usr)
	data, _ := json.Marshal(usr)
	body := bytes.NewBuffer(data)
	req, err := http.Post(target, "application/json;charset=utf-8", body)
	if err != nil {
		log.Println("error")
		return nil, err
	}
	if req.StatusCode==403{
		return nil, errors.New("密码错误")
	}
	result, err := ioutil.ReadAll(req.Body)
	var res *response
	_ = json.Unmarshal(result, &res)
	return res, nil
}
