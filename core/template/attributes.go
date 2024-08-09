package template

import (
	"fmt"
	"regexp"
	"strings"
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

// Convert Attribute map to html string
//
// Receiver:
// - a (Attributes): the attributes map
//
// Returns:
// - string: html string
// ex: class="btn"
//
// Since: 0.1.0
func (a Attributes) toString() string {
	parts := make([]string, 0, len(a))
	for key, value := range a {
		parts = append(parts, fmt.Sprintf(`%s="%s"`, key, value))
	}
	return strings.Join(parts, " ")
}

// Extracts the @attributes map from content
//
// Params:
// - content (string): the content
//
// Returns:
// - string: html attribute string
//
// Since: 0.1.0
func extractAttributesString(content string) string {
	regex := regexp.MustCompile(`@attributes\(([^)]+)\)`)
	matches := regex.FindStringSubmatch(content)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

// Convert html @attribute string to Attributes map
//
// Params:
// - attributesString (string): html attribute string
//
// Returns:
// - Attributes: attribute map
//
// Since: 0.1.0
func parseAttributesString(attributesString string) Attributes {
	attributePairs := strings.Split(attributesString, ",")
	attributes := make(Attributes)

	for _, pair := range attributePairs {
		parts := strings.SplitN(strings.TrimSpace(pair), ":", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			key = strings.Trim(key, `"`)
			value = strings.Trim(value, `"`)
			attributes[key] = value
		}
	}

	return attributes
}

// Pass parent attributes to @attributes directive in
// the child component
//
// Params:
// - content (string): the content to parse
// - parentAttributes (Attributes): parent attributes map
//
// Returns:
// - string: html element with parent attributes mapped
// to the child component @attribute directive
//
// Since: 0.1.0
func applyAttributesDirective(content string, parentAttributes Attributes) string {
	regex := regexp.MustCompile(`@attributes\(([^)]+)\)`)

	return regex.ReplaceAllStringFunc(content, func(match string) string {
		attributesString := extractAttributesString(match)
		childAttributes := parseAttributesString(attributesString)
		mergedAttributes := parentAttributes.mergeAttributes(childAttributes)
		return mergedAttributes.toString()
	})
}
