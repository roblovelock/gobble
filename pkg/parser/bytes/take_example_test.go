package bytes_test

import (
	"fmt"
	"github.com/roblovelock/gobble/pkg/parser/bytes"
	"io"
	"strings"
)

func ExampleTake_match() {
	input := strings.NewReader("abc")
	byteParser := bytes.Take(2)

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: 'ab', Error: <nil>, Remainder: 'c'
}

func ExampleTake_unexpectedEndOfFile() {
	input := strings.NewReader("abc")
	byteParser := bytes.Take(4)

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: '%v', Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: 'EOF', Remainder: 'abc'
}

func ExampleTake_endOfFile() {
	input := strings.NewReader("")
	byteParser := bytes.Take(4)

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: '%v', Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: 'EOF', Remainder: ''
}
