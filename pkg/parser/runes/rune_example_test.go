package runes_test

import (
	"fmt"
	"github.com/roblovelock/gobble/pkg/parser/runes"
	"io"
	"strings"
)

func ExampleRune_match() {
	input := strings.NewReader("𒀀a𒀀")
	byteParser := runes.Rune('𒀀')

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '𒀀', Error: <nil>, Remainder: 'a𒀀'
}

func ExampleRune_noMatch() {
	input := strings.NewReader("𒀀a𒀀")
	byteParser := runes.Rune('𒀁')

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'not matched', Remainder: '𒀀a𒀀'
}

func ExampleRune_endOfFile() {
	input := strings.NewReader("")
	byteParser := runes.Rune('𒀁')

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'EOF', Remainder: ''
}
