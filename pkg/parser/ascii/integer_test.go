package ascii_test

import (
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/roblovelock/gobble/pkg/parser/ascii"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"math"
	"strings"
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
			args:    args{input: strings.NewReader("")},
			wantErr: io.EOF,
		},
		{
			name:       "non digit => no match",
			args:       args{input: strings.NewReader("a")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'a'},
		},
		{
			name:       "negative => no match",
			args:       args{input: strings.NewReader("-a")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'-', 'a'},
		},
		{
			name:       "positive => no match",
			args:       args{input: strings.NewReader("+a")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'+', 'a'},
		},
		{
			name:       "digit => match",
			args:       args{input: strings.NewReader("9")},
			wantMatch:  9,
			wantRemain: []byte{},
		},
		{
			name:       "positive digit => match",
			args:       args{input: strings.NewReader("+9")},
			wantMatch:  9,
			wantRemain: []byte{},
		},
		{
			name:       "min number => match",
			args:       args{input: strings.NewReader("0")},
			wantMatch:  0,
			wantRemain: []byte{},
		},
		{
			name:       "max number => match",
			args:       args{input: strings.NewReader("255")},
			wantMatch:  math.MaxUint8,
			wantRemain: []byte{},
		},
		{
			name:       "one over max number => match",
			args:       args{input: strings.NewReader("256")},
			wantErr:    ascii.ErrOverflow,
			wantRemain: []byte("256"),
		},
		{
			name:       "large number => overflow",
			args:       args{input: strings.NewReader("1000")},
			wantErr:    ascii.ErrOverflow,
			wantRemain: []byte("1000"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ascii.UInt8()
			s, err := p.Parse(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			if err != io.EOF {
				remain, err := io.ReadAll(tt.args.input)
				require.NoError(t, err)
				assert.Equal(t, tt.wantRemain, remain)
			}
		})
	}
}

func TestUInt16(t *testing.T) {
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
			args:    args{input: strings.NewReader("")},
			wantErr: io.EOF,
		},
		{
			name:       "non digit => no match",
			args:       args{input: strings.NewReader("a")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'a'},
		},
		{
			name:       "negative => no match",
			args:       args{input: strings.NewReader("-a")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'-', 'a'},
		},
		{
			name:       "positive => no match",
			args:       args{input: strings.NewReader("+a")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'+', 'a'},
		},
		{
			name:       "digit => match",
			args:       args{input: strings.NewReader("9")},
			wantMatch:  9,
			wantRemain: []byte{},
		},
		{
			name:       "positive digit => match",
			args:       args{input: strings.NewReader("+9")},
			wantMatch:  9,
			wantRemain: []byte{},
		},
		{
			name:       "min number => match",
			args:       args{input: strings.NewReader("0")},
			wantMatch:  0,
			wantRemain: []byte{},
		},
		{
			name:       "max number => match",
			args:       args{input: strings.NewReader("65535")},
			wantMatch:  math.MaxUint16,
			wantRemain: []byte{},
		},
		{
			name:       "one over max number => match",
			args:       args{input: strings.NewReader("65536")},
			wantErr:    ascii.ErrOverflow,
			wantRemain: []byte("65536"),
		},
		{
			name:       "large number => overflow",
			args:       args{input: strings.NewReader("100000")},
			wantErr:    ascii.ErrOverflow,
			wantRemain: []byte("100000"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ascii.UInt16()
			s, err := p.Parse(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			if err != io.EOF {
				remain, err := io.ReadAll(tt.args.input)
				require.NoError(t, err)
				assert.Equal(t, tt.wantRemain, remain)
			}
		})
	}
}

