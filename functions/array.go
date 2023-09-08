package functions

// Function that returns all of its args inside of a array, basically is an array constructor
func Array(args ...any) []any {
	arr := make([]any, 0, len(args))

	arr = append(arr, args...)

	return arr
}
