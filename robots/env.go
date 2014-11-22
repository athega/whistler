package robots

import "os"

// EnvString returns a string from the ENV, or fallback variable
func EnvString(key, fallback string) string {
	v := os.Getenv(key)
	if v != "" {
		return v
	}
	return fallback
}
