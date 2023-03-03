package main

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gobble/pkg/parser"
	"image"
	"image/png"
	"os"
	"testing"
)

func Test_parseImage(t *testing.T) {
	type args struct {
		in parser.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    image.Image
		wantErr error
	}{
		{
			name: "dice",
			args: args{in: mustReadQOI(t, "dice.qoi")},
			want: mustReadPNG(t, "dice.png"),
		},
		{
			name: "kodim10",
			args: args{in: mustReadQOI(t, "kodim10.qoi")},
			want: mustReadPNG(t, "kodim10.png"),
		},
		{
			name: "kodim23",
			args: args{in: mustReadQOI(t, "kodim23.qoi")},
			want: mustReadPNG(t, "kodim23.png"),
		},
		{
			name: "qoi_logo",
			args: args{in: mustReadQOI(t, "qoi_logo.qoi")},
			want: mustReadPNG(t, "qoi_logo.png"),
		},
		{
			name: "testcard",
			args: args{in: mustReadQOI(t, "testcard.qoi")},
			want: mustReadPNG(t, "testcard.png"),
		},
		{
			name: "testcard_rgba",
			args: args{in: mustReadQOI(t, "testcard_rgba.qoi")},
			want: mustReadPNG(t, "testcard_rgba.png"),
		},
		{
			name: "wikipedia_008",
			args: args{in: mustReadQOI(t, "wikipedia_008.qoi")},
			want: mustReadPNG(t, "wikipedia_008.png"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseImage(tt.args.in)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want.Bounds(), got.Bounds())

			for x := 0; x < tt.want.Bounds().Dx(); x++ {
				for y := 0; y < tt.want.Bounds().Dy(); y++ {
					w := [4]uint32{}
					g := [4]uint32{}
					w[0], w[1], w[2], w[3] = tt.want.At(x, y).RGBA()
					g[0], g[1], g[2], g[3] = got.At(x, y).RGBA()

					if !assert.Equal(t, w, g) {
						fmt.Printf("x:%d, y:%d, i:%d\n", x, y, x+y*tt.want.Bounds().Dx())

						for i := -10; i < 10; i++ {
							dx := x + i
							dy := y
							if x < 0 {
								dx = tt.want.Bounds().Dx() - 1
								dy--
							} else if x >= tt.want.Bounds().Dx() {
								dx = 0
								dy++
							}

							fmt.Printf("%v, %v\n", tt.want.At(dx, dy), got.At(dx, dy))
						}

						require.Fail(t, "Failed")
					}
				}
			}

		})
	}
}

func mustReadQOI(t *testing.T, filename string) parser.Reader {
	data, err := os.ReadFile(fmt.Sprintf("testdata/%s", filename))
	require.NoError(t, err)
	return bytes.NewReader(data)
}

func mustReadPNG(t *testing.T, filename string) image.Image {
	file, err := os.Open(fmt.Sprintf("testdata/%s", filename))
	require.NoError(t, err)
	pngImg, err := png.Decode(file)
	require.NoError(t, err)
	return pngImg
}
