package template

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// ParseFile reads and returns the content of a .lamb.html file
func CheckFile(filepath string) (string, error) {
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

	// Replace for directives
	content = regexp.MustCompile(`@for\s+(\w+)\s+in\s+(\w+)`).ReplaceAllString(content, "{{ range .$2 }}")

	// Replace generic @end
	content = regexp.MustCompile(`@end`).ReplaceAllString(content, "{{ end }}")

	return content
}

func ReplaceComponents(content string) (string, error) {
	componentRegex := regexp.MustCompile(`<ui-([\w-]+)\s*/>`)
	matches := componentRegex.FindAllStringSubmatch(content, -1)

	for _, match := range matches {
		componentName := match[1]
		componentFile := fmt.Sprintf("components/%s.lamb.html", componentName)
		componentContent, err := ParseFile(componentFile)
		if err != nil {
			return "", err
		}
		content = strings.Replace(content, match[0], componentContent, 1)
	}

	// Handle wrapped components
	wrappedComponentRegex := regexp.MustCompile(`<ui-([\w-]+)>(.*?)</ui-[\w-]+>`)
	matches = wrappedComponentRegex.FindAllStringSubmatch(content, -1)

	for _, match := range matches {
		componentName := match[1]
		innerContent := match[2]
		componentFile := fmt.Sprintf("components/%s.lamb.html", componentName)
		wrapperContent, err := ParseFile(componentFile)
		if err != nil {
			return "", err
		}
		// Replace the <ui-slot /> placeholder with the wrapped content
		wrapperContent = strings.Replace(wrapperContent, "<ui-slot />", innerContent, 1)
		content = strings.Replace(content, match[0], wrapperContent, 1)
	}

	return content, nil
}

// Parses the .lamb.html file and converts its syntax to Go template syntax
func ParseFile(filepath string) (string, error) {
	content, err := CheckFile(filepath)
	if err != nil {
		return "", err
	}
	parsedContent := ParseLambSyntax(content)

	return parsedContent, nil
}
