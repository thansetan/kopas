package helpers

import "os"

func GetEnvOrDefault(key, def string) string {
	val, exists := os.LookupEnv(key)
	if !exists {
		return def
	}
	return val
}
