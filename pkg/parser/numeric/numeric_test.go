package numeric_test

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gobble/pkg/parser"
	"gobble/pkg/parser/numeric"
	"io"
	"testing"
)

func ExampleUInt8_match() {
	input := bytes.NewReader([]byte{1, 2, 3})
	numericParser := numeric.UInt8()

	match, err := numericParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: %v, Remainder: %v", match, err, remainder)

	// Output:
	// Match: 1, Error: <nil>, Remainder: [2 3]
}

func ExampleUInt8_endOfFile() {
	input := bytes.NewReader([]byte{})
	numericParser := numeric.UInt8()

	match, err := numericParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: %v", match, err, remainder)

	// Output:
	// Match: 0, Error: 'EOF', Remainder: []
}

func TestUInt8(t *testing.T) {
	type args struct {
		input parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  uint8
		wantRemain []byte
		wantErr    error
	}{
		{
			name:    "empty input => EOF",
			args:    args{input: bytes.NewReader([]byte{})},
			wantErr: io.EOF, wantRemain: []byte{},
		},
		{
			name:       "uint8 => match",
			args:       args{input: bytes.NewReader([]byte{1, 2, 3})},
			wantMatch:  1,
			wantRemain: []byte{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := numeric.UInt8()
			s, err := p(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)
		})
	}
}

func ExampleInt8_match() {
	input := bytes.NewReader([]byte{1, 2, 3})
	numericParser := numeric.Int8()

	match, err := numericParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: %v, Remainder: %v", match, err, remainder)

	// Output:
	// Match: 1, Error: <nil>, Remainder: [2 3]
}

func ExampleInt8_endOfFile() {
	input := bytes.NewReader([]byte{})
	numericParser := numeric.Int8()

	match, err := numericParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: %v", match, err, remainder)

	// Output:
	// Match: 0, Error: 'EOF', Remainder: []
}

func TestInt8(t *testing.T) {
	type args struct {
		input parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  int8
		wantRemain []byte
		wantErr    error
	}{
		{
			name:    "empty input => EOF",
			args:    args{input: bytes.NewReader([]byte{})},
			wantErr: io.EOF, wantRemain: []byte{},
		},
		{
			name:       "positive int8 => match",
			args:       args{input: bytes.NewReader([]byte{1, 2, 3})},
			wantMatch:  1,
			wantRemain: []byte{2, 3},
		},
		{
			name:       "negative int8 => match",
			args:       args{input: bytes.NewReader([]byte{Int8ToByte(-1), 2, 3})},
			wantMatch:  -1,
			wantRemain: []byte{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := numeric.Int8()
			s, err := p(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)
		})
	}
}

func ExampleUInt16LE_match() {
	input := bytes.NewReader([]byte{0x01, 0x00, 3})
	numericParser := numeric.UInt16LE()

	match, err := numericParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: %v, Remainder: %v", match, err, remainder)

	// Output:
	// Match: 1, Error: <nil>, Remainder: [3]
}

func ExampleUInt16LE_endOfFile() {
	input := bytes.NewReader([]byte{})
	numericParser := numeric.UInt16LE()

	match, err := numericParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: %v", match, err, remainder)

	// Output:
	// Match: 0, Error: 'EOF', Remainder: []
}

func TestUInt16LE(t *testing.T) {
	type args struct {
		input parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  uint16
		wantRemain []byte
		wantErr    error
	}{
		{
			name:    "empty input => EOF",
			args:    args{input: bytes.NewReader([]byte{})},
			wantErr: io.EOF, wantRemain: []byte{},
		},
		{
			name:       "uint16 => match",
			args:       args{input: bytes.NewReader([]byte{0xFE, 0xFF, 3})},
			wantMatch:  0xfffe,
			wantRemain: []byte{3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := numeric.UInt16LE()
			s, err := p(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)
		})
	}
}

func ExampleInt16LE_match() {
	input := bytes.NewReader([]byte{0x01, 0x00, 3})
	numericParser := numeric.Int16LE()

	match, err := numericParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: %v, Remainder: %v", match, err, remainder)

	// Output:
	// Match: 1, Error: <nil>, Remainder: [3]
}

func ExampleInt16LE_endOfFile() {
	input := bytes.NewReader([]byte{})
	numericParser := numeric.Int16LE()

	match, err := numericParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: %v", match, err, remainder)

	// Output:
	// Match: 0, Error: 'EOF', Remainder: []
}

