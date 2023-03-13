package bytes_test

import (
	"fmt"
	"github.com/roblovelock/gobble/pkg/parser/bytes"
	"io"
	"strings"
)

func ExampleByte_match() {
	input := strings.NewReader("abc")
	byteParser := bytes.Byte('a')

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: 'a', Error: <nil>, Remainder: 'bc'
}

func ExampleByte_noMatch() {
	input := strings.NewReader("abc")
	byteParser := bytes.Byte('b')

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'not matched', Remainder: 'abc'
}

func ExampleByte_endOfFile() {
	input := strings.NewReader("")
	byteParser := bytes.Byte('a')

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'EOF', Remainder: ''
}
