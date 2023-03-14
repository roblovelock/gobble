package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
	text := "4 + 123 + 23 + 67 +89 + 87 *78\n/67-98-		 199"
	for i := 0; i < b.N; i++ {
		val, err := ParseExpr(text)
		assert.NoError(b, err)
		assert.Equal(b, int64(110), val)
	}
	b.SetBytes(int64(len(text)))
}

func BenchmarkBytesExpr(b *testing.B) {
	text := "4 + 123 + 23 + 67 +89 + 87 *78\n/67-98-		 199"
	for i := 0; i < b.N; i++ {
		val, _, err := ParseBytesExpr(text)
		assert.NoError(b, err)
		assert.Equal(b, int64(110), val)
	}
	b.SetBytes(int64(len(text)))
}
