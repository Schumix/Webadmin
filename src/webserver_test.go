package main

import (
	"testing"
)

func TestLoadServer(t *testing.T) {

}

func BenchmarkLoadServer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		loadServer(":12345")
	}
}
