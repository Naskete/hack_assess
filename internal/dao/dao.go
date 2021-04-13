package dao

import (
	"HackAssess/internal/model"
	commonGorm "github.com/ncuhome/go-common/db/gorm"
	mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// 初始化数据库连接，重复代码

type Dao struct {
	DSN string
	DB *gorm.DB
}

var d *Dao

func DatabaseInit() {
	d = &Dao{
		DSN : getDsn(),
		DB: commonGorm.GetDB(),
	}
	d.initOrm()
	d.autoMigrate()
}

func (d *Dao) initOrm()  {
	commonGorm.Init(mysql.Open(d.DSN), &gorm.Config{})
	d.DB = commonGorm.GetDB()
}

func getDsn() string{
	values, err := env.NewReader().Values()
	if err != nil{
		log.Println(err)
	}
	return values.Get("HackAssess_DSN").String("")
}

func (d *Dao) autoMigrate(){
	_ = d.DB.AutoMigrate(&model.Score{})
}
