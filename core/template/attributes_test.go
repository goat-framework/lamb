package template

import (
	"testing"
    "reflect"
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
