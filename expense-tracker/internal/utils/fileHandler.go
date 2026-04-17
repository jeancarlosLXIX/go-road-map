package utils

import (
	// "fmt"
	"os"
)

func FileExist(filePath string) error {
	_, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		// create the file
		return os.WriteFile(filePath, []byte("[]"), 0644)
	}

	return err
}
