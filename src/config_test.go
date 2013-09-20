package main

import (
	"testing"
)

func TestloadConfig(t *testing.T) {

}

func BenchmarkLoadConfig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		loadConfig()
	}
}
