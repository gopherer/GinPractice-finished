package model

//Go系列：结构体标签  https://juejin.cn/post/7005465902804123679#heading-11
//https://www.cnblogs.com/chnmig/p/11382390.html

//binding tag 可用 但不好 最好专门写个tool函数处理前端输入的数据
//使用匿名字段 在结构体映射为数据表时无法得到实际效果
//form tag 可接收html form表单 application/x-www-form-urlencoded的类容 可以以接受json文本的内容
//json tag 只可以接收json的文本内容
type UserRegister struct {
	Id           int64  `xorm:"bigint notnull" json:"id"`
	UserAccount  string `xorm:"pk varchar(20)" form:"user_account" json:"user_account" binding:"required,min=6,max=20"`
	UserPassWord string `xorm:"varchar(20) notnull" form:"user_pass_word" json:"user_pass_word" binding:"required,min=6,max=20"`
}
type UserInfo struct {
	Id       int64  `xorm:"bigint notnull" json:"id"`
	UserName string `xorm:"pk varchar(40)" form:"user_account"  json:"user_name" binding:"required,max=40"`
	UserBio  string `xorm:"varchar(255)" form:"user_bio" json:"user_bio" binding:"required,max=255"`
}
type UserLogin struct {
	Id           int64
	UserAccount  string `form:"user_account"  json:"user_account"`
	UserPassWord string `form:"user_pass_word" json:"user_pass_word"`
}
