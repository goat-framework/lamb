<div align="center">


[![Test](https://github.com/goat-framework/lamb/actions/workflows/go.yml/badge.svg)](https://github.com/goat-framework/lamb/actions)
[![Report Card](https://goreportcard.com/badge/github.com/goat-framework/lamb)](https://goreportcard.com/report/github.com/goat-framework/lamb)
![License](https://img.shields.io/github/license/goat-framework/lamb)


</div>

# Lamb

Intuitive templating engine for Go!

## Directives

These new directives clean up your templates making them easy as pie to read!

```<h1>Title</h1>
@if LoggedIn
<p>Hello {{ User }}</p>
@else
<a href="login">Login</a>
@end
```

These directives compile directly into the standard libary's syntax!

```<h1>Title</h1>
{{ if .LoggedIn }}
<p>Hello {{ .User }}</p>
{{ else }}
<a href="login">Login</a>
{{ end }}
```

Easier on the eyes right?

## UI Components

Having troubles reusing components in your application?
Check out the ui components.

_main.lamb.html_
```
<form>
  <ui-input />
  <button>Submit</button>
</form>
```

_components/input.lamb.html_
```
<input type="text" />
```

_compiled_
```
<form>
  <input type="text" />
  <button>Submit</button>
</form>
```

## Wrap UI Components

Need to wrap some content in a custom component?

_main.lamb.html_
```
<form>
  <input />
  <ui-button>Submit</ui-button>
</form>
```

_components/button.lamb.html_
```
<button class="btn primary"><slot /></button>
```

_compiled_
```
<form>
  <input />
  <button class="btn primary">Submit</button>
</form>
```

## Create Themes With Ease

Let's mix and match to create an awesome form!

_main.lamb.html_
```
<ui-form>
  <ui-label>Email<ui-label>
  <ui-input />
  <ui-label>Name<ui-label>
  <ui-input />
  <ui-button>Submit</ui-button>
</ui-form>
```

_components/form.lamb.html_
```
<form class="flex flex-col space-y-4">
  <slot />
</form>
```

_components/label.lamb.html_
```
<label class="text-gray-800 font-bold"><slot /></label>
```

_components/input.lamb.html_
```
<input type="text" class="rounded-md border-blue-500" />
```

_components/button.lamb.html_
```
<button class="bg-blue-500 rounded-lg text-white">
  <slot />
</button>
```

_compiled_
```
<form class="flex flex-col space-y-4">
  <label class="text-gray-800 font-bold">Email</label>
  <input type="text" class="rounded-md border-blue-500" />
  <label class="text-gray-800 font-bold">Name</label>
  <input type="text" class="rounded-md border-blue-500" />
  <button class="bg-blue-500 rounded-lg text-white">Submit</button>
</form>
```
