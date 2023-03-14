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

func TestParseBytesExpr(t *testing.T) {
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
			got, out, err := ParseBytesExpr(tt.args.expr)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Empty(t, out)
			assert.Equal(t, tt.want, got)
		})
	}
}
