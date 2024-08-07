package template_test

import (
    "testing"

    "github.com/goat-framework/lamb/core/template"
)

func TestParseLambFile(t *testing.T) {
    // Path to the .lamb.html file for testing
    filepath := "./example.lamb.html"

    // Attempt to parse the .lamb.html file
    _, err := template.ParseFile(filepath)
    if err != nil {
        t.Fatalf("Error parsing .lamb.html file: %v", err)
    }
}
