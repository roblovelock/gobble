package numeric_test

import (
	"bytes"
	"fmt"
	"github.com/roblovelock/gobble/pkg/parser/numeric"
	"io"
)

func ExampleUInt8_match() {
	input := bytes.NewReader([]byte{1, 2, 3})
	numericParser := numeric.UInt8()

	match, err := numericParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: %v, Remainder: %v", match, err, remainder)

	// Output:
	// Match: 1, Error: <nil>, Remainder: [2 3]
}

func ExampleUInt8_endOfFile() {
	input := bytes.NewReader([]byte{})
	numericParser := numeric.UInt8()

	match, err := numericParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: %v", match, err, remainder)

	// Output:
	// Match: 0, Error: 'EOF', Remainder: []
}

func ExampleInt8_match() {
	input := bytes.NewReader([]byte{1, 2, 3})
	numericParser := numeric.Int8()

	match, err := numericParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: %v, Remainder: %v", match, err, remainder)

	// Output:
	// Match: 1, Error: <nil>, Remainder: [2 3]
}

func ExampleInt8_endOfFile() {
	input := bytes.NewReader([]byte{})
	numericParser := numeric.Int8()

	match, err := numericParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: %v", match, err, remainder)

	// Output:
	// Match: 0, Error: 'EOF', Remainder: []
}

func ExampleUInt16LE_match() {
	input := bytes.NewReader([]byte{0x01, 0x00, 3})
	numericParser := numeric.Uint16LE()

	match, err := numericParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: %v, Remainder: %v", match, err, remainder)

	// Output:
	// Match: 1, Error: <nil>, Remainder: [3]
}

func ExampleUInt16LE_endOfFile() {
	input := bytes.NewReader([]byte{})
	numericParser := numeric.Uint16LE()

	match, err := numericParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: %v", match, err, remainder)

	// Output:
	// Match: 0, Error: 'EOF', Remainder: []
}

func ExampleInt16LE_match() {
	input := bytes.NewReader([]byte{0x01, 0x00, 3})
	numericParser := numeric.Int16LE()

	match, err := numericParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: %v, Remainder: %v", match, err, remainder)

	// Output:
	// Match: 1, Error: <nil>, Remainder: [3]
}

func ExampleInt16LE_endOfFile() {
	input := bytes.NewReader([]byte{})
	numericParser := numeric.Int16LE()

	match, err := numericParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: %v", match, err, remainder)

	// Output:
	// Match: 0, Error: 'EOF', Remainder: []
}

func ExampleUInt16BE_match() {
	input := bytes.NewReader([]byte{0x00, 0x01, 3})
	numericParser := numeric.Uint16BE()

	match, err := numericParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: %v, Remainder: %v", match, err, remainder)

	// Output:
	// Match: 1, Error: <nil>, Remainder: [3]
}

func ExampleUInt16BE_endOfFile() {
	input := bytes.NewReader([]byte{})
	numericParser := numeric.Uint16BE()

	match, err := numericParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: %v", match, err, remainder)

	// Output:
	// Match: 0, Error: 'EOF', Remainder: []
}

func ExampleInt16BE_match() {
	input := bytes.NewReader([]byte{0x00, 0x01, 3})
	numericParser := numeric.Int16BE()

	match, err := numericParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: %v, Remainder: %v", match, err, remainder)

	// Output:
	// Match: 1, Error: <nil>, Remainder: [3]
}

func ExampleInt16BE_endOfFile() {
	input := bytes.NewReader([]byte{})
	numericParser := numeric.Int16BE()

	match, err := numericParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: %v", match, err, remainder)

	// Output:
	// Match: 0, Error: 'EOF', Remainder: []
}

func ExampleUInt32LE_match() {
	input := bytes.NewReader([]byte{0x01, 0x00, 0x00, 0x00, 3})
	numericParser := numeric.Uint32LE()

	match, err := numericParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: %v, Remainder: %v", match, err, remainder)

	// Output:
	// Match: 1, Error: <nil>, Remainder: [3]
}

