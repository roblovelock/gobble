package bits

import (
	"bytes"
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"math"
	"testing"
)

func Test_bitReader_Read(t *testing.T) {
	type fields struct {
		data  []byte
		cache byte
		bits  byte
	}
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantN   int
		wantErr error
	}{
		{
			name:    "EOF if no data remains",
			fields:  fields{data: []byte{}, bits: 0},
			args:    args{p: make([]byte, 1)},
			want:    make([]byte, 1),
			wantN:   0,
			wantErr: io.EOF,
		},
		{
			name:    "read multiple unaligned byte, EOF if no data remains",
			fields:  fields{data: []byte{2, 3}, cache: 0xFF, bits: 7},
			args:    args{p: make([]byte, 3)},
			want:    []byte{0xFE, 0x04, 0x00},
			wantN:   2,
			wantErr: io.EOF,
		},
		{
			name:   "read single byte",
			fields: fields{data: []byte{1}, bits: 0},
			args:   args{p: make([]byte, 1)},
			want:   []byte{1},
			wantN:  1,
		},
		{
			name:   "read multi bytes",
			fields: fields{data: []byte{1, 2, 3, 4, 5}, bits: 0},
			args:   args{p: make([]byte, 4)},
			want:   []byte{1, 2, 3, 4},
			wantN:  4,
		},
		{
			name:   "read single unaligned byte",
			fields: fields{data: []byte{3}, cache: 1, bits: 1},
			args:   args{p: make([]byte, 1)},
			want:   []byte{0x81},
			wantN:  1,
		},
		{
			name:   "read multiple unaligned byte",
			fields: fields{data: []byte{2, 3}, cache: 0xFF, bits: 7},
			args:   args{p: make([]byte, 2)},
			want:   []byte{0xFE, 0x04},
			wantN:  2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &bitReader{
				Reader: bytes.NewReader(tt.fields.data),
				cache:  tt.fields.cache,
				bits:   tt.fields.bits,
			}
			gotN, err := r.Read(tt.args.p)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.wantN, gotN)
			assert.Equal(t, tt.want, tt.args.p)
		})
	}
}

