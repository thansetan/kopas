package helpers

import (
	"errors"
	"path"
	"runtime"
)

func GetCurrentFileDir() (string, error) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return "", errors.New("error")
	}
	return path.Dir(filename), nil
}
