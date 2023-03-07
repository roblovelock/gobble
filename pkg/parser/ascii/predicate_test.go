package ascii_test

import (
	"github.com/roblovelock/gobble/pkg/parser/ascii"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestIsHexDigit(t *testing.T) {
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
			name: "digit characters match",
			args: args{start: '0', end: '9'},
			want: true,
		},
		{
			name: "lowercase hex letter characters match",
			args: args{start: 'a', end: 'f'},
			want: true,
		},
		{
			name: "uppercase hex letter characters match",
			args: args{start: 'A', end: 'F'},
			want: true,
		},
		{
			name: "non hex characters < 0 don't match",
			args: args{start: 0, end: '0' - 1},
			want: false,
		},
		{
			name: "non hex characters (9 < character < A) don't match",
			args: args{start: '9' + 1, end: 'A' - 1},
			want: false,
		},
		{
			name: "non alphanumeric characters (F < character < a) don't match",
			args: args{start: 'F' + 1, end: 'a' - 1},
			want: false,
		},
		{
			name: "non alphanumeric characters > f don't match",
			args: args{start: 'f' + 1, end: math.MaxUint8},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := tt.args.start; i < tt.args.end; i++ {
				assert.Equalf(t, tt.want, ascii.IsHexDigit(i), "IsHexDigit(%v)", i)
			}
		})
	}
}

func TestIsDigit(t *testing.T) {
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
			name: "digit characters match",
			args: args{start: '0', end: '9'},
			want: true,
		},
		{
			name: "non digit characters < 0 don't match",
			args: args{start: 0, end: '0' - 1},
			want: false,
		},
		{
			name: "non digit characters > 9 don't match",
			args: args{start: '9' + 1, end: math.MaxUint8},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := tt.args.start; i < tt.args.end; i++ {
				assert.Equalf(t, tt.want, ascii.IsDigit(i), "IsDigit(%v)", i)
			}
		})
	}
}

func TestIsLetter(t *testing.T) {
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
			name: "lowercase letter characters match",
			args: args{start: 'a', end: 'z'},
			want: true,
		},
		{
			name: "uppercase letter characters match",
			args: args{start: 'A', end: 'Z'},
			want: true,
		},
		{
			name: "non letter characters < A don't match",
			args: args{start: 0, end: 'A' - 1},
			want: false,
		},
		{
			name: "non letter characters (Z < character < a) don't match",
			args: args{start: 'Z' + 1, end: 'a' - 1},
			want: false,
		},
		{
			name: "non letter characters > z don't match",
			args: args{start: 'z' + 1, end: math.MaxUint8},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := tt.args.start; i < tt.args.end; i++ {
				assert.Equalf(t, tt.want, ascii.IsLetter(i), "IsLetter(%v)", i)
			}
		})
	}
}

func TestIsAlphanumeric(t *testing.T) {
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
			name: "digit characters match",
			args: args{start: '0', end: '9'},
			want: true,
		},
		{
			name: "lowercase letter characters match",
			args: args{start: 'a', end: 'z'},
			want: true,
		},
		{
			name: "uppercase letter characters match",
			args: args{start: 'A', end: 'Z'},
			want: true,
		},
		{
			name: "non alphanumeric characters < 0 don't match",
			args: args{start: 0, end: '0' - 1},
			want: false,
		},
		{
			name: "non alphanumeric characters (9 < character < A) don't match",
			args: args{start: '9' + 1, end: 'A' - 1},
			want: false,
		},
		{
			name: "non alphanumeric characters (Z < character < a) don't match",
			args: args{start: 'Z' + 1, end: 'a' - 1},
			want: false,
		},
		{
			name: "non alphanumeric characters > z don't match",
			args: args{start: 'z' + 1, end: math.MaxUint8},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := tt.args.start; i < tt.args.end; i++ {
				assert.Equalf(t, tt.want, ascii.IsAlphanumeric(i), "IsAlphanumeric(%v)", i)
			}
		})
	}
}

func TestIsASCII(t *testing.T) {
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
			name: "ascii characters match",
			args: args{start: 0x00, end: 0x7F},
			want: true,
		},
		{
			name: "non ascii characters don't match",
			args: args{start: 0x7F + 1, end: math.MaxUint8},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := tt.args.start; i < tt.args.end; i++ {
				assert.Equalf(t, tt.want, ascii.IsASCII(i), "IsASCII(%v)", i)
			}
		})
	}
}

func TestIsOctDigit(t *testing.T) {
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
			name: "octal digit characters match",
			args: args{start: '0', end: '7'},
			want: true,
		},
		{
			name: "non octal digit characters < '0' don't match",
			args: args{start: 0, end: '0' - 1},
			want: false,
		},
		{
			name: "non octal digit characters > '7' characters don't match",
			args: args{start: '7' + 1, end: math.MaxUint8},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := tt.args.start; i < tt.args.end; i++ {
				assert.Equalf(t, tt.want, ascii.IsOctDigit(i), "IsOctDigit(%v)", i)
			}
		})
	}
}
