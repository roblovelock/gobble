package bytes_test

import (
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/roblovelock/gobble/pkg/parser/ascii"
	"github.com/roblovelock/gobble/pkg/parser/bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"strings"
	"testing"
)

func TestTakeWhileMinMax(t *testing.T) {
	type args struct {
		input     parser.Reader
		predicate parser.Predicate[byte]
		max       int
		min       int
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  []byte
		wantRemain []byte
		wantErr    error
	}{
		{
			name:       "empty input min 1 => EOF",
			args:       args{min: 1, max: 10, predicate: ascii.IsDigit, input: strings.NewReader("")},
			wantRemain: []byte{},
			wantErr:    io.EOF,
		},
		{
			name:       "empty input min 0 => empty",
			args:       args{min: 0, max: 10, predicate: ascii.IsDigit, input: strings.NewReader("")},
			wantMatch:  []byte{},
			wantRemain: []byte{},
		},
		{
			name:       "take min 1 no match => empty",
			args:       args{min: 1, max: 10, predicate: ascii.IsDigit, input: strings.NewReader("b")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'b'},
		},
		{
			name:       "take min 0 no match => empty",
			args:       args{min: 0, max: 10, predicate: ascii.IsDigit, input: strings.NewReader("b")},
			wantMatch:  []byte{},
			wantRemain: []byte{'b'},
		},
		{
			name:       "take max 1 match => match 1",
			args:       args{min: 0, max: 1, predicate: ascii.IsDigit, input: strings.NewReader("12b")},
			wantMatch:  []byte{'1'},
			wantRemain: []byte{'2', 'b'},
		},
		{
			name:       "take max 10 match => match all",
			args:       args{min: 0, max: 10, predicate: ascii.IsDigit, input: strings.NewReader("12b")},
			wantMatch:  []byte{'1', '2'},
			wantRemain: []byte{'b'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := bytes.TakeWhileMinMax(tt.args.min, tt.args.max, tt.args.predicate)
			s, err := p.Parse(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)
		})
	}
}

func TestTakeWhile1(t *testing.T) {
	type args struct {
		input     parser.Reader
		predicate parser.Predicate[byte]
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  []byte
		wantRemain []byte
		wantErr    error
	}{
		{
			name:       "empty input => EOF",
			args:       args{predicate: ascii.IsDigit, input: strings.NewReader("")},
			wantRemain: []byte{},
			wantErr:    io.EOF,
		},
		{
			name:       "take no match => empty",
			args:       args{predicate: ascii.IsDigit, input: strings.NewReader("b")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'b'},
		},
		{
			name:       "take match => match all",
			args:       args{predicate: ascii.IsDigit, input: strings.NewReader("12b")},
			wantMatch:  []byte{'1', '2'},
			wantRemain: []byte{'b'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := bytes.TakeWhile1(tt.args.predicate)
			s, err := p.Parse(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)
		})
	}
}

func TestTakeWhile(t *testing.T) {
	type args struct {
		input     parser.Reader
		predicate parser.Predicate[byte]
		max       int
		min       int
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  []byte
		wantRemain []byte
	}{
		{
			name:       "empty input => empty",
			args:       args{min: 1, max: 10, predicate: ascii.IsDigit, input: strings.NewReader("")},
			wantRemain: []byte{},
			wantMatch:  []byte{},
		},
		{
			name:       "take no match => empty",
			args:       args{min: 0, max: 10, predicate: ascii.IsDigit, input: strings.NewReader("b")},
			wantMatch:  []byte{},
			wantRemain: []byte{'b'},
		},
		{
			name:       "take match => match all",
			args:       args{min: 0, max: 10, predicate: ascii.IsDigit, input: strings.NewReader("12b")},
			wantMatch:  []byte{'1', '2'},
			wantRemain: []byte{'b'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := bytes.TakeWhile(tt.args.predicate)
			s, err := p.Parse(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.NoError(t, err)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)
		})
	}
}
