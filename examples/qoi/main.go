package main

import (
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

func parseImage(in parser.Reader) (image.Image, error) {
	return imageParser().Parse(in)
}

func imageParser() parser.Parser[parser.Reader, image.Image] {
	return modifier.Map(
		sequence.Terminated(
			sequence.Pair(
				headerParser(),
				pixelParser(),
			),
			endParser(),
		),
		func(p parser.Pair[image.Rectangle, []color.Color]) (image.Image, error) {
			img := image.NewNRGBA(p.First)
			for i := 0; i < len(p.Second); i++ {
				img.Set(i%p.First.Dx(), i/p.First.Dx(), p.Second[i])
			}
			return img, nil
		},
	)
}

func headerParser() parser.Parser[parser.Reader, image.Rectangle] {
	return modifier.Map(
		sequence.Preceded(
			bytes.Tag([]byte("qoif")),
			sequence.Terminated(
				sequence.Tuple[parser.Reader, uint32](numeric.Uint32BE(), numeric.Uint32BE()),
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

func indexParser(ctx *pixelContext) parser.Parser[parser.BitReader, []color.Color] {
	return modifier.Map[parser.BitReader, uint8, []color.Color](
		branch.Alt(
			sequence.Terminated(
				bits.Tag[uint8](6, 0),
				modifier.Cut(modifier.Not(bits.Tag[uint8](8, 0))),
			),
			bits.Take[uint8](6),
		),
		func(i uint8) ([]color.Color, error) {
			ctx.setColor(ctx.p[i])
			return []color.Color{ctx.c}, nil
		},
	)
}

func diffParser(ctx *pixelContext) parser.Parser[parser.BitReader, []color.Color] {
	return modifier.Map(
		sequence.Tuple[parser.BitReader, uint8](
			bits.Take[uint8](2), bits.Take[uint8](2), bits.Take[uint8](2),
		),
		func(i []uint8) ([]color.Color, error) {
			ctx.setColor(color.NRGBA{
				R: ctx.c.R + i[0] - 2, G: ctx.c.G + i[1] - 2, B: ctx.c.B + i[2] - 2, A: ctx.c.A,
			})
			return []color.Color{ctx.c}, nil
		},
	)
}

func lumaParser(ctx *pixelContext) parser.Parser[parser.BitReader, []color.Color] {
	return modifier.Map(
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
	)
}

func runParser(ctx *pixelContext) parser.Parser[parser.BitReader, []color.Color] {
	return modifier.Map(
		bits.Take[uint8](6),
		func(count uint8) ([]color.Color, error) {
			c := make([]color.Color, count+1)
			for i := uint8(0); i <= count; i++ {
				c[i] = ctx.c
			}
			return c, nil
		},
	)
}

func bitPixelParser(ctx *pixelContext) parser.Parser[parser.Reader, []color.Color] {
	parsers := map[uint8]parser.Parser[parser.BitReader, []color.Color]{
		0x00: indexParser(ctx),
		0x01: diffParser(ctx),
		0x02: lumaParser(ctx),
		0x03: runParser(ctx),
	}

	return bits.Bits(branch.Case(bits.Take[uint8](2), parsers))
}

func pixelParser() parser.Parser[parser.Reader, []color.Color] {
	ctx := pixelContext{
		c: color.NRGBA{A: 255},
	}

	return multi.FoldMany0(
		branch.Alt(
			rgbParser(&ctx),
			rgbaParser(&ctx),
			bitPixelParser(&ctx),
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
