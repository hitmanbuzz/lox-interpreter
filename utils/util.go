package utils

import (
	"errors"
	"os"
)

func IsFileExist(fname string) bool {
	_, err := os.Stat(fname)
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}
