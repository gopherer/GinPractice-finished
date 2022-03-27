package dataAccess

import (
	"goWeb/model"
	"goWeb/tools"
)

type UserLoginDao struct {
	engine *tools.Orm
}

func (userLoginDao *UserLoginDao) GetUserLogin(userLogin *model.UserLogin) bool {
	userLoginDao.engine = tools.DbEngine
	//无论engine.SQL结果的对错 err结构都为nil
	result, _ := userLoginDao.engine.SQL("select id, user_account, user_pass_word from user_register where user_account = ? and user_pass_word = ?", userLogin.UserAccount, userLogin.UserPassWord).Get(userLogin)
	if !result {
		result, _ = userLoginDao.engine.SQL("select id, user_account, user_pass_word from user_register where user_account = ?", userLogin.UserAccount).Get(userLogin)
		if result {
			userLogin.UserPassWord = ""
		}
		return false
	}
	return true
}
