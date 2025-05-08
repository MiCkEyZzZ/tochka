package tochka

import (
	"math"
	"testing"
)

// BenchmarkSplit measures the performance of the Split() method.
func BenchmarkSplit(b *testing.B) {
	a := NewAffine2D(2, 1, 3, 1, 2, 4)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a.Split()
	}
}

// BenchmarkShear measures the performance of the Shear() method.
func BenchmarkShear(b *testing.B) {
	a := NewAffine2D(1, 0, 0, 0, 1, 0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a.Shear(Point{X: 0, Y: 0}, float32(math.Pi/6), float32(math.Pi/6))
	}
}

// BenchmarkElems measures the performance of the Elems() method.
func BenchmarkElems(b *testing.B) {
	a := NewAffine2D(2, 1, 3, 1, 2, 4)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a.Elems()
	}
}

// BenchmarkAffineString measures the performance of the String() method.
func BenchmarkAffineString(b *testing.B) {
	a := NewAffine2D(1, 0, 2, 0, 1, 3)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a.String()
	}
}