func ExampleUInt32LE_endOfFile() {
	input := bytes.NewReader([]byte{})
	numericParser := numeric.Uint32LE()

	match, err := numericParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: %v", match, err, remainder)

	// Output:
	// Match: 0, Error: 'EOF', Remainder: []
}

func ExampleInt32LE_match() {
	input := bytes.NewReader([]byte{0x01, 0x00, 0x00, 0x00, 3})
	numericParser := numeric.Int32LE()

	match, err := numericParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: %v, Remainder: %v", match, err, remainder)

	// Output:
	// Match: 1, Error: <nil>, Remainder: [3]
}

func ExampleInt32LE_endOfFile() {
	input := bytes.NewReader([]byte{})
	numericParser := numeric.Int32LE()

	match, err := numericParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: %v", match, err, remainder)

	// Output:
	// Match: 0, Error: 'EOF', Remainder: []
}

func ExampleUInt32BE_match() {
	input := bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x01, 3})
	numericParser := numeric.Uint32BE()

	match, err := numericParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: %v, Remainder: %v", match, err, remainder)

	// Output:
	// Match: 1, Error: <nil>, Remainder: [3]
}

func ExampleUInt32BE_endOfFile() {
	input := bytes.NewReader([]byte{})
	numericParser := numeric.Uint32BE()

	match, err := numericParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: %v", match, err, remainder)

	// Output:
	// Match: 0, Error: 'EOF', Remainder: []
}

func ExampleInt32BE_match() {
	input := bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x01, 3})
	numericParser := numeric.Int32BE()

	match, err := numericParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: %v, Remainder: %v", match, err, remainder)

	// Output:
	// Match: 1, Error: <nil>, Remainder: [3]
}

func ExampleInt32BE_endOfFile() {
	input := bytes.NewReader([]byte{})
	numericParser := numeric.Int32BE()

	match, err := numericParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: %v", match, err, remainder)

	// Output:
	// Match: 0, Error: 'EOF', Remainder: []
}

func ExampleUInt64LE_match() {
	input := bytes.NewReader([]byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 3})
	numericParser := numeric.UInt64LE()

	match, err := numericParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: %v, Remainder: %v", match, err, remainder)

	// Output:
	// Match: 1, Error: <nil>, Remainder: [3]
}

func ExampleUInt64LE_endOfFile() {
	input := bytes.NewReader([]byte{})
	numericParser := numeric.UInt64LE()

	match, err := numericParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: %v", match, err, remainder)

	// Output:
	// Match: 0, Error: 'EOF', Remainder: []
}

func ExampleInt64LE_match() {
	input := bytes.NewReader([]byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 3})
	numericParser := numeric.Int64LE()

	match, err := numericParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: %v, Remainder: %v", match, err, remainder)

	// Output:
	// Match: 1, Error: <nil>, Remainder: [3]
}

func ExampleInt64LE_endOfFile() {
	input := bytes.NewReader([]byte{})
	numericParser := numeric.Int64LE()

	match, err := numericParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: %v", match, err, remainder)

	// Output:
	// Match: 0, Error: 'EOF', Remainder: []
}

func ExampleUInt64BE_match() {
	input := bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 3})
	numericParser := numeric.Uint64BE()

	match, err := numericParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: %v, Remainder: %v", match, err, remainder)

	// Output:
	// Match: 1, Error: <nil>, Remainder: [3]
}

func ExampleUInt64BE_endOfFile() {
	input := bytes.NewReader([]byte{})
	numericParser := numeric.Uint64BE()

	match, err := numericParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: %v", match, err, remainder)

	// Output:
	// Match: 0, Error: 'EOF', Remainder: []
}

func ExampleInt64BE_match() {
	input := bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 3})
	numericParser := numeric.Int64BE()

	match, err := numericParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: %v, Remainder: %v", match, err, remainder)

	// Output:
	// Match: 1, Error: <nil>, Remainder: [3]
}

func ExampleInt64BE_endOfFile() {
	input := bytes.NewReader([]byte{})
	numericParser := numeric.Int64BE()

	match, err := numericParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: %v", match, err, remainder)

	// Output:
	// Match: 0, Error: 'EOF', Remainder: []
}
