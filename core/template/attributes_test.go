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
