package is

import (
	"fmt"
	"testing"
)

func TestNewNotPointerError(t *testing.T) {
	var nilInterface interface{} = nil

	testCases := []interface{}{
		nilInterface,
		42,
		new(int),
	}

	for _, testCase := range testCases {
		t.Run("NewNotPointerError", func(t *testing.T) {
			err := NewNotPointerError(testCase)
			if err.Value != testCase {
				t.Errorf("Value not stored on NotPointerError, expected: %+v, got: %+v", err.Value, testCase)
			}

			expectedMessage := "Provided value is nil"

			if testCase == nil && err.Error() != "Provided value is nil" {
				t.Errorf("Expected error message not returned. Expected: %+v; Got: %+v", expectedMessage, err.Error())
			}

			if testCase != nil {
				expectedMessage = fmt.Sprintf("Provided value (%T) is not a pointer: %+v", err.Value, err.Value)

				if err.Error() != expectedMessage {
					t.Errorf("Expected error message not returned. Expected: %+v; Got: %+v", expectedMessage, err.Error())
				}
			}
		})
	}
}
