package functions

import (
	"fmt"
	"strings"
)

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
	if strings.TrimSpace(key) == "" {
		return nil, fmt.Errorf("")
	}
	m := map[string]any{}
	m[key] = value

	if len(otherMap) == 0 {
		return m, nil
	}

	for k, v := range otherMap[0] {
		m[k] = v
	}

	return m, nil
}
