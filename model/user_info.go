package model

import "time"

// 用户信息结构体
type UserInfoModel struct {
	Id         int        `db:"id" json:"id"`                  // 自增长id
	Uid        string     `db:"uid" json:"uid"`                // 系统自动生成用户UUID
	Username   string     `db:"username" json:"username"`      // 用户名
	HeadPic    *string    `db:"head_pic" json:"headPic"`       // 头像id
	Email      string     `db:"email" json:"email"`            // 邮箱
	Password   string     `db:"password" json:"password"`      // 密码
	Status     int        `db:"status" json:"status"`          // 状态 0 失效 1 正常 2 停用 3 删除
	CreateTime *time.Time `db:"create_time" json:"createTime"` // 创建时间
	UpdateTime *time.Time `db:"update_time" json:"updateTime"` // 修改时间
}
