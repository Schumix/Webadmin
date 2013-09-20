package main

import (
	"testing"
)

func TestConnectToSql(t *testing.T) {

}

func BenchmarkConnectToSql(b *testing.B) {
	for i := 0; i < b.N; i++ {
		connectToSql()
	}
}
