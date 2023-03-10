package runes_test

import (
	"fmt"
	"github.com/roblovelock/gobble/pkg/parser/runes"
	"io"
	"strings"
)

func ExampleRune_match() {
	input := strings.NewReader("ğağ")
	byteParser := runes.Rune('ğ')

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: 'ğ', Error: <nil>, Remainder: 'ağ'
}

func ExampleRune_noMatch() {
	input := strings.NewReader("ğağ")
	byteParser := runes.Rune('ğ')

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'not matched', Remainder: 'ğağ'
}

func ExampleRune_endOfFile() {
	input := strings.NewReader("")
	byteParser := runes.Rune('ğ')

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'EOF', Remainder: ''
}
