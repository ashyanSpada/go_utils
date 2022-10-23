package geohash

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeoHash(t *testing.T) {
	lat := 30.559545
	lng := 104.059684
	b := EncodeInt(lat, lng)
	fmt.Printf("%2b", b)
	c := EncodeInt2(lat, lng)
	d := EncodeInt3(lat, lng)
	e := EncodeIntAsm(lat, lng)
	assert.Equal(t, b, c, d, e)
	// fmt.Println(b, c, d, e)
}

func BenchmarkEncodeInt(b *testing.B) {
	var lat, lng = 40.463833, -79.972422
	for i := 0; i < b.N; i++ {
		EncodeInt(lat, lng)
	}
}

func BenchmarkEncodeInt2(b *testing.B) {
	var lat, lng = 40.463833, -79.972422
	for i := 0; i < b.N; i++ {
		EncodeInt2(lat, lng)
	}
}

func BenchmarkEncodeIntAsm(b *testing.B) {
	var lat, lng = 40.463833, -79.972422
	for i := 0; i < b.N; i++ {
		EncodeIntAsm(lat, lng)
	}
}

func BenchmarkEncodeInt4(b *testing.B) {
	var lat, lng = 40.463833, -79.972422
	for i := 0; i < b.N; i++ {
		EncodeInt3(lat, lng)
	}
}
