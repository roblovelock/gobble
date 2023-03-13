package bytes_test

import (
	"github.com/roblovelock/gobble/pkg/parser/ascii"
	"github.com/roblovelock/gobble/pkg/parser/bytes"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func BenchmarkTakeWhile(b *testing.B) {
	parser := bytes.TakeWhile(ascii.IsDigit)
	for i := 0; i < b.N; i++ {
		_, err := parser.Parse(strings.NewReader("123456789"))
		assert.NoError(b, err)
	}
}
