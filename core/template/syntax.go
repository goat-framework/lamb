package template

import (
	"regexp"
)

// Convert lamb syntax to standard go template syntax
//
// Params:
// - content (string): content to parse
//
// Returns:
// - string: the parsed content
//
// Since 0.1.0
func replaceSyntax(content string) string {
	// Replace variable interpolation
	content = regexp.MustCompile(`{{\s*(\w+)\s*}}`).ReplaceAllString(content, "{{ .$1 }}")

	// Replace if directives
	content = regexp.MustCompile(`@if\s+(\w+)`).ReplaceAllString(content, "{{ if .$1 }}")
	content = regexp.MustCompile(`@elseif\s+(\w+)`).ReplaceAllString(content, "{{ else if .$1 }}")
	content = regexp.MustCompile(`@else`).ReplaceAllString(content, "{{ else }}")

	// Replace for directives
	content = regexp.MustCompile(`@for\s+(\w+)\s+in\s+(\w+)`).ReplaceAllString(content, "{{ range .$2 }}")

	// Replace generic @end
	content = regexp.MustCompile(`@end`).ReplaceAllString(content, "{{ end }}")

	return content
}
