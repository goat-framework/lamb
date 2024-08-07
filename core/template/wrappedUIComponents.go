package template

import (
	"regexp"
)

type WrappedUIComponent struct {
	ComponentName     string
	ComponentFilePath string
	InnerContent      string
}

func GetWrappedUIComponents(content string) []WrappedUIComponent {
	components := FindWrappedUIComponents(content)
	names := CreateUIComponentNames(components)
	contents := getComponentInnerContent(components)
	paths := CreateUIComponentFilePaths(names)

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

func FindWrappedUIComponents(content string) [][]string {
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