func TestInt16LE(t *testing.T) {
	type args struct {
		input parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  int16
		wantRemain []byte
		wantErr    error
	}{
		{
			name:       "empty input => EOF",
			args:       args{input: bytes.NewReader([]byte{})},
			wantErr:    io.EOF,
			wantRemain: []byte{},
		},
		{
			name:       "short input => EOF",
			args:       args{input: bytes.NewReader([]byte{1})},
			wantErr:    io.EOF,
			wantRemain: []byte{1},
		},
		{
			name:       "positive int16 => match",
			args:       args{input: bytes.NewReader([]byte{0x01, 0x00, 2, 3})},
			wantMatch:  1,
			wantRemain: []byte{2, 3},
		},
		{
			name:       "negative int16 => match",
			args:       args{input: bytes.NewReader([]byte{0xFE, 0xFF, 2, 3})},
			wantMatch:  -2,
			wantRemain: []byte{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := numeric.Int16LE()
			s, err := p(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)
		})
	}
}

func ExampleUInt16BE_match() {
	input := bytes.NewReader([]byte{0x00, 0x01, 3})
	numericParser := numeric.UInt16BE()

	match, err := numericParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: %v, Remainder: %v", match, err, remainder)

	// Output:
	// Match: 1, Error: <nil>, Remainder: [3]
}

func ExampleUInt16BE_endOfFile() {
	input := bytes.NewReader([]byte{})
	numericParser := numeric.UInt16BE()

	match, err := numericParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: %v", match, err, remainder)

	// Output:
	// Match: 0, Error: 'EOF', Remainder: []
}

func TestUInt16BE(t *testing.T) {
	type args struct {
		input parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  uint16
		wantRemain []byte
		wantErr    error
	}{
		{
			name:    "empty input => EOF",
			args:    args{input: bytes.NewReader([]byte{})},
			wantErr: io.EOF, wantRemain: []byte{},
		},
		{
			name:       "uint16 => match",
			args:       args{input: bytes.NewReader([]byte{0xFF, 0xFE, 3})},
			wantMatch:  0xfffe,
			wantRemain: []byte{3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := numeric.UInt16BE()
			s, err := p(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)
		})
	}
}

func ExampleInt16BE_match() {
	input := bytes.NewReader([]byte{0x00, 0x01, 3})
	numericParser := numeric.Int16BE()

	match, err := numericParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: %v, Remainder: %v", match, err, remainder)

	// Output:
	// Match: 1, Error: <nil>, Remainder: [3]
}

func ExampleInt16BE_endOfFile() {
	input := bytes.NewReader([]byte{})
	numericParser := numeric.Int16BE()

	match, err := numericParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: %v", match, err, remainder)

	// Output:
	// Match: 0, Error: 'EOF', Remainder: []
}

func TestInt16BE(t *testing.T) {
	type args struct {
		input parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  int16
		wantRemain []byte
		wantErr    error
	}{
		{
			name:       "empty input => EOF",
			args:       args{input: bytes.NewReader([]byte{})},
			wantErr:    io.EOF,
			wantRemain: []byte{},
		},
		{
			name:       "short input => EOF",
			args:       args{input: bytes.NewReader([]byte{1})},
			wantErr:    io.EOF,
			wantRemain: []byte{1},
		},
		{
			name:       "positive int16 => match",
			args:       args{input: bytes.NewReader([]byte{0x00, 0x01, 2, 3})},
			wantMatch:  1,
			wantRemain: []byte{2, 3},
		},
		{
			name:       "negative int16 => match",
			args:       args{input: bytes.NewReader([]byte{0xFF, 0xFE, 2, 3})},
			wantMatch:  -2,
			wantRemain: []byte{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := numeric.Int16BE()
			s, err := p(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)
		})
	}
}

func ExampleUInt32LE_match() {
	input := bytes.NewReader([]byte{0x01, 0x00, 0x00, 0x00, 3})
	numericParser := numeric.UInt32LE()

	match, err := numericParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: %v, Remainder: %v", match, err, remainder)

	// Output:
	// Match: 1, Error: <nil>, Remainder: [3]
}

func ExampleUInt32LE_endOfFile() {
	input := bytes.NewReader([]byte{})
	numericParser := numeric.UInt32LE()

	match, err := numericParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: %v", match, err, remainder)

	// Output:
	// Match: 0, Error: 'EOF', Remainder: []
}

func TestUInt32LE(t *testing.T) {
	type args struct {
		input parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  uint32
		wantRemain []byte
		wantErr    error
	}{
		{
			name:    "empty input => EOF",
			args:    args{input: bytes.NewReader([]byte{})},
			wantErr: io.EOF, wantRemain: []byte{},
		},
		{
			name:       "uint32 => match",
			args:       args{input: bytes.NewReader([]byte{0xFE, 0xFF, 0xFF, 0xFF, 3})},
			wantMatch:  0xfffffffe,
			wantRemain: []byte{3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := numeric.UInt32LE()
			s, err := p(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)
		})
	}
}

