package template

import (
	"reflect"
	"testing"
)

var exampleHtml string = "<h1>Title</h1><p>some text</p><ui-link /><ui-button>Submit</ui-button><ui-input /><ui-container><div>Container</div></ui-container>"

func TestFindSelfClosingUIComponents(t *testing.T) {
	expected := [][]string{
		{"<ui-link />", "link"},
		{"<ui-input />", "input"},
	}

	result := findSelfClosingUIComponents(exampleHtml)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestFindWrappedUIComponents(t *testing.T) {
	expected := [][]string{
		{"<ui-button>Submit</ui-button>", "button", "Submit"},
		{"<ui-container><div>Container</div></ui-container>", "container", "<div>Container</div>"},
	}

	result := findWrappedUIComponents(exampleHtml)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestGetComponentInnerContent(t *testing.T) {
	components := findWrappedUIComponents(exampleHtml)

	expected := []string{"Submit", "<div>Container</div>"}

	result := getComponentInnerContent(components)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestCreateComponentUINames(t *testing.T) {
	components := findSelfClosingUIComponents(exampleHtml)

	expected := []string{"link", "input"}

	result := createUIComponentNames(components)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestCreateUIComponentFilePaths(t *testing.T) {
	components := findSelfClosingUIComponents(exampleHtml)
	names := createUIComponentNames(components)

	expected := []string{"components/link.lamb.html", "components/input.lamb.html"}

	result := createUIComponentFilePaths(names)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
