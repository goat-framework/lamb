package integration_test

import (
	"testing"

	"github.com/goat-framework/lamb/core/template"
)

// Path to the example lamb file
var filepath string = "./example.lamb.html"

func TestCompileLamb(t *testing.T) {
	err := template.Compile(filepath, "components")
	if err != nil {
		t.Errorf("Expected no error, but got error: %s", err.Error())
	}
}
