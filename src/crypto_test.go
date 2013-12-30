package main

import (
	"testing"
)

func TestSha1_gen(t *testing.T) {

}

func BenchmarkSha1_gen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sha1_gen("udweh782gtoFzgnxoG(OXF8nYFBOYdnwhauog/I=f67Fo8goIGo8")
	}
}

func TestMd5_gen(t *testing.T) {

}

func BenchmarkMd5_gen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		md5_gen("udweh782gtoFzgnxoG(OXF8nYFBOYdnwhauog/I=f67Fo8goIGo8")
	}
}
