package ascii_test

import (
	"fmt"
	"github.com/roblovelock/gobble/pkg/parser/ascii"
	"io"
	"strings"
)

func ExampleDigit_match() {
	input := strings.NewReader("123")
	byteParser := ascii.Digit()

	match, err := byteParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '1', Error: <nil>, Remainder: '23'
}

func ExampleDigit_noMatch() {
	input := strings.NewReader("abc")
	byteParser := ascii.Digit()

	match, err := byteParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'not matched', Remainder: 'abc'
}

func ExampleDigit_endOfFile() {
	input := strings.NewReader("")
	byteParser := ascii.Digit()

	match, err := byteParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'EOF', Remainder: ''
}

func ExampleDigit1_match() {
	input := strings.NewReader("123abc")
	byteParser := ascii.Digit1()

	match, err := byteParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '123', Error: <nil>, Remainder: 'abc'
}

func ExampleDigit1_noMatch() {
	input := strings.NewReader("abc")
	byteParser := ascii.Digit1()

	match, err := byteParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: '%v', Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: 'not matched', Remainder: 'abc'
}

func ExampleDigit1_endOfFile() {
	input := strings.NewReader("")
	byteParser := ascii.Digit1()

	match, err := byteParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: '%v', Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: 'EOF', Remainder: ''
}

func ExampleDigit0_match() {
	input := strings.NewReader("123abc")
	byteParser := ascii.Digit0()

	match, err := byteParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '123', Error: <nil>, Remainder: 'abc'
}

func ExampleDigit0_noMatch() {
	input := strings.NewReader("abc")
	byteParser := ascii.Digit0()

	match, err := byteParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: <nil>, Remainder: 'abc'
}

func ExampleDigit0_endOfFile() {
	input := strings.NewReader("")
	byteParser := ascii.Digit0()

	match, err := byteParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: <nil>, Remainder: ''
}
