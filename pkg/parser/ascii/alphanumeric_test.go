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

func TestAlphanumeric(t *testing.T) {
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
			name:    "empty input => EOF",
			args:    args{input: strings.NewReader("")},
			wantErr: io.EOF,
		},
		{
			name:       "non alphanumeric => no match",
			args:       args{input: strings.NewReader("+")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'+'},
		},
		{
			name:       "alphanumeric digit => no match",
			args:       args{input: strings.NewReader("8")},
			wantMatch:  '8',
			wantRemain: []byte{},
		},
		{
			name:       "alphanumeric letter => match",
			args:       args{input: strings.NewReader("a")},
			wantMatch:  'a',
			wantRemain: []byte{},
		},
		{
			name:       "alphanumeric => match with remainder",
			args:       args{input: strings.NewReader("Ab")},
			wantMatch:  'A',
			wantRemain: []byte{'b'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ascii.Alphanumeric()
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

func TestAlphanumeric1(t *testing.T) {
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
			name:    "empty input => EOF",
			args:    args{input: strings.NewReader("")},
			wantErr: io.EOF,
		},
		{
			name:       "non alphanumeric => no match",
			args:       args{input: strings.NewReader("+")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'+'},
		},
		{
			name:       "alphanumeric digit => match",
			args:       args{input: strings.NewReader("8")},
			wantMatch:  []byte{'8'},
			wantRemain: []byte{},
		},
		{
			name:       "alphanumeric => match",
			args:       args{input: strings.NewReader("a")},
			wantMatch:  []byte{'a'},
			wantRemain: []byte{},
		},
		{
			name:       "alphanumeric => match with remainder",
			args:       args{input: strings.NewReader("Ab8+")},
			wantMatch:  []byte{'A', 'b', '8'},
			wantRemain: []byte{'+'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ascii.Alphanumeric1()
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

func TestAlphanumeric0(t *testing.T) {
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
			name:       "empty input => empty match",
			args:       args{input: strings.NewReader("")},
			wantMatch:  []byte{},
			wantRemain: []byte{},
		},
		{
			name:       "non alphanumeric => empty match",
			args:       args{input: strings.NewReader("+")},
			wantMatch:  []byte{},
			wantRemain: []byte{'+'},
		},
		{
			name:       "alphanumeric digit => match",
			args:       args{input: strings.NewReader("8")},
			wantMatch:  []byte{'8'},
			wantRemain: []byte{},
		},
		{
			name:       "alphanumeric => match",
			args:       args{input: strings.NewReader("a")},
			wantMatch:  []byte{'a'},
			wantRemain: []byte{},
		},
		{
			name:       "alphanumeric => match with remainder",
			args:       args{input: strings.NewReader("Ab8+")},
			wantMatch:  []byte{'A', 'b', '8'},
			wantRemain: []byte{'+'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ascii.Alphanumeric0()
			s, err := p.Parse(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.NoError(t, err)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)
		})
	}
}
