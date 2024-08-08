package template

import (
	"errors"
	"os"
	"strings"
)

// Confirm filepath is to a lamb file
// Ex: example.lamb.html
//
// Params:
// - filepath (string): filepath to check
//
// Returns:
// - string: the original filepath
// - error: if path is not to a lamb file
//
// Since 0.1.0
func checkFile(filepath string) (string, error) {
	if !strings.HasSuffix(filepath, ".lamb.html") {
		return "", errors.New("unsupported file type: only .lamb.html files are allowed")
	}

	return filepath, nil
}

// Read the lamb file
//
// Params:
// - filepath (string): filepath to read
//
// Returns:
// - string: file contents
// - error: if something goes wrong
//
// Since 0.1.0
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
