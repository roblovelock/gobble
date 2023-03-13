package runes_test

import (
	"bytes"
	"fmt"
	"github.com/roblovelock/gobble/pkg/parser/runes"
	"io"
	"strings"
)

func ExampleOne_match() {
	input := strings.NewReader("𒀀𒀀")
	numericParser := runes.One()

	match, err := numericParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, remainder)

	// Output:
	// Match: '𒀀', Error: <nil>, Remainder: '𒀀'
}

func ExampleOne_endOfFile() {
	input := bytes.NewReader([]byte{})
	numericParser := runes.One()

	match, err := numericParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: %v", match, err, remainder)

	// Output:
	// Match: 0, Error: 'EOF', Remainder: []
}
