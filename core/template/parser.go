package template

import (
	"strings"
    "os"
    "fmt"
    "path/filepath"
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

func CompileLamb(filePath string) error {
	// Parse the file to get the content
	parsedContent, err := ParseLamb(filePath)
	if err != nil {
		return err
	}

	// Define the .cache folder path
	cacheDir := ".cache"

	// Create the .cache directory if it doesn't exist
	if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
		err = os.Mkdir(cacheDir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create .cache directory: %w", err)
		}
	}

	// Remove the ".lamb" from the filename and replace it with ".html"
	outputFileName := strings.TrimSuffix(filepath.Base(filePath), ".lamb.html") + ".html"

	// Define the path for the compiled .html file
	outputFilePath := filepath.Join(cacheDir, outputFileName)

	// Write the parsed content to the .html file in the .cache directory
	err = os.WriteFile(outputFilePath, []byte(parsedContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to write compiled file: %w", err)
	}

	fmt.Printf("Compiled file written to %s\n", outputFilePath)
	return nil
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

	componentContent = applyAttributesDirective(componentContent, *component.Attributes)

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

	componentContent = applyAttributesDirective(componentContent, *component.Attributes)

	// Replace the <ui-slot /> placeholder with the wrapped content
	updatedContent := strings.Replace(componentContent, "<slot />", component.InnerContent, 1)
	updatedContent = strings.Replace(content, component.Element, updatedContent, 1)
	return updatedContent, nil
}
