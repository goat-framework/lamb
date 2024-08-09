package template

import (
	"regexp"
)

// A map of key value pairs for attributes
//
// ex: map[string]string{"class": "btn", "type": "submit",}
//
// Since: 0.1.0
type Attributes map[string]string

// Get the attributes from the specified element
//
// Params:
// - element (string): html element
//
// Returns:
// - map[string]string: attribute key value pairs
//
// Since: 0.1.0
func getAttributes(element string) Attributes {
	attributes := make(Attributes)

    element = getRootElement(element)

	regex := regexp.MustCompile(`(\w+)="([^"]*)"`)
	matches := regex.FindAllStringSubmatch(element, -1)

	for _, match := range matches {
		key := match[1]
		value := match[2]
		attributes[key] = value
	}

	return attributes
}

// Gets the root UI element tag
//
// Params:
// - element (string): element to pull tag from
//
// Returns:
// - string: the first ui tag
//
// Since: 0.1.0
func getRootElement(element string) string {
    regex := regexp.MustCompile(`<ui-[\w-]+(?:\s[^>]*)?/?>`)

    match := regex.FindString(element)

    return match
}

// Merges another Attributes map into the current one.
// If a key exists in both maps, the other map will override.
// If key is class, then combine them.
//
// Receiver:
// a (Attributes): current attributes map
//
// Params:
// - other (Attributes): attributes to be added
//
// Returns:
// - Attributes: combined attributes map
//
// Since: 0.1.0
func (a Attributes) mergeAttributes(other Attributes) Attributes {
	for key, value := range other {
		if key == "class" {
			a[key] = a[key] + " " + value
		} else {
			a[key] = value
		}
	}

	return a
}
