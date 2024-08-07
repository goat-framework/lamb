package template

import (
    "os"
    "regexp"
    "strings"
    "errors"
)

// ParseFile reads and returns the content of a .lamb.html file
func ParseFile(filepath string) (string, error) {
    if !strings.HasSuffix(filepath, ".lamb.html") {
		return "", errors.New("unsupported file type: only .lamb.html files are allowed")
	}
    content, err := os.ReadFile(filepath)
    if err != nil {
        return "", err
    }
    return string(content), nil
}

// ParseLambSyntax converts lamb syntax to Go template syntax
func ParseLambSyntax(content string) string {
	// Replace variable interpolation
	content = regexp.MustCompile(`{{\s*(\w+)\s*}}`).ReplaceAllString(content, "{{ .$1 }}")

	// Replace if directives
	content = regexp.MustCompile(`@if\s+(\w+)`).ReplaceAllString(content, "{{ if .$1 }}")
	content = regexp.MustCompile(`@else`).ReplaceAllString(content, "{{ else }}")
	content = regexp.MustCompile(`@endif`).ReplaceAllString(content, "{{ end }}")

	// Replace for directives
	content = regexp.MustCompile(`@for\s+(\w+)\s+in\s+(\w+)`).ReplaceAllString(content, "{{ range .$2 }}")
	content = regexp.MustCompile(`@endfor`).ReplaceAllString(content, "{{ end }}")

	// Replace generic @end
	content = regexp.MustCompile(`@end`).ReplaceAllString(content, "{{ end }}")

	return content
}
