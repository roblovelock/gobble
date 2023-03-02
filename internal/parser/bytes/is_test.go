package bytes

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gobble/internal/parser"
	"io"
	"strings"
	"testing"
)

func TestIsA(t *testing.T) {
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
			wantErr:    parser.ErrNotMatched,
			wantRemain: []byte{'a'},
		},
		{
			name:       "byte mismatch => no match",
			args:       args{bytes: []byte{'a'}, input: strings.NewReader("b")},
			wantErr:    parser.ErrNotMatched,
			wantRemain: []byte{'b'},
		},
		{
			name:       "bytes mismatch => no match",
			args:       args{bytes: []byte{'a', 'b', 'c'}, input: strings.NewReader("d")},
			wantErr:    parser.ErrNotMatched,
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
			p := IsA(tt.args.bytes...)
			s, err := p(tt.args.input)

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
