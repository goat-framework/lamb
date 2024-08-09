package template

import (
	"reflect"
	"testing"
)

var exampleHtml string = "<h1>Title</h1><p>some text</p><ui-link /><ui-button>Submit</ui-button><ui-input /><ui-container><div>Container</div></ui-container>"

func TestFindSelfClosingUIComponents(t *testing.T) {
	expected := []string{
		"<ui-link />",
		"<ui-input />",
	}

	result := findSelfClosingUIElements(exampleHtml)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestFindSelfClosingUIComponentsWithAttributes(t *testing.T) {
	expected := []string{
		"<ui-link class=\"link blue brown\" id=\"primary\" />",
		"<ui-button class=\"button\" />",
	}

	example := `<h1>Title</h1><ui-link class="link blue brown" id="primary" /><ui-button class="button" /><footer>end</footer>`

	result := findSelfClosingUIElements(example)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestFindWrappedUIComponents(t *testing.T) {
	expected := []string{
		"<ui-button>Submit</ui-button>",
		"<ui-container><div>Container</div></ui-container>",
	}

	result := findWrappedUIElements(exampleHtml)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestGetComponentInnerContent(t *testing.T) {
	components := findWrappedUIElements(exampleHtml)

	expected := []string{"Submit", "<div>Container</div>"}

	result := getComponentInnerContent(components)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestCreateComponentUINames(t *testing.T) {
	components := findSelfClosingUIElements(exampleHtml)

	expected := []string{"link", "input"}

	result := createUIComponentNames(components)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestCreateUIComponentFilePaths(t *testing.T) {
	components := findSelfClosingUIElements(exampleHtml)
	names := createUIComponentNames(components)

	expected := []string{"components/link.lamb.html", "components/input.lamb.html"}

	result := createUIComponentFilePaths(names)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
