package template

import (
    "os"
)

// ParseFile reads and returns the content of a .lamb.html file
func ParseFile(filepath string) (string, error) {
    content, err := os.ReadFile(filepath)
    if err != nil {
        return "", err
    }
    return string(content), nil
}
