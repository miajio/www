package store

import (
	"github.com/miajio/www/lib"
	"github.com/miajio/www/model"
)

type userInfoStoreImpl struct{}

type userInfoStore interface {
	Login(email, password string) (model.UserInfoModel, error) // Login 登录操作
}

var UserInfoStore userInfoStore = (*userInfoStoreImpl)(nil)

// Login 登录操作
func (*userInfoStoreImpl) Login(email, password string) (model.UserInfoModel, error) {
	var result model.UserInfoModel
	searchSQL := "SELECT * FROM `user_info` WHERE `email` = ? AND `password` = MD5(?) AND `status` = 1"
	err := lib.DB.Get(&result, searchSQL, email, password)
	return result, err
}
