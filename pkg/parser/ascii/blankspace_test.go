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

func TestBlankSpace(t *testing.T) {
	type args struct {
		input parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  byte
		wantRemain []byte
		wantErr    error
	}{
		{
			name:       "empty input => EOF",
			args:       args{input: strings.NewReader("")},
			wantErr:    io.EOF,
			wantRemain: []byte{},
		},
		{
			name:       "no match => not matched",
			args:       args{input: strings.NewReader("a")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'a'},
		},
		{
			name:       "match space => match",
			args:       args{input: strings.NewReader(" ")},
			wantMatch:  ' ',
			wantRemain: []byte{},
		},
		{
			name:       "match tab => match",
			args:       args{input: strings.NewReader("\t")},
			wantMatch:  '\t',
			wantRemain: []byte{},
		},
		{
			name:       "match CR => match",
			args:       args{input: strings.NewReader("\r")},
			wantMatch:  '\r',
			wantRemain: []byte{},
		},
		{
			name:       "match LF => match",
			args:       args{input: strings.NewReader("\n")},
			wantMatch:  '\n',
			wantRemain: []byte{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ascii.BlankSpace()
			s, err := p(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)

		})
	}
}

func TestBlankSpace0(t *testing.T) {
	type args struct {
		input parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  []byte
		wantRemain []byte
	}{
		{
			name:       "empty input => EOF",
			args:       args{input: strings.NewReader("")},
			wantMatch:  []byte{},
			wantRemain: []byte{},
		},
		{
			name:       "no match => empty match",
			args:       args{input: strings.NewReader("a")},
			wantMatch:  []byte{},
			wantRemain: []byte{'a'},
		},
		{
			name:       "match => match",
			args:       args{input: strings.NewReader(" \t\r\n")},
			wantMatch:  []byte(" \t\r\n"),
			wantRemain: []byte{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ascii.BlankSpace0()
			s, err := p(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.NoError(t, err)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)

		})
	}
}

func TestBlankSpace1(t *testing.T) {
	type args struct {
		input parser.Reader
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
			args:       args{input: strings.NewReader("")},
			wantErr:    io.EOF,
			wantRemain: []byte{},
		},
		{
			name:       "no match => not matched",
			args:       args{input: strings.NewReader("a")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'a'},
		},
		{
			name:       "match => match",
			args:       args{input: strings.NewReader(" \t\r\n")},
			wantMatch:  []byte(" \t\r\n"),
			wantRemain: []byte{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ascii.BlankSpace1()
			s, err := p(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)

		})
	}
}

func TestSkipBlankSpace(t *testing.T) {
	type args struct {
		input parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantRemain []byte
		wantErr    error
	}{
		{
			name:       "empty input => EOF",
			args:       args{input: strings.NewReader("")},
			wantErr:    io.EOF,
			wantRemain: []byte{},
		},
		{
			name:       "no match => not matched",
			args:       args{input: strings.NewReader("a")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'a'},
		},
		{
			name:       "match space => match",
			args:       args{input: strings.NewReader(" ")},
			wantRemain: []byte{},
		},
		{
			name:       "match tab => match",
			args:       args{input: strings.NewReader("\t")},
			wantRemain: []byte{},
		},
		{
			name:       "match CR => match",
			args:       args{input: strings.NewReader("\r")},
			wantRemain: []byte{},
		},
		{
			name:       "match LF => match",
			args:       args{input: strings.NewReader("\n")},
			wantRemain: []byte{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ascii.SkipBlankSpace()
			s, err := p(tt.args.input)

			assert.Nil(t, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)

		})
	}
}

func TestSkipBlankSpace0(t *testing.T) {
	type args struct {
		input parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantRemain []byte
	}{
		{
			name:       "empty input => EOF",
			args:       args{input: strings.NewReader("")},
			wantRemain: []byte{},
		},
		{
			name: "no match => empty match",
			args: args{input: strings.NewReader("a")},

			wantRemain: []byte{'a'},
		},
		{
			name:       "match => match",
			args:       args{input: strings.NewReader(" \t\r\n")},
			wantRemain: []byte{},
		},
		{
			name:       "match with remainder  => match",
			args:       args{input: strings.NewReader(" \t\r\na")},
			wantRemain: []byte{'a'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ascii.SkipBlankSpace0()
			s, err := p(tt.args.input)

			assert.Nil(t, s)
			assert.NoError(t, err)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)

		})
	}
}

func TestSkipBlankSpace1(t *testing.T) {
	type args struct {
		input parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantRemain []byte
		wantErr    error
	}{
		{
			name:       "empty input => EOF",
			args:       args{input: strings.NewReader("")},
			wantErr:    io.EOF,
			wantRemain: []byte{},
		},
		{
			name:       "no match => not matched",
			args:       args{input: strings.NewReader("a")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'a'},
		},
		{
			name:       "match => match",
			args:       args{input: strings.NewReader(" \t\r\n")},
			wantRemain: []byte{},
		},
		{
			name:       "match with remainder => match",
			args:       args{input: strings.NewReader(" \t\r\na")},
			wantRemain: []byte{'a'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ascii.SkipBlankSpace1()
			s, err := p(tt.args.input)

			assert.Nil(t, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)

		})
	}
}
