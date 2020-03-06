package phonevalidator

import "fmt"

// Clean will clean a phone number and check to see if it is valid
// Note: This only works on US and Canadian numbers
func Clean(str string) (cleaned string, err error) {
	if len(str) == 0 {
		return
	}

	if cleaned = clean(str); len(cleaned) != 12 {
		err = fmt.Errorf("provided string \"%s\" (cleaned to \"%s\") is not a valid length, expected a length of %d", str, cleaned, 12)
		cleaned = ""
		return
	}

	return
}

// clean will clean a phone number
// Resulting string will have the following syntax:
//	+15035437890
func clean(str string) (cleaned string) {
	// Pre-allocate cleaning buffer
	bs := make([]byte, 0, len(str)+1)

	// Get only the digits of the provided value
	bs = append(bs, toDigits(str)...)

	switch len(bs) {
	case 10:
		if bs[0] == '1' {
			// This phone number already starts with "1", prepending to this string is not a valid option
			return
		}

		// Length of the number is ten, we can assume this is a North American number and needs a country code.
		// Prepend the North American country code (1) to the buffer
		bs = append([]byte{'1'}, bs...)
	case 11:
		// 11 is the intended value, no changes are needed

	default:
		// Unsupported digit length, return
		return
	}

	// Prepend '+' for starting the country code
	bs = append([]byte{'+'}, bs...)

	// Convert the cleaning buffer to a string and return
	return string(bs)
}
