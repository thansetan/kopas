package helpers

import gonanoid "github.com/matoous/go-nanoid/v2"

func GenerateID() (string, error) {
	id, err := gonanoid.New(8)
	if err != nil {
		return "", err
	}
	return id, nil
}
