package bytes_test

import (
	bytesio "bytes"
	"fmt"
	"github.com/roblovelock/gobble/pkg/parser/bytes"
	"io"
)

func ExampleOne_match() {
	input := bytesio.NewReader([]byte{1, 2, 3})
	bytesParser := bytes.One()

	match, err := bytesParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: %v, Remainder: %v", match, err, remainder)

	// Output:
	// Match: 1, Error: <nil>, Remainder: [2 3]
}

func ExampleOne_endOfFile() {
	input := bytesio.NewReader([]byte{})
	bytesParser := bytes.One()

	match, err := bytesParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: %v", match, err, remainder)

	// Output:
	// Match: 0, Error: 'EOF', Remainder: []
}
