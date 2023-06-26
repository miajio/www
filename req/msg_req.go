package req

type MsgLeaveRequest struct {
	Mobile string `form:"mobile" uri:"mobile" json:"mobile" binding:"required,len=11"` // 手机号
	Name   string `form:"name" uri:"name" json:"name" binding:"required,max=32"`       // 姓名
	Msg    string `form:"" uri:"" json:"" binding:"required,max=500"`                  // 信息
}
