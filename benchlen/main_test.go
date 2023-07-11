package main

import (
	"testing"
	"unicode/utf8"

	"github.com/rivo/uniseg"
)

const prideFlag = "🏳️‍🌈"

var prideFlagRune = []rune("🏳️‍🌈")
var prideFlagBytes = []byte("🏳️‍🌈")

func BenchmarkLen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = len(prideFlag)
	}
}

func BenchmarkRuneLen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = len(prideFlagRune)
	}
}

func BenchmarkRuneCountInString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = utf8.RuneCountInString(prideFlag)
	}
}

func BenchmarkRuneCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = utf8.RuneCount(prideFlagBytes)
	}
}

func BenchmarkGraphemeClusterCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = uniseg.GraphemeClusterCount(prideFlag)
	}
}
