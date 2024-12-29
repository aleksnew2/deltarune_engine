package exception

import (
	"fmt"
	"testing"
)

func TestFormat(t *testing.T) {
	tests := []struct {
		error_code int32
		args       []string
		expected   string
	}{
		{123, []string{}, "deltarune_exception [123]"},
		{456, []string{"Error occurred"}, "deltarune_exception [456]: Error occurred"},
		{789, []string{"File not found", "in directory"}, "deltarune_exception [789]: File not found in directory"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("error_code_%d", test.error_code), func(t *testing.T) {
			got := format(test.error_code, test.args...)
			if got != test.expected {
				t.Errorf("format(%d, %v) = %v; want %v", test.error_code, test.args, got, test.expected)
			}
		})
	}
}


