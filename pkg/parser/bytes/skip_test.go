package bytes_test

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/roblovelock/gobble/pkg/parser/ascii"
	"github.com/roblovelock/gobble/pkg/parser/bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"strings"
	"testing"
)

func TestSkip(t *testing.T) {
	type args struct {
		input     parser.Reader
		predicate parser.Predicate[byte]
	}
	tests := []struct {
		name       string
		args       args
		wantRemain []byte
		wantErr    error
	}{
		{
			name:       "empty input => EOF",
			args:       args{predicate: ascii.IsWhitespace, input: strings.NewReader("")},
			wantErr:    io.EOF,
			wantRemain: []byte{},
		},
		{
			name:       "no match => not matched",
			args:       args{predicate: ascii.IsWhitespace, input: strings.NewReader("a")},
			wantErr:    parser.ErrNotMatched,
			wantRemain: []byte{'a'},
		},
		{
			name:       "match space => match",
			args:       args{predicate: ascii.IsWhitespace, input: strings.NewReader(" ")},
			wantRemain: []byte{},
		},
		{
			name:       "match tab => match",
			args:       args{predicate: ascii.IsWhitespace, input: strings.NewReader("\t")},
			wantRemain: []byte{},
		},
		{
			name:       "match CR => match",
			args:       args{predicate: ascii.IsWhitespace, input: strings.NewReader("\r")},
			wantRemain: []byte{},
		},
		{
			name:       "match LF => match",
			args:       args{predicate: ascii.IsWhitespace, input: strings.NewReader("\n")},
			wantRemain: []byte{},
		},
		{
			name:       "match vertical tab => match",
			args:       args{predicate: ascii.IsWhitespace, input: strings.NewReader("\v")},
			wantRemain: []byte{},
		},
		{
			name:       "match form feed => match",
			args:       args{predicate: ascii.IsWhitespace, input: strings.NewReader("\f")},
			wantRemain: []byte{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := bytes.Skip(tt.args.predicate)
			s, err := p(tt.args.input)

			assert.Nil(t, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)

		})
	}
}

func TestSkip0(t *testing.T) {
	type args struct {
		input     parser.Reader
		predicate parser.Predicate[byte]
	}
	tests := []struct {
		name       string
		args       args
		wantRemain []byte
	}{
		{
			name:       "empty input => EOF",
			args:       args{predicate: ascii.IsWhitespace, input: strings.NewReader("")},
			wantRemain: []byte{},
		},
		{
			name: "no match => empty match",
			args: args{predicate: ascii.IsWhitespace, input: strings.NewReader("a")},

			wantRemain: []byte{'a'},
		},
		{
			name:       "match => match",
			args:       args{predicate: ascii.IsWhitespace, input: strings.NewReader(" \t\r\n\v\f")},
			wantRemain: []byte{},
		},
		{
			name:       "match with remainder  => match",
			args:       args{predicate: ascii.IsWhitespace, input: strings.NewReader(" \t\r\n\v\fa")},
			wantRemain: []byte{'a'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := bytes.SkipWhile(tt.args.predicate)
			s, err := p(tt.args.input)

			assert.Nil(t, s)
			assert.NoError(t, err)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)

		})
	}
}

func TestSkip1(t *testing.T) {
	type args struct {
		input     parser.Reader
		predicate parser.Predicate[byte]
	}
	tests := []struct {
		name       string
		args       args
		wantRemain []byte
		wantErr    error
	}{
		{
			name:       "empty input => EOF",
			args:       args{predicate: ascii.IsWhitespace, input: strings.NewReader("")},
			wantErr:    io.EOF,
			wantRemain: []byte{},
		},
		{
			name:       "no match => not matched",
			args:       args{predicate: ascii.IsWhitespace, input: strings.NewReader("a")},
			wantErr:    parser.ErrNotMatched,
			wantRemain: []byte{'a'},
		},
		{
			name:       "match => match",
			args:       args{predicate: ascii.IsWhitespace, input: strings.NewReader(" \t\r\n\v\f")},
			wantRemain: []byte{},
		},
		{
			name:       "match with remainder => match",
			args:       args{predicate: ascii.IsWhitespace, input: strings.NewReader(" \t\r\n\v\fa")},
			wantRemain: []byte{'a'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := bytes.SkipWhile1(tt.args.predicate)
			s, err := p(tt.args.input)

			assert.Nil(t, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)

		})
	}
}
