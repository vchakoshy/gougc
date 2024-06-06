package pkg

import "os"

func GetEnv(key, fallback string) string {
	if v, exists := os.LookupEnv(key); exists {
		return v
	}

	return fallback
}
