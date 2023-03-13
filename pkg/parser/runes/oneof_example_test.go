package runes_test

import (
	"fmt"
	"github.com/roblovelock/gobble/pkg/parser/runes"
	"io"
	"strings"
)

func ExampleOneOf_match() {
	input := strings.NewReader("𒀀a𒀀")
	runeParser := runes.OneOf('𒀀', 'a')

	match, err := runeParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '𒀀', Error: <nil>, Remainder: 'a𒀀'
}

func ExampleOneOf_noMatch() {
	input := strings.NewReader("123")
	runeParser := runes.OneOf('𒀀', 'a')

	match, err := runeParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'not matched', Remainder: '123'
}

func ExampleOneOf_endOfFile() {
	input := strings.NewReader("")
	runeParser := runes.OneOf('𒀀', 'a')

	match, err := runeParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'EOF', Remainder: ''
}

func ExampleOneOf1_match() {
	input := strings.NewReader("𒀀a𒀀123")
	runeParser := runes.OneOf1('𒀀', 'a')

	match, err := runeParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '𒀀a𒀀', Error: <nil>, Remainder: '123'
}

func ExampleOneOf1_noMatch() {
	input := strings.NewReader("123")
	runeParser := runes.OneOf1('𒀀', 'a')

	match, err := runeParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: '%v', Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: 'not matched', Remainder: '123'
}

func ExampleOneOf1_endOfFile() {
	input := strings.NewReader("")
	runeParser := runes.OneOf1('𒀀', 'a')

	match, err := runeParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: '%v', Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: 'EOF', Remainder: ''
}

func ExampleOneOf0_match() {
	input := strings.NewReader("𒀀a𒀀123")
	runeParser := runes.OneOf0('𒀀', 'a')

	match, err := runeParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '𒀀a𒀀', Error: <nil>, Remainder: '123'
}

func ExampleOneOf0_noMatch() {
	input := strings.NewReader("123")
	runeParser := runes.OneOf0('𒀀', 'a')

	match, err := runeParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: <nil>, Remainder: '123'
}

func ExampleOneOf0_endOfFile() {
	input := strings.NewReader("")
	runeParser := runes.OneOf0('𒀀', 'a')

	match, err := runeParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: <nil>, Remainder: ''
}
