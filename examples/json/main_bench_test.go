package main

import (
	"os"
	"testing"
)

const (
	intText    = "1000"
	floatText  = "-10.11231"
	stringText = `"hello world"`
	boolText   = "true"
	nullText   = "null"
	arrayText  = `[79.1253026404128,null,"hello world", false, 10]`
	objText    = `{ 
		"inelegant":27.53096820876087,
		"horridness":true,
		"iridodesis":[79.1253026404128,null,"hello world", false, 10],
		"arrogantness":null,
		"unagrarian":false
	}`
)

//go:noinline
func sink[T any](x T) {
}

func benchmarkJSON(b *testing.B, json string) {
	for i := 0; i < b.N; i++ {
		val, _ := ParseJSON(json)
		sink(val)
	}
	b.SetBytes(int64(len(json)))
}

func BenchmarkJSONInt(b *testing.B) {
	benchmarkJSON(b, intText)

}

func BenchmarkJSONFloat(b *testing.B) {
	benchmarkJSON(b, floatText)
}

func BenchmarkJSONString(b *testing.B) {
	benchmarkJSON(b, stringText)
}

func BenchmarkJSONBool(b *testing.B) {
	benchmarkJSON(b, boolText)
}

func BenchmarkJSONNull(b *testing.B) {
	benchmarkJSON(b, nullText)
}

func BenchmarkJSONArray(b *testing.B) {
	benchmarkJSON(b, arrayText)
}

func BenchmarkJSONMap(b *testing.B) {
	benchmarkJSON(b, objText)
}

func BenchmarkJSONMedium(b *testing.B) {
	text, err := os.ReadFile("testdata/medium.json")
	if err != nil {
		b.Error(err)
	}
	benchmarkJSON(b, string(text))
}

func BenchmarkJSONLarge(b *testing.B) {
	text, err := os.ReadFile("testdata/large.json")
	if err != nil {
		b.Error(err)
	}
	benchmarkJSON(b, string(text))
}
