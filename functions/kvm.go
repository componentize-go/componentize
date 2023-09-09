package functions

import (
	"fmt"
	"regexp"
)

var validKeyRegExp = regexp.MustCompile(`^[^0-9][a-zA-Z0-9_]*$`)

// KVM stands for "Key Value Map", it's a function that returns a map with the `key` and `value`
// added to the `otherMap` parameter if it's passed in.
//
// The `otherMap` parameter is optional.
//
// The mainly use that you can give it is, when you want to pass more than one value to a template
// you have to create a struct and then if in your template you render more than one template that
// needs more data you have to create a bigger struct and that is incovenient. So you can simply use
// this function inside of your template and then create "On The Way" your struct inside of the template
// rather than in your source Go files.
func KVM(key string, value any, otherMap ...map[string]any) (map[string]any, error) {
	if !validKeyRegExp.MatchString(key) {
		return nil, fmt.Errorf(`the key "%s" must be a alpha-numeric string but must no start with a number`, key)
	}

	if len(otherMap) == 0 {
		return map[string]any{key: value}, nil
	}

	if otherMap[0] == nil {
		return nil, fmt.Errorf("the third argument is nil, must pass a valid KVM return value or a map[string]any")
	}

	m := make(map[string]any, len(otherMap[0]))
	m[key] = value

	for k, v := range otherMap[0] {
		m[k] = v
	}

	return m, nil
}
