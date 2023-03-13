package runes_test

import (
	"fmt"
	"github.com/roblovelock/gobble/pkg/parser/runes"
	"io"
	"strings"
	"unicode"
)

func ExampleTakeWhileMinMax_match() {
	input := strings.NewReader("abc")
	runeParser := runes.TakeWhileMinMax(1, 2, unicode.IsLetter)

	match, err := runeParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 'ab', Error: <nil>, Remainder: 'c'
}

func ExampleTakeWhileMinMax_noMatch() {
	input := strings.NewReader("abc")
	runeParser := runes.TakeWhileMinMax(1, 2, unicode.IsDigit)

	match, err := runeParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: '', Error: 'not matched', Remainder: 'abc'
}

func ExampleTakeWhileMinMax_endOfFile() {
	input := strings.NewReader("")
	runeParser := runes.TakeWhileMinMax(1, 2, unicode.IsDigit)

	match, err := runeParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: '', Error: 'EOF', Remainder: ''
}

func ExampleTakeWhile1_match() {
	input := strings.NewReader("abc123")
	runeParser := runes.TakeWhile1(unicode.IsLetter)

	match, err := runeParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 'abc', Error: <nil>, Remainder: '123'
}

func ExampleTakeWhile1_noMatch() {
	input := strings.NewReader("abc")
	runeParser := runes.TakeWhile1(unicode.IsDigit)

	match, err := runeParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: '', Error: 'not matched', Remainder: 'abc'
}

func ExampleTakeWhile1_endOfFile() {
	input := strings.NewReader("")
	runeParser := runes.TakeWhile1(unicode.IsDigit)

	match, err := runeParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: '', Error: 'EOF', Remainder: ''
}

func ExampleTakeWhile_match() {
	input := strings.NewReader("abc123")
	runeParser := runes.TakeWhile(unicode.IsLetter)

	match, err := runeParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 'abc', Error: <nil>, Remainder: '123'
}

func ExampleTakeWhile_noMatch() {
	input := strings.NewReader("abc")
	runeParser := runes.TakeWhile(unicode.IsDigit)

	match, err := runeParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: '', Error: <nil>, Remainder: 'abc'
}

func ExampleTakeWhile_endOfFile() {
	input := strings.NewReader("")
	runeParser := runes.TakeWhile(unicode.IsDigit)

	match, err := runeParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: '', Error: <nil>, Remainder: ''
}
