package ascii_test

import (
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/roblovelock/gobble/pkg/parser/ascii"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"strings"
	"testing"
)

func TestFloat64(t *testing.T) {
	type args struct {
		input parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  float64
		wantRemain []byte
		wantErr    error
	}{
		{
			name:       "empty input => EOF",
			args:       args{input: strings.NewReader("")},
			wantRemain: []byte{},
			wantErr:    io.EOF,
		},
		{
			name:       "non digit => no match",
			args:       args{input: strings.NewReader("a")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'a'},
		},
		{
			name:       "negative non digit => no match",
			args:       args{input: strings.NewReader("-a")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'-', 'a'},
		},
		{
			name:       "positive non digit => no match",
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
			name:       "negative digit => match",
			args:       args{input: strings.NewReader("-9")},
			wantMatch:  -9,
			wantRemain: []byte{},
		},
		{
			name:       "decimal => match",
			args:       args{input: strings.NewReader("12.345")},
			wantMatch:  12.345,
			wantRemain: []byte{},
		},
		{
			name:       "positive decimal => match",
			args:       args{input: strings.NewReader("+12.345")},
			wantMatch:  12.345,
			wantRemain: []byte{},
		},
		{
			name:       "negative decimal => match",
			args:       args{input: strings.NewReader("-12.345")},
			wantMatch:  -12.345,
			wantRemain: []byte{},
		},
		{
			name:       "scientific decimal number => match",
			args:       args{input: strings.NewReader("11e-1")},
			wantMatch:  11e-1,
			wantRemain: []byte{},
		},
		{
			name:       "uppercase scientific decimal number => match",
			args:       args{input: strings.NewReader("123E-02")},
			wantMatch:  123e-02,
			wantRemain: []byte{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ascii.Float64()
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
