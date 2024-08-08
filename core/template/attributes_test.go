package template

import (
	"testing"
)

func TestGetAttributes(t *testing.T) {
	element := `<input class="text-blue-500 rounded-lg" type="email" name="email" />`

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
