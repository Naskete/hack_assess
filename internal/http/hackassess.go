package http

import (
	"HackAssess/internal/dao"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ncuhome/go-common/format"
	account "github.com/ncuhome/us-backend/account/pkg"
	"github.com/ncuhome/us-backend/feedback/pkg"
	"io/ioutil"
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

func login(c *gin.Context) (map[string]interface{}, error) {
	// 添加评委判断
	result, err := usLogin(c)
	if err != nil {
		return format.FormatErrReturn(err)
	}
	grade := account.JudgeGrade(result.StudentID)
	if grade < 2 || result.StudentID == "6109119101" || result.StudentID == "6109119121" {
		return format.FormatErrReturn(errors.New("sorry, Contestants could not grade"))
	}
	return format.FormatNormalReturn(result, "success")
}

func giveMarks(c *gin.Context) (map[string]interface{}, error) {
	// 添加登录验证
	token := pkg.GetToken(c)
	// 验证token合法性
	_, err := account.ParseToken(token)
	if token == "" || err != nil {
		return format.FormatErrReturn(errors.New("please login"))
	}
	if e := dao.Mark(c); e != nil {
		return format.FormatErrReturn(e, "please do not grade again")
	}
	return format.FormatNormalReturn(nil, "successful")
}

func usLogin(c *gin.Context) (*response, error) {
	var usr user
	_ = c.ShouldBind(&usr)
	data, _ := json.Marshal(usr)
	body := bytes.NewBuffer(data)
	req, err := http.Post(target, "application/json;charset=utf-8", body)
	if err != nil {
		return nil, err
	}
	if req.StatusCode == 403 {
		return nil, errors.New("wrong password")
	}
	result, err := ioutil.ReadAll(req.Body)
	var res *response
	_ = json.Unmarshal(result, &res)
	return res, nil
}
