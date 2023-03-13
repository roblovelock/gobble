package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseExpr(t *testing.T) {
	type args struct {
		expr string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr error
	}{
		{
			name: "test add",
			args: args{expr: "1 + 2"},
			want: 3,
		},
		{
			name: "test subtract expression",
			args: args{expr: " ( 1 - 2 ) "},
			want: -1,
		},
		{
			name: "test multiply",
			args: args{expr: "2  *  3"},
			want: 6,
		},
		{
			name: "test add and multiply",
			args: args{expr: "19 + 10 * 20"},
			want: 219,
		},
		{
			name: "test add expression then multiply",
			args: args{expr: "(19 + 10) * 20"},
			want: 580,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseExpr(tt.args.expr)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}

func BenchmarkExpr1Op(b *testing.B) {
	text := "19 + 10"
	for i := 0; i < b.N; i++ {
		_, _ = ParseExpr(text)
	}
	b.SetBytes(int64(len(text)))
}

func BenchmarkExpr2Op(b *testing.B) {
	text := "19+10*20"
	for i := 0; i < b.N; i++ {
		_, _ = ParseExpr(text)
	}
	b.SetBytes(int64(len(text)))
}

func BenchmarkExpr3Op(b *testing.B) {
	text := "19 + 10 * 20/9"
	for i := 0; i < b.N; i++ {
		_, _ = ParseExpr(text)
	}
	b.SetBytes(int64(len(text)))
}

func BenchmarkExpr(b *testing.B) {
	text := `4 + 123 + 23 + 67 +89 + 87 *78
/67-98-		 199`
	for i := 0; i < b.N; i++ {
		_, _ = ParseExpr(text)
	}
	b.SetBytes(int64(len(text)))
}
