package numeric_test

import (
	"bytes"
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/roblovelock/gobble/pkg/parser/numeric"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"testing"
)

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
			s, err := p.Parse(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)
		})
	}
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
			s, err := p.Parse(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)
		})
	}
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
			p := numeric.Uint16LE()
			s, err := p.Parse(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)
		})
	}
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
			s, err := p.Parse(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)
		})
	}
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
			p := numeric.Uint16BE()
			s, err := p.Parse(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)
		})
	}
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
			s, err := p.Parse(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)
		})
	}
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
			p := numeric.Uint32LE()
			s, err := p.Parse(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)
		})
	}
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
			s, err := p.Parse(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)
		})
	}
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
			p := numeric.Uint32BE()
			s, err := p.Parse(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)
		})
	}
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
			s, err := p.Parse(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)
		})
	}
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
			s, err := p.Parse(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)
		})
	}
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
			s, err := p.Parse(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)
		})
	}
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
			p := numeric.Uint64BE()
			s, err := p.Parse(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)
		})
	}
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
			s, err := p.Parse(tt.args.input)

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
