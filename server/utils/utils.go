package utils

import (
	"os"
)

// PathExists will check whether a path exists or not.
// This can be a file or a directory
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
