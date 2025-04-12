package sqlutils

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
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

func LoadSQLFiles(path string) (map[string]string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read SQL directory: %w", err)
	}

	queries := make(map[string]string, len(files))

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		name := file.Name()
		if !strings.HasSuffix(name, ".sql") {
			continue
		}

		fullPath := filepath.Join(path, name)
		content, err := os.ReadFile(fullPath)
		if err != nil {
			return nil, fmt.Errorf("failed to read file %s: %w", fullPath, err)
		}

		queries[name] = string(bytes.Join(bytes.Fields(content), []byte(" ")))
	}

	if len(queries) == 0 {
		return nil, fmt.Errorf("no .sql files loaded from: %s", path)
	}

	return queries, nil
}
