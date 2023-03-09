package main

import (
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/stretchr/testify/assert"
	"image/color"
	"io"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name    string
		args    args
		want    color.Color
		wantErr error
	}{
		{
			name:    "empty input => EOF",
			args:    args{in: ""},
			wantErr: io.EOF,
		},
		{
			name:    "invalid input => EOF",
			args:    args{in: "$"},
			wantErr: errors.ErrNotMatched,
		},
		{
			name: "hex input => color",
			args: args{in: "#123456"},
			want: color.RGBA{R: 18, G: 52, B: 86, A: 255},
		},
		{
			name: "hex black => color",
			args: args{in: "#000000"},
			want: color.RGBA{R: 0, G: 0, B: 0, A: 255},
		},
		{
			name: "hex white => color",
			args: args{in: "#FFF"},
			want: color.RGBA{R: 255, G: 255, B: 255, A: 255},
		},
		{
			name: "short hex => color",
			args: args{in: "#CDE"},
			want: color.RGBA{R: 204, G: 221, B: 238, A: 255},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.in)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}