func TestUInt32(t *testing.T) {
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
			args:    args{input: strings.NewReader("")},
			wantErr: io.EOF,
		},
		{
			name:       "non digit => no match",
			args:       args{input: strings.NewReader("a")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'a'},
		},
		{
			name:       "negative => no match",
			args:       args{input: strings.NewReader("-a")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'-', 'a'},
		},
		{
			name:       "positive => no match",
			args:       args{input: strings.NewReader("+a")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'+', 'a'},
		},
		{
			name:       "digit => match",
			args:       args{input: strings.NewReader("9")},
			wantMatch:  9,
			wantRemain: []byte{},
		},
		{
			name:       "positive digit => match",
			args:       args{input: strings.NewReader("+9")},
			wantMatch:  9,
			wantRemain: []byte{},
		},
		{
			name:       "min number => match",
			args:       args{input: strings.NewReader("0")},
			wantMatch:  0,
			wantRemain: []byte{},
		},
		{
			name:       "max number => match",
			args:       args{input: strings.NewReader("4294967295")},
			wantMatch:  math.MaxUint32,
			wantRemain: []byte{},
		},
		{
			name:       "one over max number => match",
			args:       args{input: strings.NewReader("4294967296")},
			wantErr:    ascii.ErrOverflow,
			wantRemain: []byte("4294967296"),
		},
		{
			name:       "large number => overflow",
			args:       args{input: strings.NewReader("10000000000")},
			wantErr:    ascii.ErrOverflow,
			wantRemain: []byte("10000000000"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ascii.UInt32()
			s, err := p.Parse(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			if err != io.EOF {
				remain, err := io.ReadAll(tt.args.input)
				require.NoError(t, err)
				assert.Equal(t, tt.wantRemain, remain)
			}
		})
	}
}

