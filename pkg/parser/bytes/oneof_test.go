package bytes_test

import (
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/roblovelock/gobble/pkg/parser/bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"strings"
	"testing"
)

func TestOneOf(t *testing.T) {
	type args struct {
		bytes []byte
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
			name:    "empty input => EOF",
			args:    args{bytes: []byte{'a'}, input: strings.NewReader("")},
			wantErr: io.EOF,
		},
		{
			name:       "empty bytes => no match",
			args:       args{bytes: []byte{}, input: strings.NewReader("a")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'a'},
		},
		{
			name:       "byte mismatch => no match",
			args:       args{bytes: []byte{'a'}, input: strings.NewReader("b")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'b'},
		},
		{
			name:       "bytes mismatch => no match",
			args:       args{bytes: []byte{'a', 'b', 'c'}, input: strings.NewReader("d")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'d'},
		},
		{
			name:       "byte match => match",
			args:       args{bytes: []byte{'a'}, input: strings.NewReader("a")},
			wantMatch:  'a',
			wantRemain: []byte{},
		},
		{
			name:       "bytes match => match",
			args:       args{bytes: []byte{'a', 'b', 'c'}, input: strings.NewReader("bb")},
			wantMatch:  'b',
			wantRemain: []byte{'b'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := bytes.OneOf(tt.args.bytes...)
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

func TestOneOf1(t *testing.T) {
	type args struct {
		bytes []byte
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
			name:    "empty input => EOF",
			args:    args{bytes: []byte{'a'}, input: strings.NewReader("")},
			wantErr: io.EOF,
		},
		{
			name:       "empty bytes => no match",
			args:       args{bytes: []byte{}, input: strings.NewReader("a")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'a'},
		},
		{
			name:       "byte mismatch => no match",
			args:       args{bytes: []byte{'a'}, input: strings.NewReader("b")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'b'},
		},
		{
			name:       "bytes mismatch => no match",
			args:       args{bytes: []byte{'a', 'b', 'c'}, input: strings.NewReader("d")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'d'},
		},
		{
			name:       "byte match => match",
			args:       args{bytes: []byte{'a'}, input: strings.NewReader("a")},
			wantMatch:  []byte{'a'},
			wantRemain: []byte{},
		},
		{
			name:       "bytes match => match",
			args:       args{bytes: []byte{'a', 'b', 'c'}, input: strings.NewReader("b")},
			wantMatch:  []byte{'b'},
			wantRemain: []byte{},
		},
		{
			name:       "byte match many => match",
			args:       args{bytes: []byte{'a'}, input: strings.NewReader("aaaa")},
			wantMatch:  []byte{'a', 'a', 'a', 'a'},
			wantRemain: []byte{},
		},
		{
			name:       "bytes match many => match",
			args:       args{bytes: []byte{'a', 'b', 'c'}, input: strings.NewReader("cbad")},
			wantMatch:  []byte{'c', 'b', 'a'},
			wantRemain: []byte{'d'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := bytes.OneOf1(tt.args.bytes...)
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

func TestOneOf0(t *testing.T) {
	type args struct {
		bytes []byte
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
			args:       args{bytes: []byte{'a'}, input: strings.NewReader("")},
			wantMatch:  []byte{},
			wantRemain: []byte{},
		},
		{
			name:       "empty bytes => no match",
			args:       args{bytes: []byte{}, input: strings.NewReader("a")},
			wantMatch:  []byte{},
			wantRemain: []byte{'a'},
		},
		{
			name:       "byte mismatch => no match",
			args:       args{bytes: []byte{'a'}, input: strings.NewReader("b")},
			wantMatch:  []byte{},
			wantRemain: []byte{'b'},
		},
		{
			name:       "bytes mismatch => no match",
			args:       args{bytes: []byte{'a', 'b', 'c'}, input: strings.NewReader("d")},
			wantMatch:  []byte{},
			wantRemain: []byte{'d'},
		},
		{
			name:       "byte match => match",
			args:       args{bytes: []byte{'a'}, input: strings.NewReader("a")},
			wantMatch:  []byte{'a'},
			wantRemain: []byte{},
		},
		{
			name:       "bytes match => match",
			args:       args{bytes: []byte{'a', 'b', 'c'}, input: strings.NewReader("b")},
			wantMatch:  []byte{'b'},
			wantRemain: []byte{},
		},
		{
			name:       "byte match many => match",
			args:       args{bytes: []byte{'a'}, input: strings.NewReader("aaaa")},
			wantMatch:  []byte{'a', 'a', 'a', 'a'},
			wantRemain: []byte{},
		},
		{
			name:       "bytes match many => match",
			args:       args{bytes: []byte{'a', 'b', 'c'}, input: strings.NewReader("cbad")},
			wantMatch:  []byte{'c', 'b', 'a'},
			wantRemain: []byte{'d'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := bytes.OneOf0(tt.args.bytes...)
			s, err := p.Parse(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.NoError(t, err)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)
		})
	}
}
