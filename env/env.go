package env

import (
	"os"
	"strconv"
)

// Bool returns a bool from the ENV, or fallback variable
func Bool(key string, fallback bool) bool {
	switch os.Getenv(key) {
	case "true":
		return true
	case "false":
		return false
	default:
		return fallback
	}
}

// Int returns an int from the ENV, or fallback variable
func Int(key string, fallback int) int {
	if i, err := strconv.Atoi(os.Getenv(key)); err == nil {
		return i
	}

	return fallback
}

// String returns a string from the ENV, or fallback variable
func String(key, fallback string) string {
	v := os.Getenv(key)
	if v != "" {
		return v
	}
	return fallback
}