func ExampleInt32LE_match() {
	input := bytes.NewReader([]byte{0x01, 0x00, 0x00, 0x00, 3})
	numericParser := numeric.Int32LE()

	match, err := numericParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: %v, Remainder: %v", match, err, remainder)

	// Output:
	// Match: 1, Error: <nil>, Remainder: [3]
}

func ExampleInt32LE_endOfFile() {
	input := bytes.NewReader([]byte{})
	numericParser := numeric.Int32LE()

	match, err := numericParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: %v", match, err, remainder)

	// Output:
	// Match: 0, Error: 'EOF', Remainder: []
}

func TestInt32LE(t *testing.T) {
	type args struct {
		input parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  int32
		wantRemain []byte
		wantErr    error
	}{
		{
			name:       "empty input => EOF",
			args:       args{input: bytes.NewReader([]byte{})},
			wantErr:    io.EOF,
			wantRemain: []byte{},
		},
		{
			name:       "short input => EOF",
			args:       args{input: bytes.NewReader([]byte{1})},
			wantErr:    io.EOF,
			wantRemain: []byte{1},
		},
		{
			name:       "positive int32 => match",
			args:       args{input: bytes.NewReader([]byte{0x01, 0x00, 0x00, 0x00, 2, 3})},
			wantMatch:  1,
			wantRemain: []byte{2, 3},
		},
		{
			name:       "negative int32 => match",
			args:       args{input: bytes.NewReader([]byte{0xFE, 0xFF, 0xFF, 0xFF, 2, 3})},
			wantMatch:  -2,
			wantRemain: []byte{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := numeric.Int32LE()
			s, err := p(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)
		})
	}
}

func ExampleUInt32BE_match() {
	input := bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x01, 3})
	numericParser := numeric.UInt32BE()

	match, err := numericParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: %v, Remainder: %v", match, err, remainder)

	// Output:
	// Match: 1, Error: <nil>, Remainder: [3]
}

func ExampleUInt32BE_endOfFile() {
	input := bytes.NewReader([]byte{})
	numericParser := numeric.UInt32BE()

	match, err := numericParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: %v", match, err, remainder)

	// Output:
	// Match: 0, Error: 'EOF', Remainder: []
}

func TestUInt32BE(t *testing.T) {
	type args struct {
		input parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  uint32
		wantRemain []byte
		wantErr    error
	}{
		{
			name:    "empty input => EOF",
			args:    args{input: bytes.NewReader([]byte{})},
			wantErr: io.EOF, wantRemain: []byte{},
		},
		{
			name:       "uint32 => match",
			args:       args{input: bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0xFE, 3})},
			wantMatch:  0xfffffffe,
			wantRemain: []byte{3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := numeric.UInt32BE()
			s, err := p(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)
		})
	}
}

func ExampleInt32BE_match() {
	input := bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x01, 3})
	numericParser := numeric.Int32BE()

	match, err := numericParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: %v, Remainder: %v", match, err, remainder)

	// Output:
	// Match: 1, Error: <nil>, Remainder: [3]
}

func ExampleInt32BE_endOfFile() {
	input := bytes.NewReader([]byte{})
	numericParser := numeric.Int32BE()

	match, err := numericParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: %v", match, err, remainder)

	// Output:
	// Match: 0, Error: 'EOF', Remainder: []
}

func TestInt32BE(t *testing.T) {
	type args struct {
		input parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  int32
		wantRemain []byte
		wantErr    error
	}{
		{
			name:       "empty input => EOF",
			args:       args{input: bytes.NewReader([]byte{})},
			wantErr:    io.EOF,
			wantRemain: []byte{},
		},
		{
			name:       "short input => EOF",
			args:       args{input: bytes.NewReader([]byte{1})},
			wantErr:    io.EOF,
			wantRemain: []byte{1},
		},
		{
			name:       "positive int32 => match",
			args:       args{input: bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x01, 2, 3})},
			wantMatch:  1,
			wantRemain: []byte{2, 3},
		},
		{
			name:       "negative int32 => match",
			args:       args{input: bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0xFE, 2, 3})},
			wantMatch:  -2,
			wantRemain: []byte{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := numeric.Int32BE()
			s, err := p(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)
		})
	}
}

func ExampleUInt64LE_match() {
	input := bytes.NewReader([]byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 3})
	numericParser := numeric.UInt64LE()

	match, err := numericParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: %v, Remainder: %v", match, err, remainder)

	// Output:
	// Match: 1, Error: <nil>, Remainder: [3]
}

func ExampleUInt64LE_endOfFile() {
	input := bytes.NewReader([]byte{})
	numericParser := numeric.UInt64LE()

	match, err := numericParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: %v", match, err, remainder)

	// Output:
	// Match: 0, Error: 'EOF', Remainder: []
}

