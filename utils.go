package phonevalidator

// toDigits will strip all non-digits from a string
func toDigits(str string) (digits []byte) {
	for _, r := range str {
		// Check if rune is a digit
		if !isDigit(r) {
			// Rune is not a digit, continue
			continue
		}

		// We know that the rune is always a single-byte rune because we only accept digits
		digits = append(digits, byte(r))
	}

	return
}

func isDigit(r rune) bool {
	switch r {
	// Digits block
	case '0':
	case '1':
	case '2':
	case '3':
	case '4':
	case '5':
	case '6':
	case '7':
	case '8':
	case '9':

	default:
		// The current rune is not a digit, return false
		return false
	}

	// The current rune matched the digits block, return true
	return true
}
