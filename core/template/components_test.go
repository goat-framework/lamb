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

func TestFindSelfClosingWithinWrapped(t *testing.T) {
	expected := []string{
		"<ui-link />",
	}

	example := `<h1>title</h1><ui-container><ui-link /><p>Some Text</p></ui-container><ui-button>Submit</ui-button>`

	result := findSelfClosingUIElements(example)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestFindSelfClosingWithinWrappedWithAttributes(t *testing.T) {
	expected := []string{
		"<ui-link class=\"primary\" />",
	}

	example := `<h1>title</h1><ui-container class="wrapper"><ui-link class="primary" /></ui-container>`

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

func TestFindWrappedUIComponentsBlock(t *testing.T) {
	expected := []string{
		`<ui-wrapper>
    Some wrapped content
</ui-wrapper>`,
	}

	example := `<h1></h1><ui-wrapper>
    Some wrapped content
</ui-wrapper>`

	result := findWrappedUIElements(example)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestFindWrappedUIComponentsWithAttributes(t *testing.T) {
	expected := []string{
		"<ui-container class=\"wrapper\">Content</ui-container>",
		"<ui-button class=\"btn\" type=\"submit\">Submit</ui-button>",
	}

	example := `<h1>Hello</h1><ui-container class="wrapper">Content</ui-container><ui-button class="btn" type="submit">Submit</ui-button>`

	result := findWrappedUIElements(example)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestFindWrappedComponentsWithSelfClosingWithin(t *testing.T) {
	expected := []string{
		"<ui-container><ui-link /><p>Some Text</p></ui-container>",
		"<ui-button>Submit</ui-button>",
	}
	example := `<h1>title</h1><ui-container><ui-link /><p>Some Text</p></ui-container><ui-button>Submit</ui-button>`

	result := findWrappedUIElements(example)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestFindWrappedComponentsWithAttributesWithSelfClosing(t *testing.T) {
	expected := []string{
		"<ui-container class=\"wrapper\"><ui-link class=\"primary\" /></ui-container>",
	}
	example := `<h1>title</h1><ui-container class="wrapper"><ui-link class="primary" /></ui-container>`

	result := findWrappedUIElements(example)

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

func TestGetComponentInnerContentBlock(t *testing.T) {
	example := []string{`<ui-layout>
    This is some content
</ui-layout>`}

	expected := []string{`
    This is some content
`}

	result := getComponentInnerContent(example)

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

	result := createUIComponentFilePaths("components", names)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