func TestUInt64(t *testing.T) {
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
			args:    args{input: strings.NewReader("")},
			wantErr: io.EOF,
		},
		{
			name:       "non digit => no match",
			args:       args{input: strings.NewReader("a")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'a'},
		},
		{
			name:       "negative => no match",
			args:       args{input: strings.NewReader("-a")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'-', 'a'},
		},
		{
			name:       "positive => no match",
			args:       args{input: strings.NewReader("+a")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'+', 'a'},
		},
		{
			name:       "digit => match",
			args:       args{input: strings.NewReader("9")},
			wantMatch:  9,
			wantRemain: []byte{},
		},
		{
			name:       "positive digit => match",
			args:       args{input: strings.NewReader("+9")},
			wantMatch:  9,
			wantRemain: []byte{},
		},
		{
			name:       "min number => match",
			args:       args{input: strings.NewReader("0")},
			wantMatch:  0,
			wantRemain: []byte{},
		},
		{
			name:       "max number => match",
			args:       args{input: strings.NewReader("18446744073709551615")},
			wantMatch:  math.MaxUint64,
			wantRemain: []byte{},
		},
		{
			name:       "one over max number => match",
			args:       args{input: strings.NewReader("18446744073709551616")},
			wantErr:    ascii.ErrOverflow,
			wantRemain: []byte("18446744073709551616"),
		},
		{
			name:       "large number => overflow",
			args:       args{input: strings.NewReader("1000000000000000000000")},
			wantErr:    ascii.ErrOverflow,
			wantRemain: []byte("1000000000000000000000"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ascii.UInt64()
			s, err := p.Parse(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			if err != io.EOF {
				remain, err := io.ReadAll(tt.args.input)
				require.NoError(t, err)
				assert.Equal(t, tt.wantRemain, remain)
			}
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
			args:    args{input: strings.NewReader("")},
			wantErr: io.EOF,
		},
		{
			name:       "non digit => no match",
			args:       args{input: strings.NewReader("a")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'a'},
		},
		{
			name:       "negative => no match",
			args:       args{input: strings.NewReader("-a")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'-', 'a'},
		},
		{
			name:       "positive => no match",
			args:       args{input: strings.NewReader("+a")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'+', 'a'},
		},
		{
			name:       "digit => match",
			args:       args{input: strings.NewReader("9")},
			wantMatch:  9,
			wantRemain: []byte{},
		},
		{
			name:       "negative digit => match",
			args:       args{input: strings.NewReader("-9")},
			wantMatch:  -9,
			wantRemain: []byte{},
		},
		{
			name:       "positive digit => match",
			args:       args{input: strings.NewReader("+9")},
			wantMatch:  9,
			wantRemain: []byte{},
		},
		{
			name:       "one under min number => match",
			args:       args{input: strings.NewReader("-129")},
			wantErr:    ascii.ErrOverflow,
			wantRemain: []byte("-129"),
		},
		{
			name:       "min number => match",
			args:       args{input: strings.NewReader("-128")},
			wantMatch:  math.MinInt8,
			wantRemain: []byte{},
		},
		{
			name:       "max number => match",
			args:       args{input: strings.NewReader("127")},
			wantMatch:  math.MaxInt8,
			wantRemain: []byte{},
		},
		{
			name:       "one over max number => match",
			args:       args{input: strings.NewReader("128")},
			wantErr:    ascii.ErrOverflow,
			wantRemain: []byte("128"),
		},
		{
			name:       "large number => overflow",
			args:       args{input: strings.NewReader("1000")},
			wantErr:    ascii.ErrOverflow,
			wantRemain: []byte("1000"),
		},
		{
			name:       "large negative number => overflow",
			args:       args{input: strings.NewReader("-1000")},
			wantErr:    ascii.ErrOverflow,
			wantRemain: []byte("-1000"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ascii.Int8()
			s, err := p.Parse(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			if err != io.EOF {
				remain, err := io.ReadAll(tt.args.input)
				require.NoError(t, err)
				assert.Equal(t, tt.wantRemain, remain)
			}
		})
	}
}

func TestInt16(t *testing.T) {
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
			name:    "empty input => EOF",
			args:    args{input: strings.NewReader("")},
			wantErr: io.EOF,
		},
		{
			name:       "non digit => no match",
			args:       args{input: strings.NewReader("a")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'a'},
		},
		{
			name:       "negative => no match",
			args:       args{input: strings.NewReader("-a")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'-', 'a'},
		},
		{
			name:       "positive => no match",
			args:       args{input: strings.NewReader("+a")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'+', 'a'},
		},
		{
			name:       "digit => match",
			args:       args{input: strings.NewReader("9")},
			wantMatch:  9,
			wantRemain: []byte{},
		},
		{
			name:       "negative digit => match",
			args:       args{input: strings.NewReader("-9")},
			wantMatch:  -9,
			wantRemain: []byte{},
		},
		{
			name:       "positive digit => match",
			args:       args{input: strings.NewReader("+9")},
			wantMatch:  9,
			wantRemain: []byte{},
		},
		{
			name:       "one under min number => match",
			args:       args{input: strings.NewReader("-32769")},
			wantErr:    ascii.ErrOverflow,
			wantRemain: []byte("-32769"),
		},
		{
			name:       "min number => match",
			args:       args{input: strings.NewReader("-32768")},
			wantMatch:  math.MinInt16,
			wantRemain: []byte{},
		},
		{
			name:       "max number => match",
			args:       args{input: strings.NewReader("32767")},
			wantMatch:  math.MaxInt16,
			wantRemain: []byte{},
		},
		{
			name:       "one over max number => match",
			args:       args{input: strings.NewReader("32768")},
			wantErr:    ascii.ErrOverflow,
			wantRemain: []byte("32768"),
		},
		{
			name:       "large number => overflow",
			args:       args{input: strings.NewReader("100000")},
			wantErr:    ascii.ErrOverflow,
			wantRemain: []byte("100000"),
		},
		{
			name:       "large negative number => overflow",
			args:       args{input: strings.NewReader("-100000")},
			wantErr:    ascii.ErrOverflow,
			wantRemain: []byte("-100000"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ascii.Int16()
			s, err := p.Parse(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			if err != io.EOF {
				remain, err := io.ReadAll(tt.args.input)
				require.NoError(t, err)
				assert.Equal(t, tt.wantRemain, remain)
			}
		})
	}
}

