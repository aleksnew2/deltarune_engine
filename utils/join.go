package utils

import (
	"fmt"
	"strings"
)

func Join[T any](slice []T, sep string) string {
	strs := make([]string, len(slice))
	for i, v := range slice {
		strs[i] = fmt.Sprint(v)
	}
	return strings.Join(strs, sep)
}
