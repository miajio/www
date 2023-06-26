package req

type ArticleFindRequest struct {
	Group string `form:"group" uri:"group" json:"group" `
	Page  int    `form:"page" uri:"page" json:"page" `
	Limit int    `form:"limit" uri:"limit" json:"limit" `
}

func (a *ArticleFindRequest) Default() {
	if a.Page <= 0 {
		a.Page = 1
	}
	if a.Limit < 10 {
		a.Limit = 10
	}
}