func TestUInt64LE(t *testing.T) {
	type args struct {
		input parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  uint64
		wantRemain []byte
		wantErr    error
	}{
		{
			name:    "empty input => EOF",
			args:    args{input: bytes.NewReader([]byte{})},
			wantErr: io.EOF, wantRemain: []byte{},
		},
		{
			name:       "uint64 => match",
			args:       args{input: bytes.NewReader([]byte{0xFE, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 3})},
			wantMatch:  0xfffffffffffffffe,
			wantRemain: []byte{3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := numeric.UInt64LE()
			s, err := p(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)
		})
	}
}

func ExampleInt64LE_match() {
	input := bytes.NewReader([]byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 3})
	numericParser := numeric.Int64LE()

	match, err := numericParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: %v, Remainder: %v", match, err, remainder)

	// Output:
	// Match: 1, Error: <nil>, Remainder: [3]
}

func ExampleInt64LE_endOfFile() {
	input := bytes.NewReader([]byte{})
	numericParser := numeric.Int64LE()

	match, err := numericParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: %v", match, err, remainder)

	// Output:
	// Match: 0, Error: 'EOF', Remainder: []
}

func TestInt64LE(t *testing.T) {
	type args struct {
		input parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  int64
		wantRemain []byte
		wantErr    error
	}{
		{
			name:       "empty input => EOF",
			args:       args{input: bytes.NewReader([]byte{})},
			wantErr:    io.EOF,
			wantRemain: []byte{},
		},
		{
			name:       "short input => EOF",
			args:       args{input: bytes.NewReader([]byte{1})},
			wantErr:    io.EOF,
			wantRemain: []byte{1},
		},
		{
			name:       "positive int64 => match",
			args:       args{input: bytes.NewReader([]byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 2, 3})},
			wantMatch:  1,
			wantRemain: []byte{2, 3},
		},
		{
			name:       "negative int64 => match",
			args:       args{input: bytes.NewReader([]byte{0xFE, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 2, 3})},
			wantMatch:  -2,
			wantRemain: []byte{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := numeric.Int64LE()
			s, err := p(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)
		})
	}
}

func ExampleUInt64BE_match() {
	input := bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 3})
	numericParser := numeric.UInt64BE()

	match, err := numericParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: %v, Remainder: %v", match, err, remainder)

	// Output:
	// Match: 1, Error: <nil>, Remainder: [3]
}

func ExampleUInt64BE_endOfFile() {
	input := bytes.NewReader([]byte{})
	numericParser := numeric.UInt64BE()

	match, err := numericParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: %v", match, err, remainder)

	// Output:
	// Match: 0, Error: 'EOF', Remainder: []
}

func TestUInt64BE(t *testing.T) {
	type args struct {
		input parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  uint64
		wantRemain []byte
		wantErr    error
	}{
		{
			name:    "empty input => EOF",
			args:    args{input: bytes.NewReader([]byte{})},
			wantErr: io.EOF, wantRemain: []byte{},
		},
		{
			name:       "uint64 => match",
			args:       args{input: bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFE, 3})},
			wantMatch:  0xfffffffffffffffe,
			wantRemain: []byte{3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := numeric.UInt64BE()
			s, err := p(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)
		})
	}
}

func ExampleInt64BE_match() {
	input := bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 3})
	numericParser := numeric.Int64BE()

	match, err := numericParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: %v, Remainder: %v", match, err, remainder)

	// Output:
	// Match: 1, Error: <nil>, Remainder: [3]
}

func ExampleInt64BE_endOfFile() {
	input := bytes.NewReader([]byte{})
	numericParser := numeric.Int64BE()

	match, err := numericParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: %v", match, err, remainder)

	// Output:
	// Match: 0, Error: 'EOF', Remainder: []
}

func TestInt64BE(t *testing.T) {
	type args struct {
		input parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  int64
		wantRemain []byte
		wantErr    error
	}{
		{
			name:       "empty input => EOF",
			args:       args{input: bytes.NewReader([]byte{})},
			wantErr:    io.EOF,
			wantRemain: []byte{},
		},
		{
			name:       "short input => EOF",
			args:       args{input: bytes.NewReader([]byte{1})},
			wantErr:    io.EOF,
			wantRemain: []byte{1},
		},
		{
			name:       "positive int64 => match",
			args:       args{input: bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 2, 3})},
			wantMatch:  1,
			wantRemain: []byte{2, 3},
		},
		{
			name:       "negative int64 => match",
			args:       args{input: bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFE, 2, 3})},
			wantMatch:  -2,
			wantRemain: []byte{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := numeric.Int64BE()
			s, err := p(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)
		})
	}
}

func Int8ToByte(i int8) byte {
	return byte(i)
}
