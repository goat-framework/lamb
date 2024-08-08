package template

import (
	"errors"
	"os"
	"strings"
)

func checkFile(filepath string) (string, error) {
	if !strings.HasSuffix(filepath, ".lamb.html") {
		return "", errors.New("unsupported file type: only .lamb.html files are allowed")
	}

	return filepath, nil
}

func getContent(filepath string) (string, error) {
	filepath, err := checkFile(filepath)
	if err != nil {
		return "", err
	}

	content, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
