package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetOutputFileName(t *testing.T) {
	compiler := Compiler{
		ComponentDir: "views/components",
		FilePath:     "views/test.lamb.html",
	}

	expected := "test.html"
	result := compiler.getOutputFileName()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestGetOutputFilePath(t *testing.T) {
	compiler := Compiler{
		ComponentDir: "views/components",
		FilePath:     "views/test.lamb.html",
	}

	root := getLibraryRoot()
	cachePath := fmt.Sprintf("%s/.cache", root)
	expected := fmt.Sprintf("%s/test.html", cachePath)

	result := compiler.getOutputFilePath("test.html", cachePath)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
