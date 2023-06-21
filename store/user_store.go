package store

import (
	"database/sql"
	"errors"

	"github.com/miajio/www/lib"
	"github.com/miajio/www/model"
)

type userInfoStoreImpl struct{}

type userInfoStore interface {
	Login(email, password string) (model.UserInfoModel, error)               // Login 登录操作
	EmailRepeat(email string) (bool, error)                                  // EmailRepeat 判断邮箱是否存在
	Register(username, email, password string) (*model.UserInfoModel, error) // Register 注册操作
}

var UserInfoStore userInfoStore = (*userInfoStoreImpl)(nil)

// Login 登录操作
func (*userInfoStoreImpl) Login(email, password string) (model.UserInfoModel, error) {
	var result model.UserInfoModel
	searchSQL := "SELECT * FROM `user_info` WHERE `email` = ? AND `password` = MD5(?) AND `status` = 1"
	err := lib.DB.Get(&result, searchSQL, email, password)
	if err != nil && err == sql.ErrNoRows {
		err = errors.New("用户名或密码错误")
	}

	return result, err
}

// EmailRepeat 判断邮箱是否存在
func (*userInfoStoreImpl) EmailRepeat(email string) (bool, error) {
	var count int
	searchSql := "SELECT COUNT(1) FROM `user_info` WHERE `email` = ? AND `status` != 3"
	err := lib.DB.Get(&count, searchSql, email)
	return count > 0, err
}

// Register 注册操作
func (*userInfoStoreImpl) Register(username, email, password string) (*model.UserInfoModel, error) {
	uid := lib.UID()

	insertSql := "INSERT INTO `user_info` (`uid`, `username`, `email`, `password`, `status`, `create_time`, `update_time`) VALUES (?, ?, ?, MD5(?), 1, NOW(), NOW())"
	_, err := lib.DB.Exec(insertSql, uid, username, "nil", email, password)
	if err != nil {
		return nil, err
	}

	var result model.UserInfoModel
	err = lib.DB.Get(&result, "SELECT * FROM `user_info` WHERE `uid` = ?", uid)
	return &result, err
}
