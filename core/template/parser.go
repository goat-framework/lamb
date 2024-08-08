package template

import (
	"strings"
)

// Parse the lamb file
//
// Params:
// - filepath (string): Path to the lamb file
//
// Returns:
// - string: The parsed content
// - error: if something goes wrong
//
// Since: 0.1.0
func ParseLamb(filepath string) (string, error) {
	content, err := getContent(filepath)
	if err != nil {
		return "", err
	}

	content = replaceSyntax(content)
	closingComponents := getSelfClosingUIComponents(content)
	for _, closingComponent := range closingComponents {
		content, err = replaceSelfClosingComponents(&closingComponent, content)
		if err != nil {
			return "", err
		}
	}
	wrappedComponents := getWrappedUIComponents(content)
	for _, wrappedComponent := range wrappedComponents {
		content, err = replaceWrappedComponents(&wrappedComponent, content)
		if err != nil {
			return "", err
		}
	}

	return content, nil
}

// Replace self closing component syntax
// with the component file
//
// Params:
// - component (*SelfClosingUIComponent): self closing component struct
// - content (string): content to parse
//
// Returns:
// - string: parsed content
// - error: if something goes wrong
//
// Since: 0.1.0
func replaceSelfClosingComponents(component *SelfClosingUIComponent, content string) (string, error) {
	componentContent, err := ParseLamb(component.ComponentFilePath)
	if err != nil {
		return "", err
	}

	updatedContent := strings.Replace(content, component.Element, componentContent, 1)
	return updatedContent, nil
}

// Replace wrapped component syntax
// with the component file
//
// Params:
// - component (*WrappedUIComponent): wrapped component struct
// - content (string): content to parse
//
// Returns:
// - string: parsed content
// - error: if something goes wrong
//
// Since: 0.1.0
func replaceWrappedComponents(component *WrappedUIComponent, content string) (string, error) {
	componentContent, err := ParseLamb(component.ComponentFilePath)
	if err != nil {
		return "", err
	}

	// Replace the <ui-slot /> placeholder with the wrapped content
	updatedContent := strings.Replace(componentContent, "<slot />", component.InnerContent, 1)
	updatedContent = strings.Replace(content, component.Element, updatedContent, 1)
	return updatedContent, nil
}
