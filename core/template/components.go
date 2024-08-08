package template

import (
	"fmt"
	"regexp"
)

type SelfClosingUIComponent struct {
	ComponentName     string
	ComponentFilePath string
}

type WrappedUIComponent struct {
	ComponentName     string
	ComponentFilePath string
	InnerContent      string
}

func GetSelfClosingUIComponents(content string) []SelfClosingUIComponent {
	components := findSelfClosingUIComponents(content)
	names := createUIComponentNames(components)
	paths := createUIComponentFilePaths(names)

	var structs []SelfClosingUIComponent

	for i := range components {
		structs = append(structs, SelfClosingUIComponent{
			ComponentName:     names[i],
			ComponentFilePath: paths[i],
		})
	}

	return structs
}

func GetWrappedUIComponents(content string) []WrappedUIComponent {
	components := findWrappedUIComponents(content)
	names := createUIComponentNames(components)
	contents := getComponentInnerContent(components)
	paths := createUIComponentFilePaths(names)

	var structs []WrappedUIComponent

	for i := range components {
		structs = append(structs, WrappedUIComponent{
			ComponentName:     names[i],
			ComponentFilePath: paths[i],
			InnerContent:      contents[i],
		})
	}

	return structs
}

func findSelfClosingUIComponents(content string) [][]string {
	regex := regexp.MustCompile(`<ui-([\w-]+)\s*/>`)
	matches := regex.FindAllStringSubmatch(content, -1)
	return matches
}

func findWrappedUIComponents(content string) [][]string {
	regex := regexp.MustCompile(`<ui-([\w-]+)>(.*?)</ui-[\w-]+>`)
	matches := regex.FindAllStringSubmatch(content, -1)
	return matches
}

func getComponentInnerContent(components [][]string) []string {
	var contents []string
	for _, component := range components {
		content := component[2]
		contents = append(contents, content)
	}

	return contents
}

func createUIComponentNames(components [][]string) []string {
	var names []string
	for _, component := range components {
		name := component[1]
		names = append(names, name)
	}

	return names
}

func createUIComponentFilePaths(names []string) []string {
	var paths []string
	for _, name := range names {
		path := fmt.Sprintf("components/%s.lamb.html", name)
		paths = append(paths, path)
	}
	return paths
}
