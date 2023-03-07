package main

import (
	"encoding/json"
	"os"
	"testing"
)

func benchmarkGoJSON(b *testing.B, text string) {
	for i := 0; i < b.N; i++ {
		var val interface{}
		_ = json.Unmarshal([]byte(text), &val)
		sink(val)
	}
	b.SetBytes(int64(len(text)))
}

func BenchmarkGoJSONInt(b *testing.B) {
	benchmarkGoJSON(b, intText)
}

func BenchmarkGoJSONFloat(b *testing.B) {
	benchmarkGoJSON(b, floatText)
}

func BenchmarkGoJSONString(b *testing.B) {
	benchmarkGoJSON(b, stringText)
}

func BenchmarkGoJSONBool(b *testing.B) {
	benchmarkGoJSON(b, boolText)
}

func BenchmarkGoJSONNull(b *testing.B) {
	benchmarkGoJSON(b, nullText)
}

func BenchmarkGoJSONArray(b *testing.B) {
	benchmarkGoJSON(b, arrayText)
}

func BenchmarkGoJSONMap(b *testing.B) {
	benchmarkGoJSON(b, objText)
}

func BenchmarkGoJSONMedium(b *testing.B) {
	text, err := os.ReadFile("testdata/medium.json")
	if err != nil {
		b.Error(err)
	}
	benchmarkGoJSON(b, string(text))
}

func BenchmarkGoJSONLarge(b *testing.B) {
	text, err := os.ReadFile("testdata/large.json")
	if err != nil {
		b.Error(err)
	}
	benchmarkGoJSON(b, string(text))
}
