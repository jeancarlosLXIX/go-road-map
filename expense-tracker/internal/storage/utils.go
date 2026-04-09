package storage

import "os"

func fileExist(filePath string) error {
	_, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		// create the file if not
		return os.WriteFile(filePath, []byte("[]"), 0644)
	}

	return err
}
