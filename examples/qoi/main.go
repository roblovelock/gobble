package main

import (
	bytes2 "bytes"
	"fmt"
	"github.com/roblovelock/gobble/pkg/combinator/branch"
	"github.com/roblovelock/gobble/pkg/combinator/modifier"
	"github.com/roblovelock/gobble/pkg/combinator/multi"
	"github.com/roblovelock/gobble/pkg/combinator/sequence"
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/roblovelock/gobble/pkg/parser/bits"
	"github.com/roblovelock/gobble/pkg/parser/bytes"
	"github.com/roblovelock/gobble/pkg/parser/numeric"
	"github.com/roblovelock/gobble/pkg/parser/stream"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

type pixelContext struct {
	c color.NRGBA
	p [64]color.NRGBA
	i image.Image
}

func (p *pixelContext) setColor(c color.NRGBA) {
	p.c = c
	p.p[hash(c)] = c
}

func main() {
	qoiImage, err := os.ReadFile("examples/qoi/testdata/qoi_logo.qoi")
	if err != nil {
		log.Fatal(err)
	}

	img, err := parseImage(bytes2.NewReader(qoiImage))
	if err != nil {
		log.Fatal(err)
	}

	out, err := os.Create("output.png")
	if err != nil {
		log.Fatal(err)
	}

	err = png.Encode(out, img)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Generated image to output.png")
}

func parseImage(in parser.Reader) (image.Image, error) {
	return modifier.Map(
		sequence.Pair(
			headerParser(),
			pixelParser(),
		),
		func(p parser.Pair[image.Rectangle, []color.Color]) (image.Image, error) {
			img := image.NewNRGBA(p.First)
			for i := 0; i < len(p.Second); i++ {
				img.Set(i%p.First.Dx(), i/p.First.Dx(), p.Second[i])
			}
			return img, nil
		},
	)(in)
}

func headerParser() parser.Parser[parser.Reader, image.Rectangle] {
	return modifier.Map(
		sequence.Preceded(
			bytes.Tag([]byte("qoif")),
			sequence.Terminated(
				sequence.Tuple[parser.Reader, uint32](numeric.UInt32BE(), numeric.UInt32BE()),
				bytes.Take(2),
			),
		),
		func(i []uint32) (image.Rectangle, error) {
			return image.Rect(0, 0, int(i[0]), int(i[1])), nil
		},
	)
}

func endParser() parser.Parser[parser.Reader, parser.Empty] {
	return sequence.Preceded(
		bytes.Tag([]byte{0, 0, 0, 0, 0, 0, 0, 1}),
		stream.EOF(),
	)
}

func rgbParser(ctx *pixelContext) parser.Parser[parser.Reader, []color.Color] {
	return sequence.Preceded(
		bytes.Byte(0xFE),
		modifier.Map(
			bytes.Take(3),
			func(b []byte) ([]color.Color, error) {
				ctx.setColor(color.NRGBA{R: b[0], G: b[1], B: b[2], A: ctx.c.A})
				return []color.Color{ctx.c}, nil
			},
		),
	)
}

func rgbaParser(ctx *pixelContext) parser.Parser[parser.Reader, []color.Color] {
	return sequence.Preceded(
		bytes.Byte(0xFF),
		modifier.Map(
			bytes.Take(4),
			func(b []byte) ([]color.Color, error) {
				ctx.setColor(color.NRGBA{R: b[0], G: b[1], B: b[2], A: b[3]})
				return []color.Color{ctx.c}, nil
			},
		),
	)
}

func indexParser(ctx *pixelContext) parser.Parser[parser.Reader, []color.Color] {
	return bits.Bits(
		sequence.Preceded(
			bits.Tag[uint8](2, 0x00),
			modifier.Map[parser.BitReader, uint8, []color.Color](
				bits.Take[uint8](6),
				func(i uint8) ([]color.Color, error) {
					ctx.setColor(ctx.p[i])
					return []color.Color{ctx.c}, nil
				},
			),
		),
	)
}

func diffParser(ctx *pixelContext) parser.Parser[parser.Reader, []color.Color] {
	return bits.Bits(
		sequence.Preceded(
			bits.Tag[uint8](2, 0x01),
			modifier.Map(
				sequence.Tuple[parser.BitReader, uint8](
					bits.Take[uint8](2), bits.Take[uint8](2), bits.Take[uint8](2),
				),
				func(i []uint8) ([]color.Color, error) {
					ctx.setColor(color.NRGBA{
						R: ctx.c.R + i[0] - 2, G: ctx.c.G + i[1] - 2, B: ctx.c.B + i[2] - 2, A: ctx.c.A,
					})
					return []color.Color{ctx.c}, nil
				},
			),
		),
	)
}

func lumaParser(ctx *pixelContext) parser.Parser[parser.Reader, []color.Color] {
	return bits.Bits(
		sequence.Preceded(
			bits.Tag[uint8](2, 0x02),
			modifier.Map(
				sequence.Tuple[parser.BitReader, uint8](
					bits.Take[uint8](6), bits.Take[uint8](4), bits.Take[uint8](4),
				),
				func(i []uint8) ([]color.Color, error) {
					dg := i[0] - 32
					dr := i[1] - 8 + dg
					db := i[2] - 8 + dg

					ctx.setColor(color.NRGBA{R: ctx.c.R + dr, G: ctx.c.G + dg, B: ctx.c.B + db, A: ctx.c.A})
					return []color.Color{ctx.c}, nil
				},
			),
		),
	)
}

func runParser(ctx *pixelContext) parser.Parser[parser.Reader, []color.Color] {
	return bits.Bits(
		sequence.Preceded(
			bits.Tag[uint8](2, 0x03),
			modifier.Map(
				bits.Take[uint8](6),
				func(count uint8) ([]color.Color, error) {
					c := make([]color.Color, count+1)
					for i := uint8(0); i <= count; i++ {
						c[i] = ctx.c
					}
					return c, nil
				},
			),
		),
	)
}

func pixelParser() parser.Parser[parser.Reader, []color.Color] {
	ctx := pixelContext{
		c: color.NRGBA{A: 255},
	}

	return multi.FoldMany0(
		branch.Alt(
			modifier.Value(endParser(), []color.Color{}),
			rgbParser(&ctx),
			rgbaParser(&ctx),
			indexParser(&ctx),
			diffParser(&ctx),
			lumaParser(&ctx),
			runParser(&ctx),
		),
		make([]color.Color, 0),
		func(img []color.Color, c []color.Color) []color.Color {
			return append(img, c...)
		},
	)
}

func hash(c color.NRGBA) uint8 {
	return (c.R*3 + c.G*5 + c.B*7 + c.A*11) % 64
}
