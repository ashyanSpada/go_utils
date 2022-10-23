package geohash

import (
	"math"
)

func EncodeInt(lat, lng float64) uint64 {
	latMin, latMax := -90.0, 90.0
	lngMin, lngMax := -180.0, 180.0
	encodeLat := encode(latMin, latMax, lat)
	encodeLng := encode(lngMin, lngMax, lng)
	return Interleave(encodeLat, encodeLng)
}

func encode(start, end, target float64) uint32 {
	var ans uint32
	for i := 0; i < 32; i++ {
		mid := (start + end) / (2.0)
		if target <= mid {
			end = mid
			ans <<= 1
		} else {
			start = mid
			ans = (ans << 1) + 1
		}
	}
	return ans
}

func EncodeInt2(lat, lng float64) uint64 {
	a, b := Quantize(lat, lng)
	return Interleave(a, b)
}

func Quantize(lat, lng float64) (lat32 uint32, lng32 uint32) {
	lat32 = uint32(math.Ldexp((lat+90.0)/180.0, 32))
	lng32 = uint32(math.Ldexp((lng+180.0)/360.0, 32))
	return
}

func Spread(x uint32) uint64 {
	X := uint64(x)
	X = (X | (X << 16)) & 0x0000ffff0000ffff
	X = (X | (X << 8)) & 0x00ff00ff00ff00ff
	X = (X | (X << 4)) & 0x0f0f0f0f0f0f0f0f
	X = (X | (X << 2)) & 0x3333333333333333
	X = (X | (X << 1)) & 0x5555555555555555
	return X
}

func Interleave(lat, lng uint32) uint64 {
	return Spread(lat) | (Spread(lng) << 1)
}

func EncodeIntAsm(lat, lng float64) uint64

func encodeInt(lat, lng float64) uint64 {
	return Interleave(Quantize(lat, lng))
}

func QuantizeLatBits(lat float64) uint64 {
	return math.Float64bits(lat/180.0 + 1.5)
}

func QuantizeLngBits(lng float64) uint64 {
	return math.Float64bits(lng/360.0 + 1.5)
}

func EncodeInt3(lat, lng float64) uint64 {
	a, b := QuantizeLatBits(lat), QuantizeLngBits(lng)
	c, d := uint32(a>>20), uint32(b>>20)
	return Interleave(c, d)
}
