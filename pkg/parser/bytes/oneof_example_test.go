package bytes_test

import (
	"fmt"
	"github.com/roblovelock/gobble/pkg/parser/bytes"
	"io"
	"strings"
)

func ExampleOneOf_match() {
	input := strings.NewReader("abc123")
	byteParser := bytes.OneOf('a', 'b', 'c')

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: 'a', Error: <nil>, Remainder: 'bc123'
}

func ExampleOneOf_noMatch() {
	input := strings.NewReader("123")
	byteParser := bytes.OneOf('a', 'b', 'c')

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'not matched', Remainder: '123'
}

func ExampleOneOf_endOfFile() {
	input := strings.NewReader("")
	byteParser := bytes.OneOf('a', 'b', 'c')

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'EOF', Remainder: ''
}

func ExampleOneOf1_match() {
	input := strings.NewReader("abc123")
	byteParser := bytes.OneOf1('a', 'b', 'c')

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: 'abc', Error: <nil>, Remainder: '123'
}

func ExampleOneOf1_noMatch() {
	input := strings.NewReader("123")
	byteParser := bytes.OneOf1('a', 'b', 'c')

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: '%v', Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: 'not matched', Remainder: '123'
}

func ExampleOneOf1_endOfFile() {
	input := strings.NewReader("")
	byteParser := bytes.OneOf1('a', 'b', 'c')

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: '%v', Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: 'EOF', Remainder: ''
}

func ExampleOneOf0_match() {
	input := strings.NewReader("abc123")
	byteParser := bytes.OneOf0('a', 'b', 'c')

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: 'abc', Error: <nil>, Remainder: '123'
}

func ExampleOneOf0_noMatch() {
	input := strings.NewReader("123")
	byteParser := bytes.OneOf0('a', 'b', 'c')

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: <nil>, Remainder: '123'
}

func ExampleOneOf0_endOfFile() {
	input := strings.NewReader("")
	byteParser := bytes.OneOf0('a', 'b', 'c')

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: <nil>, Remainder: ''
}
