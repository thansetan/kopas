package model

type Paste struct {
	Title     []byte
	Content   []byte
	ExpiresAt int64
}
