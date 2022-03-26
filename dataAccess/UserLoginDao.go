package dataAccess

import (
	"goWeb/model"
	"goWeb/tools"
)

type UserLoginDao struct {
	engine *tools.Orm
}

func (userLoginDao *UserLoginDao) GetUserLogin(userLogin *model.UserLogin) (bool, error) {
	userLoginDao.engine = tools.DbEngine
	return userLoginDao.engine.SQL("select id, user_account, user_pass_word from user_register where user_account = ?", userLogin.UserAccount).Get(userLogin)
}
