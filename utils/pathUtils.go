package utils

import (
	"os"
)

func RootPath() string {
	current, _ := os.Getwd()
	return current
}
