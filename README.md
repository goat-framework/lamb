<div align="center">


[![Test](https://github.com/goat-framework/lamb/actions/workflows/go.yml/badge.svg)](https://github.com/goat-framework/lamb/actions)
[![Report Card](https://goreportcard.com/badge/github.com/goat-framework/lamb)](https://goreportcard.com/report/github.com/goat-framework/lamb)
![License](https://img.shields.io/github/license/goat-framework/lamb)


</div>

# Lamb

Intuitive templating engine for Go!

## Examples

These new directives clean up your templates making them easy as pie to read!

`<h1>Title</h1>
@if LoggedIn
<p>Hello {{ User }}</p>
@else
<a href="login">Login</a>
@end
`

These directives compile directly into the standard libary's syntax!

`<h1>Title</h1>
{{ if .LoggedIn }}
<p>Hello {{ .User }}</p>
{{ else }}
<a href="login">Login</a>
{{ end }}
`

Easier on the eyes right?

Having troubles reusing components in your application?
Check out the ui components.

_main.html_
`
<form>
  <ui-input />
  <button>Submit</button>
</form>
`

_components/input.html_
`
<input type="text" />
`

_compiled_
`
<form>
  <input type="text" />
  <button>Submit</button>
</form>

