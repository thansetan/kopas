package pastedto

type PasteReq struct {
	Title     string `form:"title" json:"title" binding:"required"`
	Content   string `form:"content" json:"content" binding:"required"`
	ExpiresAt string `form:"expires_at" json:"expires_at" binding:"required"`
}

type PasteResp struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	ExpiresAt int64  `json:"expires_at"`
}
