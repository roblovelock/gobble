package main

import (
	"fmt"
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strconv"
	"strings"
	"testing"
)

func benchmarkParseString(data string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = parseString(strings.NewReader(data))
	}
}
func BenchmarkParseStringSimple(b *testing.B) { benchmarkParseString("\"abc\"", b) }
func BenchmarkParseString(b *testing.B) {
	benchmarkParseString(
		"\"tab:\\tafter tab, newline:\\nnew line, quote: \\\", emoji: \\u{1F602}, newline:\\nescaped whitespace: \\    abc\"",
		b,
	)
}

func benchmarkUnquote(data string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = strconv.Unquote(data)
	}
}

func BenchmarkUnquoteSimple(b *testing.B) { benchmarkUnquote("\"abc\"", b) }
func BenchmarkUnquote(b *testing.B) {
	benchmarkUnquote(
		"\"tab:\\tafter tab, newline:\\nnew line, quote: \\\", emoji: \\u{1F602}, newline:\\nescaped whitespace: \\    abc\"",
		b,
	)
}

func Test_strconv_Unquote(t *testing.T) {
	data := "\"abc\""

	want, err := strconv.Unquote(data)
	require.NoError(t, err)
	got, err := parseString(strings.NewReader(data))
	require.NoError(t, err)
	assert.Equal(t, want, got)
}

func Test_parseString(t *testing.T) {
	type args struct {
		in parser.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "basic",
			args:    args{in: strings.NewReader(`"abc"`)},
			want:    "abc",
			wantErr: assert.NoError,
		},
		{
			name:    "tab",
			args:    args{in: strings.NewReader(`"tab:\tafter tab"`)},
			want:    "tab:\tafter tab",
			wantErr: assert.NoError,
		},
		{
			name:    "newline",
			args:    args{in: strings.NewReader(`"newline:\nnew line"`)},
			want:    "newline:\nnew line",
			wantErr: assert.NoError,
		},
		{
			name:    "quote",
			args:    args{in: strings.NewReader(`"quote: \""`)},
			want:    "quote: \"",
			wantErr: assert.NoError,
		},
		{
			name:    "emoji",
			args:    args{in: strings.NewReader(`"emoji: \u{1F602}"`)},
			want:    "emoji: 😂",
			wantErr: assert.NoError,
		},
		{
			name:    "emoji",
			args:    args{in: strings.NewReader(`"escaped whitespace: \    abc"`)},
			want:    "escaped whitespace: abc",
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseString(tt.args.in)
			if !tt.wantErr(t, err, fmt.Sprintf("parseString(%v)", tt.args.in)) {
				return
			}
			assert.Equalf(t, tt.want, got, "parseString(%v)", tt.args.in)
		})
	}
}
