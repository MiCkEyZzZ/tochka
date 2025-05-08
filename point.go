package tochka

import (
	"errors"
	"fmt"
	"image"
	"math"
)

// Point represents a point in a two-dimensional coordinate system with X and Y coordinates.
type Point struct {
	X, Y float32
}

// NewPoint creates and returns a new point with the specified x and y coordinates.
func NewPoint(x, y float32) Point {
	return Point{X: x, Y: y}
}

// Add returns a new point obtained by adding the current point to another.
func (p Point) Add(point Point) Point {
	return Point{X: p.X + point.X, Y: p.Y + point.Y}
}

// Sub returns a new point whose coordinates are the difference between the current point and another.
func (p Point) Sub(point Point) Point {
	return Point{X: p.X - point.X, Y: p.Y - point.Y}
}

// Mul returns a new point with coordinates multiplied by a given factor s.
func (p Point) Mul(s float32) Point {
	return Point{X: p.X * s, Y: p.Y * s}
}

// Div returns a new point with coordinates divided by a given factor s.
func (p Point) Div(s float32) (Point, error) {
	if s == 0 {
		return Point{}, errors.New("division by zero")
	}
	return Point{X: p.X / s, Y: p.Y / s}, nil
}

// Distance returns the distance between the current point and a given point.
func (p Point) Distance(point Point) float32 {
	dx := p.X - point.X
	dy := p.Y - point.Y
	return float32(math.Sqrt(float64(dx*dx + dy*dy)))
}

// Dot returns the dot product of two vectors.
func (p Point) Dot(point Point) float32 {
	return p.X*point.X + p.Y*point.Y
}

// Cross returns the pseudovector (determinant) product of two vectors in 2D.
func (p Point) Cross(point Point) float32 {
	return p.X*point.Y - p.Y*point.X
}

// Magnitude returns the length of the vector.
func (p Point) Magnitude() float32 {
	return float32(math.Sqrt(float64(p.X*p.X + p.Y*p.Y)))
}

// Round rounds the point's coordinates to the nearest integers and returns an image.Point object.
func (p Point) Round() image.Point {
	return image.Point{
		X: int(math.Round(float64(p.X))),
		Y: int(math.Round(float64(p.Y))),
	}
}

// String returns a string representation of the point in the format "(X, Y)".
func (p Point) String() string {
	return fmt.Sprintf("(%.6f, %.6f)", p.X, p.Y)
}
