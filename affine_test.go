package tochka

import (
	"math"
	"testing"
)

// TestNewAffine2D tests that the constructor creates the correct transformation.
func TestNewAffine2D(t *testing.T) {
	a := NewAffine2D(1, 0, 0, 0, 1, 0)
	expected := Affine2D{0, 0, 0, 0, 0, 0}
	if a != expected {
		t.Errorf("Expected %v, got %v", expected, a)
	}
}

// TestAffine2D_Offset tests the correctness of offsetting along the X and Y axes.
func TestAffine2D_Offset(t *testing.T) {
	a := NewAffine2D(1, 0, 0, 0, 1, 0)
	offset := Point{X: 2, Y: 3}
	result := a.Offset(offset)
	if result.c != 2 || result.f != 3 {
		t.Errorf("Offset failed. Expected offset c: 2, f: 3, got c: %v, f: %v", result.c, result.f)
	}
}

// TestAffine2D_Scale tests that scaling correctly modifies the coefficients.
func TestAffine2D_Scale(t *testing.T) {
	a := NewAffine2D(1, 0, 0, 0, 1, 0)
	origin := Point{X: 0, Y: 0}
	factor := Point{X: 2, Y: 2}
	result := a.Scale(origin, factor)
	if result.a+1 != 2 || result.e+1 != 2 {
		t.Errorf("Scale failed. Expected scale a + 1: 2, e + 1: 2, got: a + 1: %v, e + 1: %v", result.a+1, result.e+1)
	}
}

// TestAffine2D_Rotate tests a 45-degree rotation.
func TestAffine2D_Rotate(t *testing.T) {
	a := NewAffine2D(1, 0, 0, 0, 1, 0)
	origin := Point{X: 0, Y: 0}
	radians := float32(math.Pi / 4) // 45 degrees
	result := a.Rotate(origin, radians)
	if math.Abs(float64(result.a+1)-math.Cos(math.Pi/4)) > 1e-6 {
		t.Errorf("Rotate failed. Expected a + 1 close to %v, got %v", math.Cos(math.Pi/4), result.a+1)
	}
}

// TestAffine2D_Invert tests matrix inversion by comparing with expected values.
func TestAffine2D_Invert(t *testing.T) {
	a := NewAffine2D(2, 0, 0, 0, 2, 0)
	inv := a.Invert()
	expected := Affine2D{-0.5, 0, 0, 0, -0.5, 0}
	if inv.a != expected.a || inv.e != expected.e {
		t.Errorf("Invert failed. Expected %v, got %v", expected, inv)
	}
}

// TestAffine2D_Transform tests applying the transformation to a point.
func TestAffine2D_Transform(t *testing.T) {
	a := NewAffine2D(1, 0, 1, 0, 1, 1)
	p := Point{X: 2, Y: 2}
	transformed := a.Transform(p)
	expected := Point{X: 3, Y: 3}
	if transformed != expected {
		t.Errorf("Transform failed. Expected %v, got %v", expected, transformed)
	}
}

// TestAffine2D_Mul tests matrix multiplication by verifying the Mul() method.
func TestAffine2D_Mul(t *testing.T) {
	a := NewAffine2D(4, 0, 3, 0, 4, 3)
	b := NewAffine2D(2, 0, 2, 0, 2, 2)
	result := a.Mul(b)
	expected := NewAffine2D(8, 0, 11, 0, 8, 11)

	if result != expected {
		t.Errorf("Mul failed. Expected %v, got %v", expected, result)
	}
}

// TestSplit tests the Split() method to separate the matrix into linear part and offset.
func TestSplit(t *testing.T) {
	a := NewAffine2D(2, 1, 3, 1, 2, 4)
	mat, off := a.Split()
	expectedMat := NewAffine2D(2, 1, 0, 1, 2, 0)
	expectedOff := Point{X: 3, Y: 4}
	if mat != expectedMat || off != expectedOff {
		t.Errorf("Split failed: got mat %+v, offset %+v", mat, off)
	}
}

// TestShear tests the Shear() method for matrix shearing.
func TestShare(t *testing.T) {
	a := NewAffine2D(1, 0, 0, 0, 1, 0)
	b := a.Shear(Point{X: 0, Y: 0}, float32(math.Pi/6), float32(math.Pi/6))
	if !almostEqual(b.b, float32(math.Tan(math.Pi/6)), 1e-6) {
		t.Errorf("Shear failed: got %+v", b)
	}
}

// TestElems tests the Elems() method for retrieving the elements of the matrix.
func TestElems(t *testing.T) {
	a := NewAffine2D(2, 1, 3, 1, 2, 4)
	sx, hx, ox, hy, sy, oy := a.Elems()
	if sx != 2 || hx != 1 || ox != 3 || hy != 1 || sy != 2 || oy != 4 {
		t.Errorf("Elems failed: got %v, %v, %v, %v, %v, %v", sx, hx, ox, hy, sy, oy)
	}
}

// TestAffineString tests the String() method for the string representation of the matrix.
func TestAffineString(t *testing.T) {
	a := NewAffine2D(1, 0, 2, 0, 1, 3)
	str := a.String()
	expected := "[[1 0 2] [0 1 3]]"
	if str != expected {
		t.Errorf("String failed: expected %s, got %s", expected, str)
	}
}

// almostEqual compares two values within a given precision.
func almostEqual(a, b float32, epsilon float32) bool {
	return math.Abs(float64(a-b)) < float64(epsilon)
}
