package pastedto

type Paste struct {
	Title   string `form:"title" json:"title" binding:"required"`
	Content string `form:"content" json:"content" binding:"required"`
}
