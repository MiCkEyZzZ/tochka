// The tochka package provides tools for working with two-dimensional
// points and affine transformations in a two-dimensional coordinate
// system. The package is suitable for performing operations such as
// translation, scaling, rotation, and shearing, making it
// useful for manipulating 2D geometry in graphical applications.
//
// # Point Type
//
// The Point type represents a point in a two-dimensional coordinate
// system with X and Y coordinates (of type float32). This type includes
// basic arithmetic operations and helper methods for working with
// points.
//
// Key methods:
//   - NewPoint(x, y float32) Point: Creates a point with specified coordinates.
//   - String() string: Returns a string representation of the point in the format "(X, Y)".
//   - Add(other Point) Point: Adds the current point to another point.
//   - Sub(other Point) Point: Subtracts another point from the current point.
//   - Mul(scale float32) Point: Multiplies the point's coordinates by a scale factor.
//   - Div(scale float32) (Point, error): Divides the point's coordinates by a scale factor, returning an error when dividing by zero.
//   - Distance(other Point) float32: Returns the distance between the current point and another point.
//   - Round() image.Point: Rounds the coordinates to the nearest integers and returns an image.Point object.
//   - Dot(other Point) float32: Computes the dot product of two vectors.
//   - Cross(other Point) float32: Computes the pseudo-vector (determinant) product of two vectors in 2D.
//   - Magnitude() float32: Returns the length of the vector.
//
// # Affine2D Type
//
// The Affine2D type represents an affine transformation described by a 3x3
// matrix, which supports transformation operations in two-dimensional space,
// including translation, scaling, rotation, and shearing.
//
// Key methods:
//   - NewAffine2D(sx, hx, ox, hy, sy, oy float32) Affine2D: Creates a transformation from the matrix elements.
//   - Offset(offset Point) Affine2D: Performs a translation of the matrix by the specified vector.
//   - Scale(origin, factor Point) Affine2D: Scales relative to a given point.
//   - Rotate(origin Point, radians float32) Affine2D: Rotates around a given point by an angle in radians.
//   - Shear(origin Point, radiansX, radiansY float32) Affine2D: Applies a shearing transformation to the matrix with specified angles.
//   - Mul(other Affine2D) Affine2D: Multiplies the current transformation by another transformation.
//   - Invert() Affine2D: Computes the inverse transformation, if possible.
//   - Transform(p Point) Point: Applies the transformation to a point and returns a new point.
//   - Elems() (sx, hx, ox, hy, sy, oy float32): Returns the elements of the transformation matrix.
//   - Split() (Affine2D, Point): Splits the transformation into a matrix without translation and a translation vector.
//   - String() string: Returns a string representation of the transformation matrix in the format "[[sx hx ox] [hy sy oy]]".
//
// The package is designed for integration into graphical applications and
// for processing 2D geometry. It is useful for both educational and production
// projects where working with points and affine transformations in two-dimensional
// space is required.
package tochka
