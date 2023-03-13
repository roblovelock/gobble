package ascii_test

import (
	"fmt"
	"github.com/roblovelock/gobble/pkg/parser/ascii"
	"io"
	"strings"
)

func ExampleUInt8_match() {
	input := strings.NewReader("123abc")
	byteParser := ascii.UInt8()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: %v, Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 123, Error: <nil>, Remainder: 'abc'
}

func ExampleUInt8_noMatch() {
	input := strings.NewReader("abc")
	byteParser := ascii.UInt8()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'not matched', Remainder: 'abc'
}

func ExampleUInt8_overflow() {
	input := strings.NewReader("1234a")
	byteParser := ascii.UInt8()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'overflow', Remainder: '1234a'
}

func ExampleUInt8_endOfFile() {
	input := strings.NewReader("")
	byteParser := ascii.UInt8()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'EOF', Remainder: ''
}

func ExampleUInt16_match() {
	input := strings.NewReader("123abc")
	byteParser := ascii.UInt16()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: %v, Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 123, Error: <nil>, Remainder: 'abc'
}

func ExampleUInt16_noMatch() {
	input := strings.NewReader("abc")
	byteParser := ascii.UInt16()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'not matched', Remainder: 'abc'
}

func ExampleUInt16_overflow() {
	input := strings.NewReader("65536a")
	byteParser := ascii.UInt16()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'overflow', Remainder: '65536a'
}

func ExampleUInt16_endOfFile() {
	input := strings.NewReader("")
	byteParser := ascii.UInt16()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'EOF', Remainder: ''
}

func ExampleUInt32_match() {
	input := strings.NewReader("123abc")
	byteParser := ascii.UInt32()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: %v, Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 123, Error: <nil>, Remainder: 'abc'
}

func ExampleUInt32_noMatch() {
	input := strings.NewReader("abc")
	byteParser := ascii.UInt32()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'not matched', Remainder: 'abc'
}

func ExampleUInt32_overflow() {
	input := strings.NewReader("42949672950")
	byteParser := ascii.UInt32()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'overflow', Remainder: '42949672950'
}

func ExampleUInt32_endOfFile() {
	input := strings.NewReader("")
	byteParser := ascii.UInt32()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'EOF', Remainder: ''
}

func ExampleUInt64_match() {
	input := strings.NewReader("123abc")
	byteParser := ascii.UInt64()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: %v, Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 123, Error: <nil>, Remainder: 'abc'
}

func ExampleUInt64_noMatch() {
	input := strings.NewReader("abc")
	byteParser := ascii.UInt64()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'not matched', Remainder: 'abc'
}

func ExampleUInt64_overflow() {
	input := strings.NewReader("18446744073709551616")
	byteParser := ascii.UInt64()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'overflow', Remainder: '18446744073709551616'
}

func ExampleUInt64_endOfFile() {
	input := strings.NewReader("")
	byteParser := ascii.UInt64()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'EOF', Remainder: ''
}

func ExampleInt8_match() {
	input := strings.NewReader("123abc")
	byteParser := ascii.Int8()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: %v, Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 123, Error: <nil>, Remainder: 'abc'
}

func ExampleInt8_noMatch() {
	input := strings.NewReader("abc")
	byteParser := ascii.Int8()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'not matched', Remainder: 'abc'
}

func ExampleInt8_overflow() {
	input := strings.NewReader("1234a")
	byteParser := ascii.Int8()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'overflow', Remainder: '1234a'
}

func ExampleInt8_endOfFile() {
	input := strings.NewReader("")
	byteParser := ascii.Int8()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'EOF', Remainder: ''
}

func ExampleInt16_match() {
	input := strings.NewReader("123abc")
	byteParser := ascii.Int16()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: %v, Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 123, Error: <nil>, Remainder: 'abc'
}

func ExampleInt16_noMatch() {
	input := strings.NewReader("abc")
	byteParser := ascii.Int16()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'not matched', Remainder: 'abc'
}

func ExampleInt16_overflow() {
	input := strings.NewReader("65536a")
	byteParser := ascii.Int16()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'overflow', Remainder: '65536a'
}

func ExampleInt16_endOfFile() {
	input := strings.NewReader("")
	byteParser := ascii.Int16()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'EOF', Remainder: ''
}

func ExampleInt32_match() {
	input := strings.NewReader("123abc")
	byteParser := ascii.Int32()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: %v, Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 123, Error: <nil>, Remainder: 'abc'
}

func ExampleInt32_noMatch() {
	input := strings.NewReader("abc")
	byteParser := ascii.Int32()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'not matched', Remainder: 'abc'
}

func ExampleInt32_overflow() {
	input := strings.NewReader("42949672950")
	byteParser := ascii.Int32()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'overflow', Remainder: '42949672950'
}

func ExampleInt32_endOfFile() {
	input := strings.NewReader("")
	byteParser := ascii.Int32()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'EOF', Remainder: ''
}

func ExampleInt64_match() {
	input := strings.NewReader("123abc")
	byteParser := ascii.Int64()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: %v, Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 123, Error: <nil>, Remainder: 'abc'
}

func ExampleInt64_noMatch() {
	input := strings.NewReader("abc")
	byteParser := ascii.Int64()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'not matched', Remainder: 'abc'
}

func ExampleInt64_overflow() {
	input := strings.NewReader("18446744073709551616")
	byteParser := ascii.Int64()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'overflow', Remainder: '18446744073709551616'
}

func ExampleInt64_endOfFile() {
	input := strings.NewReader("")
	byteParser := ascii.Int64()

	match, err := byteParser.Parse(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'EOF', Remainder: ''
}
