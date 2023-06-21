package req

// UserEmailLoginRequest 用户邮箱登录请求
type UserEmailLoginRequest struct {
	Email    string `form:"email" uri:"email" json:"email" binding:"required,email"`                 // 邮箱地址
	Password string `form:"password" uri:"password" json:"password" binding:"required,min=6,max=32"` // 密码
}

// UserEmailRegisterRequest 用户邮箱注册请求
type UserEmailRegisterRequest struct {
	Uid       string `form:"uid" uri:"uid" json:"uid" binding:"required"`                             // 邮箱验证码的uid
	Username  string `form:"username" uri:"username" json:"username" binding:"required,min=2,max=32"` // 用户名
	Email     string `form:"email" uri:"email" json:"email" binding:"required,email"`                 // 邮箱地址
	CheckCode string `form:"checkCode" uri:"checkCode" json:"checkCode" binding:"required,len=6"`     // 验证码
	Password  string `form:"password" uri:"password" json:"password" binding:"required,min=6,max=32"` // 密码
}
