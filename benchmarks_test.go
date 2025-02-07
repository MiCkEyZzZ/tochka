package tochka

import (
	"math"
	"testing"
)

// BenchmarkSplit измеряет производительность метода Split().
func BenchmarkSplit(b *testing.B) {
	a := NewAffine2D(2, 1, 3, 1, 2, 4)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a.Split()
	}
}

// BenchmarkShear измеряет производительность метода Share().
func BenchmarkShear(b *testing.B) {
	a := NewAffine2D(1, 0, 0, 0, 1, 0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a.Shear(Point{X: 0, Y: 0}, float32(math.Pi/6), float32(math.Pi/6))
	}
}

// BenchmarkElems измеряет производительность метода Elems().
func BenchmarkElems(b *testing.B) {
	a := NewAffine2D(2, 1, 3, 1, 2, 4)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a.Elems()
	}
}

func BenchmarkAffineString(b *testing.B) {
	a := NewAffine2D(1, 0, 2, 0, 1, 3)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a.String()
	}
}
