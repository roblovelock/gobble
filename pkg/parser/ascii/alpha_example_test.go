package ascii_test

import (
	"fmt"
	"github.com/roblovelock/gobble/pkg/parser/ascii"
	"io"
	"strings"
)

func ExampleAlpha_match() {
	input := strings.NewReader("abc")
	byteParser := ascii.Alpha()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: 'a', Error: <nil>, Remainder: 'bc'
}

func ExampleAlpha_noMatch() {
	input := strings.NewReader("123")
	byteParser := ascii.Alpha()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'not matched', Remainder: '123'
}

func ExampleAlpha_endOfFile() {
	input := strings.NewReader("")
	byteParser := ascii.Alpha()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'EOF', Remainder: ''
}

func ExampleAlpha1_match() {
	input := strings.NewReader("abc123")
	byteParser := ascii.Alpha1()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: 'abc', Error: <nil>, Remainder: '123'
}

func ExampleAlpha1_noMatch() {
	input := strings.NewReader("123")
	byteParser := ascii.Alpha1()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: '%v', Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: 'not matched', Remainder: '123'
}

func ExampleAlpha1_endOfFile() {
	input := strings.NewReader("")
	byteParser := ascii.Alpha1()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: '%v', Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: 'EOF', Remainder: ''
}

func ExampleAlpha0_match() {
	input := strings.NewReader("abc123")
	byteParser := ascii.Alpha0()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: 'abc', Error: <nil>, Remainder: '123'
}

func ExampleAlpha0_noMatch() {
	input := strings.NewReader("123")
	byteParser := ascii.Alpha0()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: <nil>, Remainder: '123'
}

func ExampleAlpha0_endOfFile() {
	input := strings.NewReader("")
	byteParser := ascii.Alpha0()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: <nil>, Remainder: ''
}
