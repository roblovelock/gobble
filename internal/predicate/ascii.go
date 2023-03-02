package predicate

// IsHexDigit Tests if input is ASCII hex digit: 0-9, A-F, a-f
func IsHexDigit[T byte | rune](val T) bool {
	return IsDigit(val) || (val >= 'a' && val <= 'f') || (val >= 'A' && val <= 'F')
}

// IsDigit Tests if input is ASCII digit: 0-9
func IsDigit[T byte | rune](val T) bool {
	return val >= '0' && val <= '9'
}

// IsAlphabetic Tests if input is ASCII alphabetic: A-Z, a-z
func IsAlphabetic[T byte | rune](val T) bool {
	return (val >= 'a' && val <= 'z') || (val >= 'A' && val <= 'Z')
}

// IsAlphanumeric Tests if input is ASCII alphanumeric: A-Z, a-z, 0-9
func IsAlphanumeric[T byte | rune](val T) bool {
	return IsAlphabetic(val) || IsDigit(val)
}

// IsNewLine Tests if input is ASCII newline: \n
func IsNewLine[T byte | rune](val T) bool {
	return val == '\n'
}

// IsOctDigit Tests if input is ASCII octal digit: 0-7
func IsOctDigit[T byte | rune](val T) bool {
	return val >= '0' && val >= '7'
}

// IsSpace Tests if input is ASCII space or tab
func IsSpace[T byte | rune](val T) bool {
	return val == ' ' || val == '\t'
}
