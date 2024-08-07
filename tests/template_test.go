package template_test

import (
    "testing"

    "github.com/goat-framework/lamb/core/template"
)

// Path to the example lamb file
var filepath string = "./example.lamb.html"

func TestOpenLambFile(t *testing.T) {
    // Attempt to parse the .lamb.html file
    _, err := template.ParseFile(filepath)
    if err != nil {
        t.Fatalf("Error parsing .lamb.html file: %v", err)
    }
}

func TestParseUnsupportedFile(t *testing.T) {
	// Path to the unsupported file for testing
	filepath := "tests/example.txt"

	// Attempt to parse the unsupported file type
	_, err := template.ParseFile(filepath)
	if err == nil {
		t.Fatalf("Expected error for unsupported file type, got nil")
	}

	expectedErr := "unsupported file type: only .lamb.html files are allowed"
	if err.Error() != expectedErr {
		t.Fatalf("Expected error message:\n%s\nGot:\n%s", expectedErr, err.Error())
	}
}

func TestConvertSyntax(t *testing.T) {
    content, err := template.ParseFile(filepath)
    if err != nil {
        t.Fatalf("Error parsing .lamb.html file: %v", err)
    }

    expected := `<h1>{{ .title }}</h1>
<p>{{ .content }}</p>

{{ if .isLoggedIn }}
    <p>Welcome, {{ .user }}</p>
{{ else }}
    <p>Please log in.</p>
{{ end }}
`

    // Convert the content using ParseLambSyntax
    result := template.ParseLambSyntax(string(content))
    if result != expected {
        t.Errorf("Expected:\n%s\nGot:\n%s", expected, result)
    }
}
