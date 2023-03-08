package modifier_test

import (
	"github.com/roblovelock/gobble/pkg/combinator/modifier"
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/roblovelock/gobble/pkg/parser/runes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"strconv"
	"strings"
	"testing"
	"unicode"
)

func TestMap(t *testing.T) {
	_, wantErr := strconv.Atoi("a")
	type args struct {
		parser  parser.Parser[parser.Reader, string]
		mapFunc func(string) (int, error)
		input   parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  int
		wantRemain []byte
		wantErr    error
	}{
		{
			name: "empty input => EOF",
			args: args{
				parser:  runes.TakeWhile1(unicode.IsDigit),
				mapFunc: strconv.Atoi,
				input:   strings.NewReader(""),
			},
			wantErr:    io.EOF,
			wantRemain: []byte{},
		},
		{
			name: "non digit => no match",
			args: args{
				parser:  runes.TakeWhile1(unicode.IsDigit),
				mapFunc: strconv.Atoi,
				input:   strings.NewReader("a"),
			},
			wantErr:    parser.ErrNotMatched,
			wantRemain: []byte{'a'},
		},
		{
			name: "map error => no match",
			args: args{
				parser:  runes.Take(1),
				mapFunc: strconv.Atoi,
				input:   strings.NewReader("a"),
			},
			wantErr:    wantErr,
			wantRemain: []byte{'a'},
		},
		{
			name: "digit => match",
			args: args{
				parser:  runes.TakeWhile1(unicode.IsDigit),
				mapFunc: strconv.Atoi,
				input:   strings.NewReader("9"),
			},
			wantMatch:  9,
			wantRemain: []byte{},
		},
		{
			name: "digits => match",
			args: args{
				parser:  runes.TakeWhile1(unicode.IsDigit),
				mapFunc: strconv.Atoi,
				input:   strings.NewReader("12"),
			},
			wantMatch:  12,
			wantRemain: []byte{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := modifier.Map(tt.args.parser, tt.args.mapFunc)
			s, err := p(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			if tt.wantErr != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
			} else {
				assert.NoError(t, err)
			}

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)

		})
	}
}
