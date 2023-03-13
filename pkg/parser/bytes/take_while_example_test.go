package bytes_test

import (
	"fmt"
	"github.com/roblovelock/gobble/pkg/parser/ascii"
	"github.com/roblovelock/gobble/pkg/parser/bytes"
	"io"
	"strings"
)

func ExampleTakeWhileMinMax_match() {
	input := strings.NewReader("abc")
	byteParser := bytes.TakeWhileMinMax(1, 2, ascii.IsLetter)

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: 'ab', Error: <nil>, Remainder: 'c'
}

func ExampleTakeWhileMinMax_noMatch() {
	input := strings.NewReader("abc")
	byteParser := bytes.TakeWhileMinMax(1, 2, ascii.IsDigit)

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: '%v', Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: 'not matched', Remainder: 'abc'
}

func ExampleTakeWhileMinMax_endOfFile() {
	input := strings.NewReader("")
	byteParser := bytes.TakeWhileMinMax(1, 2, ascii.IsDigit)

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: '%v', Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: 'EOF', Remainder: ''
}

func ExampleTakeWhile1_match() {
	input := strings.NewReader("abc123")
	byteParser := bytes.TakeWhile1(ascii.IsLetter)

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: 'abc', Error: <nil>, Remainder: '123'
}

func ExampleTakeWhile1_noMatch() {
	input := strings.NewReader("abc")
	byteParser := bytes.TakeWhile1(ascii.IsDigit)

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: '%v', Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: 'not matched', Remainder: 'abc'
}

func ExampleTakeWhile1_endOfFile() {
	input := strings.NewReader("")
	byteParser := bytes.TakeWhile1(ascii.IsDigit)

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: '%v', Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: 'EOF', Remainder: ''
}

func ExampleTakeWhile_match() {
	input := strings.NewReader("abc123")
	byteParser := bytes.TakeWhile(ascii.IsLetter)

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: 'abc', Error: <nil>, Remainder: '123'
}

func ExampleTakeWhile_noMatch() {
	input := strings.NewReader("abc")
	byteParser := bytes.TakeWhile(ascii.IsDigit)

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: <nil>, Remainder: 'abc'
}

func ExampleTakeWhile_endOfFile() {
	input := strings.NewReader("")
	byteParser := bytes.TakeWhile(ascii.IsDigit)

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: <nil>, Remainder: ''
}
