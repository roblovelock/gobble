package ascii

const (
	digitFlag        uint8 = 0x01
	hexFlag          uint8 = 0x02
	octFlag          uint8 = 0x04
	letterFlag       uint8 = 0x08
	lowerLetterFlag  uint8 = 0x10
	upperLetterFlag  uint8 = 0x20
	alphanumericFlag uint8 = 0x40
)

var (
	lookupTable [256]uint8
)

func init() {
	for i := '0'; i <= '7'; i++ {
		lookupTable[i] = digitFlag | hexFlag | octFlag | alphanumericFlag
	}
	for i := '8'; i <= '9'; i++ {
		lookupTable[i] = digitFlag | hexFlag | alphanumericFlag
	}
	for i := 'a'; i <= 'f'; i++ {
		lookupTable[i] = hexFlag | alphanumericFlag | letterFlag | lowerLetterFlag
	}
	for i := 'A'; i <= 'F'; i++ {
		lookupTable[i] = hexFlag | alphanumericFlag | letterFlag | upperLetterFlag
	}
	for i := 'g'; i <= 'z'; i++ {
		lookupTable[i] = alphanumericFlag | letterFlag | lowerLetterFlag
	}
	for i := 'G'; i <= 'Z'; i++ {
		lookupTable[i] = alphanumericFlag | letterFlag | upperLetterFlag
	}
}

// IsHexDigit Tests if input is ASCII hex digit: 0-9, A-F, a-f
func IsHexDigit(val byte) bool {
	return lookupTable[val]&hexFlag == hexFlag
}

// IsDigit Tests if input is ASCII digit: 0-9
func IsDigit(val byte) bool {
	return lookupTable[val]&digitFlag == digitFlag
}

// IsLetter Tests if input is ASCII letter: A-Z, a-z
func IsLetter(val byte) bool {
	return lookupTable[val]&letterFlag == letterFlag
}

// IsLowercaseLetter Tests if input is ASCII lowercase letter: a-z
func IsLowercaseLetter(val byte) bool {
	return lookupTable[val]&lowerLetterFlag == lowerLetterFlag
}

// IsUppercaseLetter Tests if input is ASCII uppercase letter: A-Z
func IsUppercaseLetter(val byte) bool {
	return lookupTable[val]&upperLetterFlag == upperLetterFlag
}

// IsAlphanumeric Tests if input is ASCII alphanumeric: A-Z, a-z, 0-9
func IsAlphanumeric(val byte) bool {
	return lookupTable[val]&alphanumericFlag == alphanumericFlag
}

// IsASCII Tests if input is an ASCII character: 0x00-0x7F
func IsASCII(val byte) bool {
	return val <= 0x7F
}

// IsOctDigit Tests if input is ASCII octal digit: 0-7
func IsOctDigit(val byte) bool {
	return lookupTable[val]&octFlag == octFlag
}
