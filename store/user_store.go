package store

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/miajio/www/lib"
	"github.com/miajio/www/model"
)

type userInfoStoreImpl struct{}

type userInfoStore interface {
	FindUserInfoByUid(uid string) (model.UserInfoModel, error)              // FindUserInfoByUid 依据用户uid获取用户信息
	Login(email, password string) (model.UserInfoModel, error)              // Login 登录操作
	EmailRepeat(email string) (bool, error)                                 // EmailRepeat 判断邮箱是否存在
	Register(username, email, password string) (model.UserInfoModel, error) // Register 注册操作

	Update(uid, headPic, username string) (model.UserInfoModel, error) // Update 修改用户信息(修改头像、昵称)

	FindHeadPicPath(headPic string) (string, error) // FindHeadPicPath 查询头像地址

}

var UserInfoStore userInfoStore = (*userInfoStoreImpl)(nil)

// FindUserInfoByUid 依据用户uid获取用户信息
func (*userInfoStoreImpl) FindUserInfoByUid(uid string) (model.UserInfoModel, error) {
	var result model.UserInfoModel
	searchSQL := "SELECT * FROM `user_info` WHERE `uid` = ? AND `status` = 1"
	err := lib.DB.Get(&result, searchSQL, uid)
	return result, err
}

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
func (*userInfoStoreImpl) Register(username, email, password string) (model.UserInfoModel, error) {
	result := model.UserInfoModel{}
	uid := lib.UID()

	insertSql := "INSERT INTO `user_info` (`uid`, `username`, `email`, `password`, `status`, `create_time`, `update_time`) VALUES (?, ?, ?, MD5(?), 1, NOW(), NOW())"
	_, err := lib.DB.Exec(insertSql, uid, username, email, password)
	if err != nil {
		return result, err
	}

	err = lib.DB.Get(&result, "SELECT * FROM `user_info` WHERE `uid` = ?", uid)
	return result, err
}

// Update 修改用户信息(修改头像、昵称)
func (u *userInfoStoreImpl) Update(uid, headPic, username string) (model.UserInfoModel, error) {
	userInfo, err := u.FindUserInfoByUid(uid)
	if err != nil {
		return userInfo, err
	}
	oldHeadPic, _ := u.FindHeadPicPath(userInfo.HeadPic)
	tx := lib.DB.MustBegin()

	if oldHeadPic == "" {
		userInfo.HeadPic = lib.UID()
		insertSql := "INSERT INFO `user_head_info` (`uid`, `head_path`, `create_time`, `update_time`) (?, ?, NOW(), NOW())"
		tx.MustExec(insertSql, userInfo.HeadPic, headPic)
	} else {
		updateUserHeadSql := "UPDATE `user_head_info` SET `head_path` = ?, update_time = NOW() WHERE `uid` = ?"
		tx.MustExec(updateUserHeadSql, headPic, userInfo.HeadPic)
	}
	updateSql := "UPDATE `user_info` SET `username` = ?, `head_pic` = ?, `update_time` = NOW() WHERE `uid` = ?"
	tx.MustExec(updateSql, username, userInfo.HeadPic, uid)

	if err := tx.Commit(); err != nil {
		return userInfo, tx.Rollback()
	}
	return userInfo, nil
}

// FindHeadPicPath 查询头像地址
func (*userInfoStoreImpl) FindHeadPicPath(headPic string) (string, error) {
	if strings.TrimSpace(headPic) == "" {
		return "", nil
	}
	var headPath string
	searchSql := "SELECT `head_path` FROM `user_head_info` WHERE uid = ?"
	err := lib.DB.Get(&headPath, searchSql, headPic)
	return headPath, err
}
