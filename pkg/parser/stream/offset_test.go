package stream_test

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/roblovelock/gobble/pkg/parser/stream"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"strings"
	"testing"
)

func TestOffset(t *testing.T) {
	type args struct {
		readBytes int
		input     parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantRemain string
		want       int64
	}{
		{
			name: "empty input => 0",
			args: args{input: strings.NewReader("")},
		},
		{
			name:       "at start => 0",
			args:       args{input: strings.NewReader("a")},
			wantRemain: "a",
		},
		{
			name:       "in middle => 1",
			args:       args{readBytes: 1, input: strings.NewReader("ab")},
			wantRemain: "b",
			want:       1,
		},
		{
			name:       "at end => 2",
			args:       args{readBytes: 2, input: strings.NewReader("ab")},
			wantRemain: "",
			want:       2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := stream.Offset()
			for i := 0; i < tt.args.readBytes; i++ {
				_, _ = tt.args.input.ReadByte()
			}
			s, err := p(tt.args.input)

			assert.Equal(t, tt.want, s)
			assert.NoError(t, err)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, string(remain))
		})
	}
}
