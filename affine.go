package tochka

import (
	"math"
	"strconv"
	"strings"
)

// Affine2D represents an affine transformation in a 2D coordinate system.
// It contains the elements of the transformation matrix that allow performing operations
// such as translation, scaling, rotation, and shear.
type Affine2D struct {
	a, b, c float32
	d, e, f float32
}

// NewAffine2D creates a new affine transformation.
// sx, hx, ox are elements of the first row of the matrix (scaling along X, shear along X, translation along X).
// hy, sy, oy are elements of the second row of the matrix (shear along Y, scaling along Y, translation along Y).
func NewAffine2D(sx, hx, ox, hy, sy, oy float32) Affine2D {
	return Affine2D{
		a: sx - 1, b: hx, c: ox,
		d: hy, e: sy - 1, f: oy,
	}
}

// Offset performs a translation on the transformation matrix by a given vector.
// offset defines the amount of translation along X and Y.
func (a Affine2D) Offset(offset Point) Affine2D {
	return Affine2D{
		a.a, a.b, a.c + offset.X, // translation along X
		a.d, a.e, a.f + offset.Y, // translation along Y
	}
}

// OffsetInPlace translates the matrix in place.
// offset defines the amount of translation along X and Y.
func (a *Affine2D) OffsetInPlace(offset Point) {
	a.c += offset.X
	a.f += offset.Y
}

// Scale performs scaling on the matrix relative to a given point.
// origin defines the point around which scaling occurs.
// factor defines the scaling factors along X and Y.
// Returns a new transformation matrix that takes scaling into account.
func (a Affine2D) Scale(origin, factor Point) Affine2D {
	if origin == (Point{}) {
		return a.scale(factor)
	}
	a = a.Offset(origin.Mul(-1))
	a = a.scale(factor)
	return a.Offset(origin)
}

// ScaleInPlace scales the matrix in place.
// origin defines the point relative to which scaling occurs.
// factor defines the scaling factors along X and Y.
func (a *Affine2D) ScaleInPlace(origin, factor Point) {
	*a = a.Scale(origin, factor)
}

// Rotate performs a rotation on the matrix around a given point by a specified angle.
// origin defines the point around which the rotation occurs.
// radians defines the angle of rotation in radians.
// Returns a new transformation matrix that takes rotation into account.
func (a Affine2D) Rotate(origin Point, radians float32) Affine2D {
	if origin == (Point{}) {
		return a.rotate(radians)
	}
	a = a.Offset(origin.Mul(-1)) // shift to origin
	a = a.rotate(radians)        // rotate
	return a.Offset(origin)      // return to original position
}

// RotateInPlace rotates the matrix in place.
// origin defines the point relative to which rotation occurs.
// radians defines the angle of rotation in radians.
func (a *Affine2D) RotateInPlace(origin Point, radians float32) {
	*a = a.Rotate(origin, radians)
}

// Shear performs a shear transformation on the matrix under given angles relative to a specified point.
// origin defines the point relative to which shearing occurs.
// radiansX and radiansY define the shear angles along the X and Y axes, respectively.
// Returns a new transformation matrix that takes shear into account.
func (a Affine2D) Shear(origin Point, radiansX, radiansY float32) Affine2D {
	if origin == (Point{}) {
		return a.shear(radiansX, radiansY)
	}
	a = a.Offset(origin.Mul(-1))    // shift the coordinate system
	a = a.shear(radiansX, radiansY) // apply shear
	return a.Offset(origin)         // return to original position
}

// Mul multiplies the current matrix by another matrix.
// B is the other transformation matrix.
// Returns the result of the multiplication.
func (A Affine2D) Mul(B Affine2D) (r Affine2D) {
	r.a = (A.a+1)*(B.a+1) + A.b*B.d - 1
	r.b = (A.a+1)*B.b + A.b*(B.e+1)
	r.c = (A.a+1)*B.c + A.b*B.f + A.c
	r.d = A.d*(B.a+1) + (A.e+1)*B.d
	r.e = A.d*B.b + (A.e+1)*(B.e+1) - 1
	r.f = A.d*B.c + (A.e+1)*B.f + A.f
	return r
}

