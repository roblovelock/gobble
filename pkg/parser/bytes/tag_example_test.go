package bytes_test

import (
	"fmt"
	"github.com/roblovelock/gobble/pkg/parser/bytes"
	"io"
	"strings"
)

func ExampleTag_match() {
	input := strings.NewReader("abc")
	byteParser := bytes.Tag([]byte("ab"))

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: 'ab', Error: <nil>, Remainder: 'c'
}

func ExampleTag_noMatch() {
	input := strings.NewReader("abc")
	byteParser := bytes.Tag([]byte("bc"))

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: '%v', Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: 'not matched', Remainder: 'abc'
}

func ExampleTag_endOfFile() {
	input := strings.NewReader("")
	byteParser := bytes.Tag([]byte("ab"))

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: '%v', Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: 'EOF', Remainder: ''
}
