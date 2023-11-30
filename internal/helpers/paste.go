package helpers

import (
	"fmt"
	"strings"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

func GetRemainingTime(expiresAt int64) string {
	var expString strings.Builder
	if expiresAt == 0 {
		return "never"
	}

	expiredDate := time.Unix(expiresAt, 0)

	timeDiff := time.Until(expiredDate)

	minutes := int(timeDiff.Minutes()) % 60

	hours := int(timeDiff.Hours()) % 24

	days := int(timeDiff.Hours()) / 24

	if days > 0 {
		expString.WriteString(fmt.Sprintf("%d days ", days))
	}

	if hours > 0 {
		expString.WriteString(fmt.Sprintf("%d hours ", hours))
	}

	if minutes > 0 {
		expString.WriteString(fmt.Sprintf("%d minutes", minutes))
	}

	return expString.String()
}

func GenerateID() (string, error) {
	id, err := gonanoid.New(8)
	if err != nil {
		return "", err
	}
	return id, nil
}

func IsValidSize(content []byte) bool {
	return len(content) <= 20*1024*1024 // size can't be more than 20MB
}
