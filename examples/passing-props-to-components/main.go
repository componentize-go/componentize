package main

import (
	"html/template"
	"os"

	"github.com/componentize-go/componentize"
)

func main() {
	myIndexTmpl := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Document</title>
		</head>
		<body>
			<div>
				<ul>
					{{template "components/nav-menu-button" KVM "link" "/contact" | KVM "content" "Contact me"}}
					{{template "components/nav-menu-button" KVM "link" "/buy" | KVM "content" "Buy the product"}}
					{{template "components/nav-menu-button" KVM "link" "/blog" | KVM "content" "Visit my blog"}}
				</ul>
			</div>
		</body>
		</html>
	`
	myNavMenuBtn := `
		<li>
			<a href="{{.link}}">{{.content}}</a>
		</li>
	`
	componentizeFuncs := componentize.Default()

	tmpl, err := template.New("index").Funcs(componentizeFuncs).Parse(myIndexTmpl)
	if err != nil {
		panic(err)
	}

	tmpl, err = tmpl.New("components/nav-menu-button").Parse(myNavMenuBtn)
	if err != nil {
		panic(err)
	}

	if err := tmpl.ExecuteTemplate(os.Stdout, "index", nil); err != nil {
		panic(err)
	}
}
