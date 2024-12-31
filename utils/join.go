package utils

import (
	"fmt"
	"strings"
)

// Join concatenates the elements of the provided slice into a single string
// with the specified separator between each element. The elements of the slice
// are converted to their string representations using fmt.Sprint.
//
// T is a type parameter that can be any type.
func Join[T any](slice []T, sep string) string {
	strs := make([]string, len(slice))
	for i, v := range slice {
		strs[i] = fmt.Sprint(v)
	}
	return strings.Join(strs, sep)
}
