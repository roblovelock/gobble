package ascii

// IsHexDigit Tests if input is ASCII hex digit: 0-9, A-F, a-f
func IsHexDigit(val byte) bool {
	return IsDigit(val) || (val >= 'a' && val <= 'f') || (val >= 'A' && val <= 'F')
}

// IsDigit Tests if input is ASCII digit: 0-9
func IsDigit(val byte) bool {
	return val >= '0' && val <= '9'
}

// IsLetter Tests if input is ASCII letter: A-Z, a-z
func IsLetter(val byte) bool {
	return (val >= 'a' && val <= 'z') || (val >= 'A' && val <= 'Z')
}

// IsAlphanumeric Tests if input is ASCII alphanumeric: A-Z, a-z, 0-9
func IsAlphanumeric(val byte) bool {
	return IsLetter(val) || IsDigit(val)
}

// IsASCII Tests if input is an ASCII character: 0x00-0x7F
func IsASCII(val byte) bool {
	return val <= 0x7F
}

// IsOctDigit Tests if input is ASCII octal digit: 0-7
func IsOctDigit(val byte) bool {
	return val >= '0' && val <= '7'
}
