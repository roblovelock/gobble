package main

import (
	"github.com/roblovelock/gobble/pkg/combinator"
	"github.com/roblovelock/gobble/pkg/combinator/branch"
	"github.com/roblovelock/gobble/pkg/combinator/sequence"
	"github.com/roblovelock/gobble/pkg/parser/ascii"
	"github.com/roblovelock/gobble/pkg/parser/bytes"
	"image/color"
	"strconv"
	"strings"
)

var (
	hex       = bytes.TakeWhileMinMax(2, 2, ascii.IsHexDigit)
	rgbValue  = combinator.Map(hex, hexToUint8)
	rgbValues = sequence.Tuple(rgbValue, rgbValue, rgbValue)

	hexByte        = bytes.OneOf([]byte("0123456789abcdefABCDEF")...)
	shortRGBValue  = combinator.Map(hexByte, hexByteToUint8)
	shortRGBValues = sequence.Tuple(shortRGBValue, shortRGBValue, shortRGBValue)

	rgb = branch.Alt(rgbValues, shortRGBValues)

	colorValue = combinator.Map(rgb, func(rgb []uint8) (color.Color, error) {
		return color.RGBA{R: rgb[0], G: rgb[1], B: rgb[2], A: 255}, nil
	})

	hash        = bytes.Byte('#')
	colorParser = sequence.Preceded(hash, colorValue)
)

func hexToUint8(bytes []byte) (uint8, error) {
	i, err := strconv.ParseUint(string(bytes), 16, 8)
	return uint8(i), err
}

func hexByteToUint8(hexByte byte) (uint8, error) {
	return hexToUint8([]byte{hexByte, hexByte})
}

func Parse(in string) (color.Color, error) {
	return colorParser(strings.NewReader(in))
}
