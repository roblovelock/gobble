package ascii_test

import (
	"fmt"
	"github.com/roblovelock/gobble/pkg/parser/ascii"
	"io"
	"strings"
)

func ExampleAlphanumeric_match() {
	input := strings.NewReader("abc123")
	byteParser := ascii.Alphanumeric()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: 'a', Error: <nil>, Remainder: 'bc123'
}

func ExampleAlphanumeric_noMatch() {
	input := strings.NewReader("+-")
	byteParser := ascii.Alphanumeric()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'not matched', Remainder: '+-'
}

func ExampleAlphanumeric_endOfFile() {
	input := strings.NewReader("")
	byteParser := ascii.Alphanumeric()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'EOF', Remainder: ''
}

func ExampleAlphanumeric1_match() {
	input := strings.NewReader("abc123+")
	byteParser := ascii.Alphanumeric1()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: 'abc123', Error: <nil>, Remainder: '+'
}

func ExampleAlphanumeric1_noMatch() {
	input := strings.NewReader("+-")
	byteParser := ascii.Alphanumeric1()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: '%v', Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: 'not matched', Remainder: '+-'
}

func ExampleAlphanumeric1_endOfFile() {
	input := strings.NewReader("")
	byteParser := ascii.Alphanumeric1()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: '%v', Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: 'EOF', Remainder: ''
}

func ExampleAlphanumeric0_match() {
	input := strings.NewReader("abc123+")
	byteParser := ascii.Alphanumeric0()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: 'abc123', Error: <nil>, Remainder: '+'
}

func ExampleAlphanumeric0_noMatch() {
	input := strings.NewReader("+-")
	byteParser := ascii.Alphanumeric0()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: <nil>, Remainder: '+-'
}

func ExampleAlphanumeric0_endOfFile() {
	input := strings.NewReader("")
	byteParser := ascii.Alphanumeric0()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: <nil>, Remainder: ''
}
