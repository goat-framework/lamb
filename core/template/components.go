package template

import (
	"fmt"
	"regexp"
)

// Represents a self closing component
//
// Fields:
// - ComponentName (string): component name
// ex: link
// - ComponentFilePath (string): component filepath
// ex: components/link.lamb.html
// - Element (string): base element
// ex: <ui-link />
//
// Since: 0.1.0
type SelfClosingUIComponent struct {
	ComponentName     string
	ComponentFilePath string
	Element           string
	Attributes        *Attributes
}

// Represents a wrapped component
//
// Fields:
// - ComponentName (string): component name
// ex: button
// - ComponentFilePath (string): component filepath
// ex: components/button.lamb.html
// - InnerContent (string): component inner content
// ex: Submit
// - Element (string): base element
// ex: <ui-button>Submit</ui-button>
//
// Since: 0.1.0
type WrappedUIComponent struct {
	ComponentName     string
	ComponentFilePath string
	InnerContent      string
	Element           string
	Attributes        *Attributes
}

// Gets all of the self closing components
//
// Params:
// - content (string): content to parse
//
// Returns:
// []SelfClosingUIComponents: list of self closing components
//
// Since: 0.1.0
func getSelfClosingUIComponents(content string, componentDir string) []SelfClosingUIComponent {
	elements := findSelfClosingUIElements(content)
	names := createUIComponentNames(elements)
	paths := createUIComponentFilePaths(componentDir, names)

	var structs []SelfClosingUIComponent

	for i := range elements {
		attributes := getAttributes(elements[i])
		structs = append(structs, SelfClosingUIComponent{
			ComponentName:     names[i],
			ComponentFilePath: paths[i],
			Element:           elements[i],
			Attributes:        &attributes,
		})
	}

	return structs
}

// Gets all of the wrapped components
//
// Params:
// - content (string): content to parse
//
// Returns:
// - []WrappedUIComponent: list of wrapped components
//
// Since: 0.1.0
func getWrappedUIComponents(content string, componentDir string) []WrappedUIComponent {
	elements := findWrappedUIElements(content)
	names := createUIComponentNames(elements)
	contents := getComponentInnerContent(elements)
	paths := createUIComponentFilePaths(componentDir, names)

	var structs []WrappedUIComponent

	for i := range elements {
		attributes := getAttributes(elements[i])
		structs = append(structs, WrappedUIComponent{
			ComponentName:     names[i],
			ComponentFilePath: paths[i],
			InnerContent:      contents[i],
			Element:           elements[i],
			Attributes:        &attributes,
		})
	}

	return structs
}

// Looks for self closing element syntax
//
// Params:
// - content (string): content to parse
//
// Returns:
// - []string: list of elements
//
// Since: 0.1.0
func findSelfClosingUIElements(content string) []string {
	regex := regexp.MustCompile(`<ui-[\w-]+\s*[^>]*\/>`)
	matches := regex.FindAllString(content, -1)
	return matches
}

// Looks for wrapped element syntax
//
// Params:
// - content (string): content to parse
//
// Returns:
// - []string: list of elements
//
// Since: 0.1.0
func findWrappedUIElements(content string) []string {
	regex := regexp.MustCompile(`<ui-[\w-]+\b[^/>]*>.*?</ui-[\w-]+>`)
	matches := regex.FindAllString(content, -1)
	return matches
}

// Grabs component's inner content
// Must be a wrapped component
//
// Params:
// - elements ([]string): list of elements
//
// Returns:
// []string: list of inner contents
//
// Since: 0.1.0
func getComponentInnerContent(elements []string) []string {
	var contents []string
	regex := regexp.MustCompile(`<ui-[\w-]+[^>]*>(.*?)</ui-[\w-]+>`)

	for _, element := range elements {
		match := regex.FindStringSubmatch(element)
		if len(match) > 1 {
			contents = append(contents, match[1])
		} else {
			contents = append(contents, "")
		}
	}

	return contents
}

// Grabs component names
//
// Params:
// - elements ([]string): list of elements
//
// Returns:
// []string: list of names
//
// Since: 0.1.0
func createUIComponentNames(elements []string) []string {
	var names []string
	regex := regexp.MustCompile(`<ui-([\w-]+)`)

	for _, element := range elements {
		match := regex.FindStringSubmatch(element)
		if len(match) > 1 {
			names = append(names, match[1])
		}
	}

	return names
}

// Creates the component file paths
//
// Params:
// - names ([]string): list of component names
//
// Returns:
// - []string: list of filepaths to components
//
// Since: 0.1.0
func createUIComponentFilePaths(baseDir string, names []string) []string {
	var paths []string
	for _, name := range names {
		path := fmt.Sprintf("%s/%s.lamb.html", baseDir, name)
		paths = append(paths, path)
	}
	return paths
}
