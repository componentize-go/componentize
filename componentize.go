package componentize

import (
	"html/template"

	"github.com/componentize-go/componentize/functions"
)

// The default constructor for the `componentize` FuncMap, it appends all the functions in the `functions` package
// and returns it like a FuncMap.
//
// If you don't want all the features, either because you've already defined your own with the same
// functionality or name, or because you simply only need a few of them, you can use the `WithConfig` constructor
// with whatever functions you want.
func Default() template.FuncMap {
	m := map[string]any{}
	m["KVM"] = functions.KVM
	m["UID"] = functions.UID

	return m
}

type Config struct {
	// Mark as `true` if you want to use the `KVM` function.
	//
	// default: `false`
	UsingKVM bool

	// Mark as `true` if you want to use the `UID` function.
	//
	// default: `false`
	UsingUID bool
}

// Use this constructor if you don't want to use all of the functions that provides the Default constructor
//
// You can select what functions you want with the `config` param.
func WithConfig(config Config) template.FuncMap {
	m := map[string]any{}
	if config.UsingKVM {
		m["KVM"] = functions.KVM
	}

	if config.UsingUID {
		m["UID"] = functions.UID
	}
	return m
}
