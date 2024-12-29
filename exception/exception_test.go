package exception

import (
	"fmt"
	"testing"
)

func TestCreateException(t *testing.T) {
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
			err := CreateException(test.error_code, test.args...)
			if err == nil {
				t.Errorf("CreateException(%d, %v) = nil; want non-nil error", test.error_code, test.args)
			}
			if err.Error() != test.expected {
				t.Errorf("CreateException(%d, %v).Error() = %v; want %v", test.error_code, test.args, err.Error(), test.expected)
			}
		})
	}
}

func TestExceptionType(t *testing.T) {
	err := CreateException(100, "Test Error")
	if _, ok := err.(Exception); !ok {
		t.Errorf("Expected error of type Exception, but got %T", err)
	}
}
