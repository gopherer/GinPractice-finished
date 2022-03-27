package tools

import "goWeb/model"

func UserJsonLen(user interface{}) bool {
	if v, ok := user.(model.UserRegister); ok {
		if len(v.UserAccount) >= 6 && len(v.UserAccount) <= 20 && len(v.UserPassWord) >= 6 && len(v.UserPassWord) <= 20 {
			return true
		}
	} else if v, ok := user.(model.UserLogin); ok {
		if len(v.UserAccount) >= 6 && len(v.UserAccount) <= 20 && len(v.UserPassWord) >= 6 && len(v.UserPassWord) <= 20 {
			return true
		}
	} else if v, ok := user.(model.UserInfo); ok {
		if len(v.UserName) <= 40 && len(v.UserBio) <= 255 {
			return true
		}
	}
	return false
}
