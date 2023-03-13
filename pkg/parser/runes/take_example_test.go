package runes_test

import (
	"fmt"
	"github.com/roblovelock/gobble/pkg/parser/runes"
	"io"
	"strings"
)

func ExampleTake_match() {
	input := strings.NewReader("𒀀a𒀀")
	runeParser := runes.Take(2)

	match, err := runeParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: '𒀀a', Error: <nil>, Remainder: '𒀀'
}

func ExampleTake_unexpectedEndOfFile() {
	input := strings.NewReader("𒀀a𒀀")
	runeParser := runes.Take(4)

	match, err := runeParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: '', Error: 'EOF', Remainder: '𒀀a𒀀'
}

func ExampleTake_endOfFile() {
	input := strings.NewReader("")
	runeParser := runes.Take(4)

	match, err := runeParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: '', Error: 'EOF', Remainder: ''
}
