package sqlutils

import (
	"fmt"
	"os"
)

func CheckSQLFilesPath(filePath string) error {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return fmt.Errorf("SQL files directory does not exist at path: %s", filePath)
	} else if err != nil {
		return fmt.Errorf("error checking SQL files directory: %v", err)
	}

	files, err := os.ReadDir(filePath)
	if err != nil {
		return fmt.Errorf("error reading directory: %v", err)
	}

	if len(files) == 0 {
		return fmt.Errorf("no SQL files found in directory: %s", filePath)
	}

	return nil
}
