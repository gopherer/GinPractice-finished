package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type User struct {
	UserId   int64  `xorm:"pk autoincr"`
	UserName string `xorm:"varchar(255)"`
}
type SmsCode struct {
	Id         int64  `xorm:"pk autoincr" json:"id"`
	Phone      string `xorm:"varchar(11)" json:"phone"`
	BizId      string `xorm:"varchar(30)" json:"biz_id"`
	Code       string `xorm:"varchar(6)" json:"code"`
	CreateTime int64  `xorm:"bigint" json:"create_time"`
}
type UU struct {
	Id           int64  `xorm:"pk autoincr" json:"id" `
	UserAccount  int64  `xorm:"bigint" json:"user_account"`
	UserPassWord string `xorm:"varchar(20)" json:"user_pass_word" binding:"required,min=6,max=20"`
	UserName     string `xorm:"varchar(40)" json:"user_name" binding:"required,max=40"`
	UserBio      string `xorm:"varchar(255)" json:"user_bio"`
}

func main() {
	conStr := "root:root@tcp(127.0.0.1:3306)/goWeb?charset="
	engine, err := xorm.NewEngine("mysql", conStr)
	fmt.Println(engine, err)
	//engine.SetMapper(core.SnakeMapper{})
	_ = engine.Sync2(new(SmsCode), new(UU))
	fmt.Println(1111111111)
}