// Invert computes the inverse transformation for the current matrix.
// If the matrix is close to singular, the results may be inaccurate.
func (a Affine2D) Invert() Affine2D {
	if math.Abs(float64(a.a)) < 1e-6 && math.Abs(float64(a.b)) < 1e-6 &&
		math.Abs(float64(a.d)) < 1e-6 && math.Abs(float64(a.e)) < 1e-6 {
		return Affine2D{a: 0, b: 0, c: -a.c, d: 0, e: 0, f: -a.f}
	}
	a.a += 1
	a.e += 1
	det := a.a*a.e - a.b*a.d
	if math.Abs(float64(det)) < 1e-6 {
		return Affine2D{} // matrix is singular
	}
	a.a, a.e = a.e/det, a.a/det
	a.b, a.d = -a.b/det, -a.d/det
	temp := a.c
	a.c = -a.a*a.c - a.b*a.f
	a.f = -a.d*temp - a.e*a.f
	a.a -= 1
	a.e -= 1
	return a
}

// Transform applies the current transformation to a given point.
// p is the point to which the transformation is applied.
func (a Affine2D) Transform(p Point) Point {
	return Point{
		X: p.X*(a.a+1) + p.Y*a.b + a.c,
		Y: p.Y*a.d + p.Y*(a.e+1) + a.f,
	}
}

// Elems returns the elements of the transformation matrix.
func (a Affine2D) Elems() (sx, hx, ox, hy, sy, oy float32) {
	return a.a + 1, a.b, a.c, a.d, a.e + 1, a.f
}

// Split splits the transformation into a matrix without translation and a translation vector.
// Returns a new matrix and a translation vector.
func (a Affine2D) Split() (src Affine2D, offset Point) {
	return Affine2D{
		a: a.a, b: a.b, c: 0,
		d: a.d, e: a.e, f: 0,
	}, Point{X: a.c, Y: a.f}
}

// scale performs internal scaling of the current matrix.
// factor defines the scaling factors.
func (a Affine2D) scale(factor Point) Affine2D {
	return Affine2D{
		(a.a+1)*factor.X - 1, a.b * factor.X, a.c * factor.X, // scale along X
		a.d * factor.Y, (a.e+1)*factor.Y - 1, a.f * factor.Y, // scale along Y
	}
}

// rotate performs internal rotation of the matrix by the specified angle.
// radians defines the angle of rotation in radians.
func (a Affine2D) rotate(radians float32) Affine2D {
	// находим синус и косинус
	sin, cos := math.Sincos(float64(radians))
	s, c := float32(sin), float32(cos)
	return Affine2D{
		(a.a+1)*c - a.d*s - 1, a.b*c - (a.e+1)*s, a.c*c - a.f*s, // apply rotation
		(a.a+1)*s + a.d*c, a.b*s + (a.e+1)*c - 1, a.c*s + a.f*c, // update elements for translation along Y
	}
}

// shear performs internal shearing of the matrix under given angles.
// radiansX and radiansY define the shear angles along the X and Y axes.
func (a Affine2D) shear(radiansX, radiansY float32) Affine2D {
	// вычисляем тангенсы углов для сдвига.
	tx := float32(math.Tan(float64(radiansX)))
	ty := float32(math.Tan(float64(radiansY)))

	return Affine2D{
		(a.a + 1) + a.d*tx - 1, a.b + (a.e+1)*tx, a.c + a.f*tx, // update elements for shifting along X
		(a.a+1)*ty + a.d, a.b*ty + (a.e + 1) - 1, a.c*ty + a.f, // update elements for shifting along Y
	}
}

// String returns the string representation of the transformation matrix.
// Format: "[[sx hx ox] [hy sy oy]]".
func (a Affine2D) String() string {
	sx, hx, ox, hy, sy, oy := a.Elems()
	var b strings.Builder
	b.WriteString("[[")
	b.WriteString(strconv.FormatFloat(float64(sx), 'g', 6, 32))
	b.WriteString(" ")
	b.WriteString(strconv.FormatFloat(float64(hx), 'g', 6, 32))
	b.WriteString(" ")
	b.WriteString(strconv.FormatFloat(float64(ox), 'g', 6, 32))
	b.WriteString("] [")
	b.WriteString(strconv.FormatFloat(float64(hy), 'g', 6, 32))
	b.WriteString(" ")
	b.WriteString(strconv.FormatFloat(float64(sy), 'g', 6, 32))
	b.WriteString(" ")
	b.WriteString(strconv.FormatFloat(float64(oy), 'g', 6, 32))
	b.WriteString("]]")
	return b.String()
}
