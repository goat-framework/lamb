package template

import (
	"fmt"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"testing"
)

func TestGetAttributes(t *testing.T) {
	element := `<ui-input class="text-blue-500 rounded-lg" type="email" name="email" />`

	attributes := getAttributes(element)

	expectedAttributes := map[string]string{
		"class": "text-blue-500 rounded-lg",
		"type":  "email",
		"name":  "email",
	}

	for key, expectedValue := range expectedAttributes {
		if value, exists := attributes[key]; !exists || value != expectedValue {
			t.Errorf("Expected attribute %s=\"%s\", but got %s\"%s\"", key, expectedValue, key, value)
		}
	}
}

func TestGetRootElement(t *testing.T) {
	element := `<ui-div class="container"><button class="button"></button></ui-div>`

	result := getRootElement(element)

	expected := `<ui-div class="container">`

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNestGetAttributes(t *testing.T) {
	element := `<ui-div class="container"><button class="button"></button></ui-div>`

	attributes := getAttributes(element)

	expectedAttributes := map[string]string{
		"class": "container",
	}

	for key, expectedValue := range expectedAttributes {
		if value, exists := attributes[key]; !exists || value != expectedValue {
			t.Errorf("Expected attribute %s=\"%s\", but got %s\"%s\"", key, expectedValue, key, value)
		}
	}
}

func TestMergeAttributes(t *testing.T) {
	original := make(Attributes)
	original["class"] = "btn"
	original["type"] = "submit"

	toMerge := make(Attributes)
	toMerge["class"] = "primary"
	toMerge["id"] = "button"

	expected := map[string]string{
		"class": "btn primary",
		"type":  "submit",
		"id":    "button",
	}

	result := original.mergeAttributes(toMerge)

	for key, expectedValue := range expected {
		if value, exists := result[key]; !exists || value != expectedValue {
			t.Errorf("Expected attribute %s=\"%s\", but got %s\"%s\"", key, expectedValue, key, value)
		}
	}
}

func TestMergeAttributesOverwrite(t *testing.T) {
	original := make(Attributes)
	original["class"] = "btn"
	// should be overwritten
	original["id"] = "submit"

	toMerge := make(Attributes)
	toMerge["class"] = "primary"
	// should overwrite submit id
	toMerge["id"] = "button"

	expected := map[string]string{
		"class": "btn primary",
		"id":    "button",
	}

	result := original.mergeAttributes(toMerge)

	for key, expectedValue := range expected {
		if value, exists := result[key]; !exists || value != expectedValue {
			t.Errorf("Expected attribute %s=\"%s\", but got %s\"%s\"", key, expectedValue, key, value)
		}
	}
}

func TestExtractAttributesString(t *testing.T) {
	example := `<div @attributes("class": "container flex", "id": "div1")`

	expected := `"class": "container flex", "id": "div1"`

	result := extractAttributesString(example)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestParseAttributesString(t *testing.T) {
	example := `"class": "container flex", "id": "div1"`

	expected := make(Attributes)
	expected["class"] = "container flex"
	expected["id"] = "div1"

	result := parseAttributesString(example)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestApplyAttributesDirectiveAddID(t *testing.T) {
	attributes := make(Attributes)
	attributes["class"] = "btn primary"
	attributes["type"] = "submit"

	content := `<button @attributes("id": "btn")>Submit</button>`

	expected := `<button type="submit" id="btn" class="btn primary">Submit</button>`

	result := applyAttributesDirective(content, attributes)

	expectedSorted := sortAttributes(expected)
	resultSorted := sortAttributes(result)

	if !reflect.DeepEqual(resultSorted, expectedSorted) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestApplyAttributesDirectiveReplaceID(t *testing.T) {
	attributes := make(Attributes)
	attributes["class"] = "btn primary"
	attributes["id"] = "newID"

	content := `<button @attributes("id": "oldID")>Submit</button>`

	expected := `<button class="btn primary" id="newID">Submit</button>`

	result := applyAttributesDirective(content, attributes)

	expectedSorted := sortAttributes(expected)
	resultSorted := sortAttributes(result)

	if !reflect.DeepEqual(resultSorted, expectedSorted) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestApplyAttributesDirectiveAddClasses(t *testing.T) {
	attributes := make(Attributes)
	attributes["class"] = "newClass"
	attributes["type"] = "submit"

	content := `<button @attributes("class": "original classes")>Submit</button>`

	expected := `<button class="newClass original classes" type="submit">Submit</button>`

	result := applyAttributesDirective(content, attributes)

	expectedSorted := sortAttributes(expected)
	resultSorted := sortAttributes(result)

	if !reflect.DeepEqual(resultSorted, expectedSorted) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestApplyAttributesDirectiveAllThree(t *testing.T) {
	attributes := make(Attributes)
	attributes["class"] = "newClass"
	attributes["id"] = "newID"
	attributes["type"] = "submit"

	content := `<button @attributes("class": "original classes", "id": "oldID")>Submit</button>`

	expected := `<button class="newClass original classes" id="newID" type="submit">Submit</button>`

	result := applyAttributesDirective(content, attributes)

	expectedSorted := sortAttributes(expected)
	resultSorted := sortAttributes(result)

	if !reflect.DeepEqual(resultSorted, expectedSorted) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestApplyAttributesDirectiveEmpty(t *testing.T) {
	attributes := make(Attributes)
	attributes["id"] = "newID"

	content := `<button @attributes()>Submit</button>`

	expected := `<button id="newID">Submit</button>`

	result := applyAttributesDirective(content, attributes)

	expectedSorted := sortAttributes(expected)
	resultSorted := sortAttributes(result)

	if !reflect.DeepEqual(resultSorted, expectedSorted) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestApplyAttributesDirectiveEmptyClass(t *testing.T) {
	attributes := make(Attributes)
	attributes["class"] = "newClass"

	content := `<button @attributes()>Submit</button>`

	expected := `<button class="newClass">Submit</button>`

	result := applyAttributesDirective(content, attributes)

	expectedSorted := sortAttributes(expected)
	resultSorted := sortAttributes(result)

	if !reflect.DeepEqual(resultSorted, expectedSorted) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestApplyAttributesDirectiveEmptyMultiple(t *testing.T) {
	attributes := make(Attributes)
	attributes["class"] = "newClass"
	attributes["id"] = "newID"
	attributes["type"] = "submit"

	content := `<button @attributes()>Submit</button>`

	expected := `<button class="newClass" id="newID" type="submit">Submit</button>`

	result := applyAttributesDirective(content, attributes)

	expectedSorted := sortAttributes(expected)
	resultSorted := sortAttributes(result)

	if !reflect.DeepEqual(resultSorted, expectedSorted) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func sortAttributes(html string) string {
	regex := regexp.MustCompile(`</s*(\w+)([^>]*)>`)
	match := regex.FindStringSubmatch(html)
	if len(match) == 0 {
		return html
	}

	tag := match[1]
	attrs := strings.Fields(match[2])

	sort.Strings(attrs)

	return fmt.Sprintf("<%s %s>", tag, strings.Join(attrs, " "))
}
