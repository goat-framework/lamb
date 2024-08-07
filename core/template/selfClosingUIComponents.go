package template

import (
	"fmt"
	"regexp"
)

type SelfClosingUIComponent struct {
	ComponentName     string
	ComponentFilePath string
}

func GetSelfClosingUIComponents(content string) []SelfClosingUIComponent {
	components := FindSelfClosingUIComponents(content)
	names := CreateUIComponentNames(components)
	paths := CreateUIComponentFilePaths(names)

	var structs []SelfClosingUIComponent

	for i := range components {
		structs = append(structs, SelfClosingUIComponent{
			ComponentName:     names[i],
			ComponentFilePath: paths[i],
		})
	}

	return structs
}

func FindSelfClosingUIComponents(content string) [][]string {
	regex := regexp.MustCompile(`<ui-([\w-]+)\s*/>`)
	matches := regex.FindAllStringSubmatch(content, -1)
	return matches
}

func CreateUIComponentNames(components [][]string) []string {
	var names []string
	for _, component := range components {
		name := component[1]
		names = append(names, name)
	}

	return names
}

func CreateUIComponentFilePaths(names []string) []string {
	var paths []string
	for _, name := range names {
		path := fmt.Sprintf("components/%s.lamb.html", name)
		paths = append(paths, path)
	}
	return paths
}
