package model

import "time"

// 用户头像信息结构体
type UserHeadInfo struct {
	Id         int        `db:"id" json:"id"`                  // 自增长id
	Uid        string     `db:"uid" json:"uid"`                // 头像uid
	HeadPath   string     `db:"head_path" json:"headPath"`     // 头像路径
	CreateTime *time.Time `db:"create_time" json:"createTime"` // 创建时间
	UpdateTime *time.Time `db:"update_time" json:"updateTime"` // 修改时间
}