func Test_bitReader_ReadByte(t *testing.T) {
	type fields struct {
		data  []byte
		cache byte
		bits  byte
	}
	tests := []struct {
		name    string
		fields  fields
		want    byte
		wantErr error
	}{
		{
			name:    "EOF if no data remains",
			fields:  fields{data: []byte{}, bits: 0},
			want:    0,
			wantErr: io.EOF,
		},
		{
			name:   "read single byte",
			fields: fields{data: []byte{1}, bits: 0},
			want:   1,
		},
		{
			name:   "read single unaligned byte",
			fields: fields{data: []byte{3}, cache: 1, bits: 1},
			want:   0x81,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &bitReader{
				Reader: bytes.NewReader(tt.fields.data),
				cache:  tt.fields.cache,
				bits:   tt.fields.bits,
			}
			got, err := r.ReadByte()
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_bitReader_isAligned(t *testing.T) {
	type fields struct {
		Reader parser.Reader
		cache  byte
		bits   byte
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "true if remaining bits == 0",
			fields: fields{bits: 0},
			want:   true,
		},
		{
			name:   "false if remaining bits != 0",
			fields: fields{bits: 1},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &bitReader{bits: tt.fields.bits}
			assert.Equal(t, tt.want, r.isAligned())
		})
	}
}

func Test_bitReader_ReadBool(t *testing.T) {
	type fields struct {
		data  []byte
		cache byte
		bits  byte
	}
	tests := []struct {
		name    string
		fields  fields
		want    bool
		wantErr error
	}{
		{
			name:    "EOF if no data remains",
			fields:  fields{data: []byte{}, bits: 0},
			want:    false,
			wantErr: io.EOF,
		},
		{
			name:   "read single zero bit",
			fields: fields{data: []byte{0x7F}, bits: 0},
			want:   false,
		},
		{
			name:   "read single one bit",
			fields: fields{data: []byte{0x80}, bits: 0},
			want:   true,
		},
		{
			name:   "read single unaligned zero bit",
			fields: fields{bits: 1, cache: 0xFE},
			want:   false,
		},
		{
			name:   "read single unaligned one bit",
			fields: fields{bits: 1, cache: 0x01},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &bitReader{
				Reader: bytes.NewReader(tt.fields.data),
				cache:  tt.fields.cache,
				bits:   tt.fields.bits,
			}
			got, err := r.ReadBool()
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_bitReader_ReadBits(t *testing.T) {
	type fields struct {
		data  []byte
		cache byte
		bits  byte
	}
	type args struct {
		bits uint8
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint64
		wantN   uint8
		wantErr error
	}{
		{
			name:    "EOF if no data remains",
			fields:  fields{data: []byte{}, bits: 0},
			args:    args{bits: 1},
			want:    0,
			wantN:   0,
			wantErr: io.EOF,
		},
		{
			name:    "EOF if no data remains from multiple bytes",
			fields:  fields{data: []byte{0xFF}, bits: 0},
			args:    args{bits: 9},
			want:    0xFF,
			wantN:   8,
			wantErr: io.EOF,
		},
		{
			name:    "read multiple unaligned bits, EOF if no data remains",
			fields:  fields{data: []byte{}, cache: 0x01, bits: 1},
			args:    args{bits: 2},
			want:    1,
			wantN:   1,
			wantErr: io.EOF,
		},
		{
			name:    "read multiple unaligned bits from multiple bytes, EOF if no data remains",
			fields:  fields{data: []byte{0xFF}, cache: 0x01, bits: 1},
			args:    args{bits: 10},
			want:    0x1FF,
			wantN:   9,
			wantErr: io.EOF,
		},
		{
			name:    "read to EOF",
			fields:  fields{data: []byte{0xFF}, cache: 0x7F, bits: 7},
			args:    args{bits: 255},
			want:    0x7fff,
			wantN:   15,
			wantErr: io.EOF,
		},
		{
			name:   "read single bit",
			fields: fields{data: []byte{0x80}, bits: 0},
			args:   args{bits: 1},
			want:   1,
			wantN:  1,
		},
		{
			name:   "read multi bits",
			fields: fields{data: []byte{1, 1}, bits: 0},
			args:   args{bits: 16},
			want:   0x101,
			wantN:  16,
		},
		{
			name:   "read all bits in cache",
			fields: fields{cache: 1, bits: 1},
			args:   args{bits: 1},
			want:   1,
			wantN:  1,
		},
		{
			name:   "read single unaligned bit",
			fields: fields{cache: 0x03, bits: 2},
			args:   args{bits: 1},
			want:   1,
			wantN:  1,
		},
		{
			name:   "read multiple unaligned bits",
			fields: fields{data: []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x80}, cache: 0x7F, bits: 7},
			args:   args{bits: 64},
			want:   math.MaxUint64,
			wantN:  64,
		},
		{
			name:   "read max 64 bits",
			fields: fields{data: []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x80}, cache: 0x7F, bits: 7},
			args:   args{bits: 255},
			want:   math.MaxUint64,
			wantN:  64,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &bitReader{
				Reader: bytes.NewReader(tt.fields.data),
				cache:  tt.fields.cache,
				bits:   tt.fields.bits,
			}
			got, gotN, err := r.ReadBits(tt.args.bits)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.wantN, gotN)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_bitReader_Seek(t *testing.T) {
	type fields struct {
		data []byte
		bits byte
	}
	type args struct {
		offset int64
		whence int
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      int64
		wantCache byte
		wantBits  uint8
		wantErr   error
	}{
		{
			name:    "EOF if seek 1 bit past end",
			fields:  fields{data: []byte{}, bits: 0},
			args:    args{offset: 1, whence: io.SeekStart},
			want:    0,
			wantErr: io.EOF,
		},
		{
			name:    "EOF if seek before start",
			fields:  fields{data: []byte{}, bits: 0},
			args:    args{offset: -1, whence: io.SeekStart},
			want:    0,
			wantErr: io.EOF,
		},
		{
			name:   "seek zero bytes from start",
			fields: fields{data: []byte{0xFF}, bits: 0},
			args:   args{offset: 0, whence: io.SeekStart},
			want:   0,
		},
		{
			name:      "seek 1 bit from start",
			fields:    fields{data: []byte{0xFF}, bits: 0},
			args:      args{offset: 1, whence: io.SeekStart},
			want:      1,
			wantBits:  7,
			wantCache: 0x7F,
		},
		{
			name:      "seek  1 byte, 1 bit from start",
			fields:    fields{data: []byte{0xFF, 0xFF}, bits: 0},
			args:      args{offset: 9, whence: io.SeekStart},
			want:      9,
			wantBits:  7,
			wantCache: 0x7F,
		},
		{
			name:      "seek 1 bit from end",
			fields:    fields{data: []byte{0xFF}, bits: 0},
			args:      args{offset: 1, whence: io.SeekEnd},
			want:      7,
			wantBits:  1,
			wantCache: 0x01,
		},
		{
			name:      "seek  1 byte, 1 bit from end",
			fields:    fields{data: []byte{0xFF, 0xFF}, bits: 0},
			args:      args{offset: 9, whence: io.SeekEnd},
			want:      7,
			wantBits:  1,
			wantCache: 0x01,
		},
		{
			name:      "seek 1 bit from current at start",
			fields:    fields{data: []byte{0xFF}, bits: 0},
			args:      args{offset: 1, whence: io.SeekCurrent},
			want:      1,
			wantBits:  7,
			wantCache: 0x7F,
		},
		{
			name:      "seek  1 byte, 1 bit from current at start",
			fields:    fields{data: []byte{0xFF, 0xFF}, bits: 0},
			args:      args{offset: 9, whence: io.SeekCurrent},
			want:      9,
			wantBits:  7,
			wantCache: 0x7F,
		},
		{
			name:      "seek 1 bit from current at one bit in",
			fields:    fields{data: []byte{0xFF}, bits: 1},
			args:      args{offset: 1, whence: io.SeekCurrent},
			want:      2,
			wantBits:  6,
			wantCache: 0x3F,
		},
		{
			name:      "seek  1 byte, from current at one bit in",
			fields:    fields{data: []byte{0xFF, 0xFF}, bits: 1},
			args:      args{offset: 8, whence: io.SeekCurrent},
			want:      9,
			wantBits:  7,
			wantCache: 0x7F,
		},
		{
			name:      "seek 1 byte, 1 bit from current at one bit in",
			fields:    fields{data: []byte{0xFF, 0xFF}, bits: 1},
			args:      args{offset: 9, whence: io.SeekCurrent},
			want:      10,
			wantBits:  6,
			wantCache: 0x3F,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &bitReader{Reader: bytes.NewReader(tt.fields.data)}
			if tt.fields.bits > 0 {
				_, _, err := r.ReadBits(tt.fields.bits)
				require.NoError(t, err)
			}
			got, err := r.Seek(tt.args.offset, tt.args.whence)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantCache, r.cache)
			assert.Equal(t, tt.wantBits, r.bits)
		})
	}
}
