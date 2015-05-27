# TemplateCat

A CLI Golang template processor for appending multiple files into a single template.

## Building
```bash
go build templatecat.go
```

## usage
```bash
./templatecat variable@path, ... < template > processed_template
```

### Example

index.html.template
```html
{{ .header }}

BODY

{{ .footer }}
```

header.html
```html
<h1>This the page header!</h1>
```

footer.html
```html
<footer>Page Footer</footer>
```

```bash
./tempatecat header@header.html footer@footer.html < index.html.template > index.html
```

Produces:
```
<h1>This the page header!</h1>

BODY

<footer>Page Footer</footer>
```
