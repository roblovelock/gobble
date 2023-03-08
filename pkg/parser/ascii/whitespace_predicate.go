package ascii

const (
	spaceFlag      uint8 = 0x01
	blankSpaceFlag uint8 = 0x02
	whitespaceFlag uint8 = 0x04
)

var (
	whitespaceLookupTable [256]uint8
)

func init() {
	whitespaceLookupTable[' '] = spaceFlag | blankSpaceFlag | whitespaceFlag
	whitespaceLookupTable['\t'] = spaceFlag | blankSpaceFlag | whitespaceFlag
	whitespaceLookupTable['\r'] = blankSpaceFlag | whitespaceFlag
	whitespaceLookupTable['\n'] = blankSpaceFlag | whitespaceFlag
	whitespaceLookupTable['\v'] = whitespaceFlag
	whitespaceLookupTable['\f'] = whitespaceFlag
}

// IsNewLine Tests if input is ASCII newline: [\n]
func IsNewLine(val byte) bool {
	return val == '\n'
}

// IsSpace Tests if input is ASCII space or tab: [ \t]
func IsSpace(val byte) bool {
	return whitespaceLookupTable[val]&spaceFlag == spaceFlag
}

// IsBlankSpace Tests if input is ASCII space, tab, newline: [ \t\r\n]
func IsBlankSpace(val byte) bool {
	return whitespaceLookupTable[val]&blankSpaceFlag == blankSpaceFlag
}

// IsWhitespace Tests if input is ASCII whitespace: [ \t\r\n\v\f]
func IsWhitespace(val byte) bool {
	return whitespaceLookupTable[val]&whitespaceFlag == whitespaceFlag
}
