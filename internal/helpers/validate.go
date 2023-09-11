package helpers

func ValidSize(content []byte) bool {
	return len(content) < 20*1024*1024 // size can't be more than 20MB
}
