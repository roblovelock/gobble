package ascii_test

import (
	"github.com/roblovelock/gobble/pkg/parser/ascii"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestIsNewLine(t *testing.T) {
	type args struct {
		start byte
		end   byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "newline character match",
			args: args{start: '\n', end: '\n'},
			want: true,
		},
		{
			name: "non newline characters < \\n don't match",
			args: args{start: 0, end: '\n' - 1},
			want: false,
		},
		{
			name: "non newline characters > \\n don't match",
			args: args{start: '\n' + 1, end: math.MaxUint8},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := tt.args.start; i < tt.args.end; i++ {
				assert.Equalf(t, tt.want, ascii.IsNewLine(i), "IsNewLine(%v)", i)
			}
		})
	}
}

func TestIsSpace(t *testing.T) {
	type args struct {
		start byte
		end   byte
		skip  []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "space character match",
			args: args{start: ' ', end: ' '},
			want: true,
		},
		{
			name: "tab character match",
			args: args{start: '\t', end: '\t'},
			want: true,
		},
		{
			name: "non space or tab characters don't match",
			args: args{start: 0, end: math.MaxUint8, skip: []byte{' ', '\t'}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := tt.args.start; i < tt.args.end; i++ {
				if skipCharacter(i, tt.args.skip) {
					continue
				}
				assert.Equalf(t, tt.want, ascii.IsSpace(i), "IsSpace(%v)", i)
			}
		})
	}
}

func TestIsBlankSpace(t *testing.T) {
	type args struct {
		start byte
		end   byte
		skip  []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "space character match",
			args: args{start: ' ', end: ' '},
			want: true,
		},
		{
			name: "tab character match",
			args: args{start: '\t', end: '\t'},
			want: true,
		},
		{
			name: "\\r character match",
			args: args{start: '\r', end: '\r'},
			want: true,
		},
		{
			name: "\\n character match",
			args: args{start: '\n', end: '\n'},
			want: true,
		},
		{
			name: "non space, tab or newline characters don't match",
			args: args{start: 0, end: math.MaxUint8, skip: []byte{' ', '\t', '\r', '\n'}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := tt.args.start; i < tt.args.end; i++ {
				if skipCharacter(i, tt.args.skip) {
					continue
				}
				assert.Equalf(t, tt.want, ascii.IsBlankSpace(i), "IsBlankSpace(%v)", i)
			}
		})
	}
}

func TestIsWhitespace(t *testing.T) {
	type args struct {
		start byte
		end   byte
		skip  []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "space character match",
			args: args{start: ' ', end: ' '},
			want: true,
		},
		{
			name: "tab character match",
			args: args{start: '\t', end: '\t'},
			want: true,
		},
		{
			name: "\\r character match",
			args: args{start: '\r', end: '\r'},
			want: true,
		},
		{
			name: "\\n character match",
			args: args{start: '\n', end: '\n'},
			want: true,
		},
		{
			name: "\\v character match",
			args: args{start: '\v', end: '\v'},
			want: true,
		},
		{
			name: "\\f character match",
			args: args{start: '\f', end: '\f'},
			want: true,
		},
		{
			name: "non space, tab or newline characters don't match",
			args: args{start: 0, end: math.MaxUint8, skip: []byte{' ', '\t', '\r', '\n', '\v', '\f'}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := tt.args.start; i < tt.args.end; i++ {
				if skipCharacter(i, tt.args.skip) {
					continue
				}
				assert.Equalf(t, tt.want, ascii.IsWhitespace(i), "IsWhitespace(%v)", i)
			}
		})
	}
}

func skipCharacter(c byte, skip []byte) bool {
	for _, e := range skip {
		if e == c {
			return true
		}
	}
	return false
}
