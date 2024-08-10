package template

import (
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestCheckFileFail(t *testing.T) {
	filepath := "test.html"

	expected := "unsupported file type: only .lamb.html files are allowed"

	_, err := checkFile(filepath)
	if err == nil {
		t.Errorf("Expected an error, but got none")
	} else if err.Error() != expected {
		t.Errorf("Expected error message '%s', but got '%s'", expected, err.Error())
	}
}

func TestCheckFile(t *testing.T) {
	filepath := "test.lamb.html"

	result, err := checkFile(filepath)
	if err != nil {
		t.Errorf("Expected no error, but got error: %s", err.Error())
	}
	if result != filepath {
		t.Errorf("Expected %s, but got %s", filepath, result)
	}
}

func TestGetContent(t *testing.T) {
	filepath := "../../tests/example.lamb.html"

	expected := "<h1>{{ title }}</h1>"

	result, err := getContent(filepath)
	if err != nil {
		t.Errorf("Expected no error, but got error: %s", err.Error())
		return
	}

	if !strings.Contains(result, expected) {
		t.Errorf("Expected content to contain %s, but got %s", expected, result)
	}
}

func TestGetLibraryRoot(t *testing.T) {
	_, b, _, _ := runtime.Caller(0)
	expectedRootDir := filepath.Join(filepath.Dir(b), "../..")

	actualRootDir := getLibraryRoot()

	if filepath.Clean(actualRootDir) != filepath.Clean(expectedRootDir) {
		t.Errorf("expected root directory %s, but got %s", expectedRootDir, actualRootDir)
	}
}
