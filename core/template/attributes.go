package template

import (
	"regexp"
)

// Get the attributes from the specified element
//
// Params:
// - element (string): html element
//
// Returns:
// - map[string]string: attribute key value pairs
//
// Since: 0.1.0
func getAttributes(element string) map[string]string {
	attributes := make(map[string]string)

	regex := regexp.MustCompile(`(\w+)="([^"]*)"`)
	matches := regex.FindAllStringSubmatch(element, -1)

	for _, match := range matches {
		key := match[1]
		value := match[2]
		attributes[key] = value
	}

	return attributes
}
