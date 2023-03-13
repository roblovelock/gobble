package combinator

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"strings"
	"testing"
)

func TestSuccess(t *testing.T) {
	type args struct {
		value string
		input parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  string
		wantRemain string
	}{
		{
			name: "empty input => success",
			args: args{
				value: "success",
				input: strings.NewReader(""),
			},
			wantMatch:  "success",
			wantRemain: "",
		},
		{
			name: "input => success",
			args: args{
				value: "success",
				input: strings.NewReader("ab"),
			},
			wantMatch:  "success",
			wantRemain: "ab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Success[parser.Reader, string](tt.args.value)
			s, err := p.Parse(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.NoError(t, err)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, string(remain))
		})
	}
}
