package http

import (
	"HackAssess/internal/dao"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"github.com/ncuhome/go-common/format"
	account "github.com/ncuhome/us-backend/account/pkg"
	"github.com/ncuhome/us-backend/feedback/pkg"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// TODO 优化
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
		// TODO  区别优化？
		return nil, errors.New("wrong password or user not exists")
	}
	result, err := ioutil.ReadAll(req.Body)
	var res *response
	_ = json.Unmarshal(result, &res)
	return res, nil
}

func getExcel(c *gin.Context) (map[string]interface{},error){
	data := dao.GetExcelData()
	if data == nil{
		return format.FormatErrReturn(errors.New("failed to export excel, the data is null"))
	}
	filename := fmt.Sprint("hackweek评分结果表.xlsx")
	fileName := url.QueryEscape(filename)
	c.Writer.Header().Set("Content-type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	createExcel(data, c.Writer)
	return format.FormatNormalReturn(nil, "success")
}

func createExcel(data []float32, w http.ResponseWriter) {
	file := excelize.NewFile()
	//// 表头
	file.SetSheetRow("Sheet1", "A1", &[]interface{}{"组别", "运营", "产品", "设计", "前端","后端","路演","总分"})
	for i := 0; i < 6; i++{
		// 获取每一组的成绩
		d := data[i*6:i*6+6]
		res := d[0]*5/28 + d[1]*5/28 + d[2]*5/28 + d[3]/7 + d[4]/7 +d[5]*5/28
		rowIndex := strconv.Itoa(i+2)
		location := "A"+rowIndex
		file.SetSheetRow("Sheet1", location, &[]interface{}{i+1,d[0],d[1],d[2],d[3],d[4],d[5], fmt.Sprintf("%.2f", res)})
	}
	_ = file.Write(w)
}