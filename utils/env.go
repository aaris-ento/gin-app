package utils

import "os"

func GetEnv(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if ok {
		return val
	}
	
	return fallback
}