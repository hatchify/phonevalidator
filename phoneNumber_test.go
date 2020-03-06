package phonevalidator

import "testing"

var (
	cleanPhoneNumberTests = []cleanTest{
		// With country code block
		cleanTest{
			value:    "1 (503) 543-7890",
			expected: "+15035437890",
		},
		cleanTest{
			value:    "15035437890",
			expected: "+15035437890",
		},
		cleanTest{
			value:    "1 503 543 7890",
			expected: "+15035437890",
		},
		cleanTest{
			value:    "1-503-543-7890",
			expected: "+15035437890",
		},
		cleanTest{
			value:    "1_503_543_7890",
			expected: "+15035437890",
		},

		// Without country code block
		cleanTest{
			value:    "(503) 543-7890",
			expected: "+15035437890",
		},
		cleanTest{
			value:    "5035437890",
			expected: "+15035437890",
		},
		cleanTest{
			value:    "503 543 7890",
			expected: "+15035437890",
		},
		cleanTest{
			value:    "503-543-7890",
			expected: "+15035437890",
		},
		cleanTest{
			value:    "503_543_7890",
			expected: "+15035437890",
		},
	}
)

func TestCleanPhoneNumber(t *testing.T) {
	// Iterate through all clean phone number tests
	for i, test := range cleanPhoneNumberTests {
		// Get cleaned version of phone number
		cleaned := clean(test.value)
		// Check to see if cleaned value matches expectation
		if cleaned != test.expected {
			t.Fatalf("invalid value, expected \"%s\" and recieved \"%s\" for test #%d", test.expected, cleaned, i)
		}
	}
}

// cleanTest is a simple struct used to setup clean testing
type cleanTest struct {
	value    string
	expected string
}
