package model

//Go系列：结构体标签  https://juejin.cn/post/7005465902804123679#heading-11

//目前：用户登录和注册和公用一个结构体
//binding tag 可用 但不好 最好专门写个tool函数处理前端输入的数据
type UserLogin struct {
	Id           int64  `xorm:"bigint notnull" json:"id"`
	UserAccount  string `xorm:"pk varchar(20)" json:"user_account" binding:"required,min=6,max=20"`
	UserPassWord string `xorm:"varchar(20) notnull" json:"user_pass_word" binding:"required,min=6,max=20"`
}
type UserInfo struct {
	Id       int64  `xorm:"bigint notnull" json:"id"`
	UserName string `xorm:"pk varchar(40)" json:"user_name" binding:"required,max=40"`
	UserBio  string `xorm:"varchar(255)" json:"user_bio" binding:"required,max=255"`
}
