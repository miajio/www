package req

type EmailSendCheckCodeRequest struct {
	Email     string `form:"email" uri:"email" json:"email" binding:"required,email"`       // 收件人地址
	EmailType string `form:"emailType" uri:"emailType" json:"emailType" binding:"required"` // 邮件类型
}
