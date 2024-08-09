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
func getSelfClosingUIComponents(content string) []SelfClosingUIComponent {
	components := findSelfClosingUIComponents(content)
	elements := getComponentElements(components)
	names := createUIComponentNames(components)
	paths := createUIComponentFilePaths(names)

	var structs []SelfClosingUIComponent

	for i := range components {
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
func getWrappedUIComponents(content string) []WrappedUIComponent {
	components := findWrappedUIComponents(content)
	elements := getComponentElements(components)
	names := createUIComponentNames(components)
	contents := getComponentInnerContent(components)
	paths := createUIComponentFilePaths(names)

	var structs []WrappedUIComponent

	for i := range components {
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

// Looks for self closing component syntax
//
// Params:
// - content (string): content to parse
//
// Returns:
// - [][]string: list of listed components
// ex: {{'<ui-link />', 'link',}, {'<ui-button />', 'button',},}
//
// Since: 0.1.0
func findSelfClosingUIComponents(content string) [][]string {
	regex := regexp.MustCompile(`<ui-([\w-]+)\s*/>`)
	matches := regex.FindAllStringSubmatch(content, -1)
	return matches
}

// Looks for wrapped component syntax
//
// Params:
// - content (string): content to parse
//
// Returns:
// - [][]string: list of listed components
// ex: {{'<ui-button>submit</ui-button>', 'button', 'submit',},}
//
// Since: 0.1.0
func findWrappedUIComponents(content string) [][]string {
	regex := regexp.MustCompile(`<ui-([\w-]+)>(.*?)</ui-[\w-]+>`)
	matches := regex.FindAllStringSubmatch(content, -1)
	return matches
}

// Grabs component's base element
//
// Params:
// - components ([][]string): the listed component
//
// Returns:
// []string: list of base elements
//
// Since 0.1.0
func getComponentElements(components [][]string) []string {
	var elements []string
	for _, component := range components {
		element := component[0]
		elements = append(elements, element)
	}

	return elements
}

// Grabs component's inner content
// Must be a wrapped component
//
// Params:
// - components ([][]string): the listed component
//
// Returns:
// []string: list of inner contents
//
// Since: 0.1.0
func getComponentInnerContent(components [][]string) []string {
	var contents []string
	for _, component := range components {
		content := component[2]
		contents = append(contents, content)
	}

	return contents
}

// Grabs component names
//
// Params:
// - components ([][]string): the listed component
//
// Returns:
// []string: list of names
//
// Since: 0.1.0
func createUIComponentNames(components [][]string) []string {
	var names []string
	for _, component := range components {
		name := component[1]
		names = append(names, name)
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
func createUIComponentFilePaths(names []string) []string {
	var paths []string
	for _, name := range names {
		path := fmt.Sprintf("components/%s.lamb.html", name)
		paths = append(paths, path)
	}
	return paths
}
