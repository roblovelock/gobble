package bits

import (
	"errors"
	"gobble/internal/parser"
	"io"
)

const (
	MaxBits    = 64
	BitsInByte = 8
)

var (
	ErrRemainingBits = errors.New("remaining bits")
	ErrBitsOverflow  = errors.New("bits overflow")
)

type (
	bitReader struct {
		parser.Reader
		cache byte  // unread bits
		bits  uint8 // number of unread bits in cache
	}
)

func (r *bitReader) Read(p []byte) (n int, err error) {
	if r.bits == 0 {
		return r.Reader.Read(p)
	}

	for ; n < len(p); n++ {
		if p[n], err = r.readByte(); err != nil {
			return
		}
	}

	return
}

func (r *bitReader) Seek(offset int64, whence int) (int64, error) {
	offsetBytes := offset / BitsInByte
	offsetBits := uint8(offset % BitsInByte)
	switch whence {
	case io.SeekStart:
		r.bits = 0
	case io.SeekCurrent:
		if r.bits > 0 {
			offsetBytes--
			offsetBits += BitsInByte - r.bits
			r.bits = 0
		}
	case io.SeekEnd:
		if offsetBits > 0 {
			offsetBytes++
			offsetBits = BitsInByte - offsetBits
		}
		offsetBytes *= -1
	default:
		return 0, errors.New("bytes.Reader.Seek: invalid whence")
	}

	o, err := r.Reader.Seek(offsetBytes, whence)
	o *= BitsInByte
	if err != nil {
		return o, err
	}
	if offsetBits > 0 {
		_, n, err := r.ReadBits(offsetBits)
		o += int64(n)
		if err != nil {
			return o, err
		}
	}

	return o, nil
}

func (r *bitReader) ReadByte() (b byte, err error) {
	// r.bits will be the same after reading 8 bits, so we don't need to update that.
	if r.bits == 0 {
		return r.Reader.ReadByte()
	}
	return r.readByte()
}

func (r *bitReader) ReadBool() (b bool, err error) {
	if r.bits == 0 {
		r.cache, err = r.Reader.ReadByte()
		if err != nil {
			return
		}
		r.bits = BitsInByte
	}

	r.bits--
	b = (r.cache & (1 << r.bits)) != 0
	r.cache &= 1<<r.bits - 1
	return
}

func (r *bitReader) readByte() (b byte, err error) {
	b = r.cache << (BitsInByte - r.bits)
	r.cache, err = r.Reader.ReadByte()
	if err != nil {
		return 0, err
	}
	b |= r.cache >> r.bits
	r.cache &= 1<<r.bits - 1
	return
}

func (r *bitReader) ReadBits(n uint8) (uint64, uint8, error) {
	if n < r.bits {
		r.bits -= n
		u := uint64(r.cache >> r.bits)
		r.cache &= 1<<r.bits - 1
		return u, n, nil
	}

	if n > r.bits {
		var u uint64
		if n > MaxBits {
			n = MaxBits
		}
		remainder := n
		if r.bits > 0 {
			u = uint64(r.cache)
			remainder -= r.bits
		}

		for ; remainder >= BitsInByte; remainder -= BitsInByte {
			b, err := r.Reader.ReadByte()
			if err != nil {
				return u, n - remainder, err
			}
			u = u<<BitsInByte + uint64(b)
		}

		if remainder > 0 {
			var err error
			if r.cache, err = r.Reader.ReadByte(); err != nil {
				return u, n - remainder, err
			}
			r.bits = BitsInByte - remainder
			u = u<<remainder + uint64(r.cache>>r.bits)
			r.cache &= 1<<r.bits - 1
		} else {
			r.bits = 0
		}
		return u, n, nil
	}

	r.bits = 0
	return uint64(r.cache), n, nil
}

func (r *bitReader) isAligned() bool {
	return r.bits == 0
}

func Bits[T any](p parser.Parser[parser.BitReader, T]) parser.Parser[parser.Reader, T] {
	return func(in parser.Reader) (T, error) {
		currentOffset, _ := in.Seek(0, io.SeekCurrent)
		reader := &bitReader{Reader: in}
		result, err := p(reader)
		if err != nil {
			var t T
			return t, err
		}

		if !reader.isAligned() {
			_, _ = in.Seek(currentOffset, io.SeekStart)
			return result, ErrRemainingBits
		}

		return result, nil
	}
}
