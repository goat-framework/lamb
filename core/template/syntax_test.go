package template

import (
	"reflect"
	"testing"
)

func TestReplaceSyntax(t *testing.T) {
	example := `<h1>{{ title }}</h1>
<p>{{ text }}</p>
@if LoggedIn
<p>Hello {{ name }}</p>
@elseif NotLoggedin
<p>Sign Up</p>
@else
<p>Welcome</p>
@end
@for user in users
<p>{{ username }}</p>
@end`

	expected := `<h1>{{ .title }}</h1>
<p>{{ .text }}</p>
{{ if .LoggedIn }}
<p>Hello {{ .name }}</p>
{{ else if .NotLoggedin }}
<p>Sign Up</p>
{{ else }}
<p>Welcome</p>
{{ end }}
{{ range .users }}
<p>{{ .username }}</p>
{{ end }}`

	result := replaceSyntax(example)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
