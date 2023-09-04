package functions

import "github.com/teris-io/shortid"

// This function returns a Unique ID which can be used for almost anything you want in your template.
//
// The mainly use case is when you are using Javascript and you need an UID in some of your HTML elements
// and you don't know if that ID already exists.
func UID() (string, error) {
	return shortid.Generate()
}