func TestInt32(t *testing.T) {
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
			name:    "empty input => EOF",
			args:    args{input: strings.NewReader("")},
			wantErr: io.EOF,
		},
		{
			name:       "non digit => no match",
			args:       args{input: strings.NewReader("a")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'a'},
		},
		{
			name:       "negative => no match",
			args:       args{input: strings.NewReader("-a")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'-', 'a'},
		},
		{
			name:       "positive => no match",
			args:       args{input: strings.NewReader("+a")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'+', 'a'},
		},
		{
			name:       "digit => match",
			args:       args{input: strings.NewReader("9")},
			wantMatch:  9,
			wantRemain: []byte{},
		},
		{
			name:       "negative digit => match",
			args:       args{input: strings.NewReader("-9")},
			wantMatch:  -9,
			wantRemain: []byte{},
		},
		{
			name:       "positive digit => match",
			args:       args{input: strings.NewReader("+9")},
			wantMatch:  9,
			wantRemain: []byte{},
		},
		{
			name:       "one under min number => match",
			args:       args{input: strings.NewReader("-2147483649")},
			wantErr:    ascii.ErrOverflow,
			wantRemain: []byte("-2147483649"),
		},
		{
			name:       "min number => match",
			args:       args{input: strings.NewReader("-2147483648")},
			wantMatch:  math.MinInt32,
			wantRemain: []byte{},
		},
		{
			name:       "max number => match",
			args:       args{input: strings.NewReader("2147483647")},
			wantMatch:  math.MaxInt32,
			wantRemain: []byte{},
		},
		{
			name:       "one over max number => match",
			args:       args{input: strings.NewReader("2147483648")},
			wantErr:    ascii.ErrOverflow,
			wantRemain: []byte("2147483648"),
		},
		{
			name:       "large number => overflow",
			args:       args{input: strings.NewReader("10000000000")},
			wantErr:    ascii.ErrOverflow,
			wantRemain: []byte("10000000000"),
		},
		{
			name:       "large negative number => overflow",
			args:       args{input: strings.NewReader("-10000000000")},
			wantErr:    ascii.ErrOverflow,
			wantRemain: []byte("-10000000000"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ascii.Int32()
			s, err := p.Parse(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			if err != io.EOF {
				remain, err := io.ReadAll(tt.args.input)
				require.NoError(t, err)
				assert.Equal(t, tt.wantRemain, remain)
			}
		})
	}
}

func TestInt64(t *testing.T) {
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
			name:    "empty input => EOF",
			args:    args{input: strings.NewReader("")},
			wantErr: io.EOF,
		},
		{
			name:       "non digit => no match",
			args:       args{input: strings.NewReader("a")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'a'},
		},
		{
			name:       "negative => no match",
			args:       args{input: strings.NewReader("-a")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'-', 'a'},
		},
		{
			name:       "positive => no match",
			args:       args{input: strings.NewReader("+a")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'+', 'a'},
		},
		{
			name:       "digit => match",
			args:       args{input: strings.NewReader("9")},
			wantMatch:  9,
			wantRemain: []byte{},
		},
		{
			name:       "negative digit => match",
			args:       args{input: strings.NewReader("-9")},
			wantMatch:  -9,
			wantRemain: []byte{},
		},
		{
			name:       "positive digit => match",
			args:       args{input: strings.NewReader("+9")},
			wantMatch:  9,
			wantRemain: []byte{},
		},
		{
			name:       "min number => match",
			args:       args{input: strings.NewReader("0")},
			wantMatch:  0,
			wantRemain: []byte{},
		},
		{
			name:       "one under min number => match",
			args:       args{input: strings.NewReader("-9223372036854775809")},
			wantErr:    ascii.ErrOverflow,
			wantRemain: []byte("-9223372036854775809"),
		},
		{
			name:       "min number => match",
			args:       args{input: strings.NewReader("-9223372036854775808")},
			wantMatch:  math.MinInt64,
			wantRemain: []byte{},
		},
		{
			name:       "max number => match",
			args:       args{input: strings.NewReader("9223372036854775807")},
			wantMatch:  math.MaxInt64,
			wantRemain: []byte{},
		},
		{
			name:       "one over max number => match",
			args:       args{input: strings.NewReader("9223372036854775808")},
			wantErr:    ascii.ErrOverflow,
			wantRemain: []byte("9223372036854775808"),
		},
		{
			name:       "large number => overflow",
			args:       args{input: strings.NewReader("10000000000000000000")},
			wantErr:    ascii.ErrOverflow,
			wantRemain: []byte("10000000000000000000"),
		},
		{
			name:       "large negative number => overflow",
			args:       args{input: strings.NewReader("-10000000000000000000")},
			wantErr:    ascii.ErrOverflow,
			wantRemain: []byte("-10000000000000000000"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ascii.Int64()
			s, err := p.Parse(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			if err != io.EOF {
				remain, err := io.ReadAll(tt.args.input)
				require.NoError(t, err)
				assert.Equal(t, tt.wantRemain, remain)
			}
		})
	}
}
