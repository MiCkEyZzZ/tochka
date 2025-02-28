package tochka

import (
	"errors"
	"fmt"
	"image"
	"math"
)

// Point представляет точку в двумерной системе координат с координатами X и Y.
type Point struct {
	X, Y float32
}

// NewPoint создаёт и возвращает новую точку с заданными координатами x и y.
func NewPoint(x, y float32) Point {
	return Point{X: x, Y: y}
}

// Add возвращает новую точку, полученную сложением текущей точки с другой.
func (p Point) Add(point Point) Point {
	return Point{X: p.X + point.X, Y: p.Y + point.Y}
}

// Sub возвращает новую точку, координаты которой равны разности текущей точки и другой.
func (p Point) Sub(point Point) Point {
	return Point{X: p.X - point.X, Y: p.Y - point.Y}
}

// Mul возвращает новую точку с координатами, умноженными на заданный множитель s.
func (p Point) Mul(s float32) Point {
	return Point{X: p.X * s, Y: p.Y * s}
}

// Div возвращает новую точку с координатами, разделенными на заданный коэффициент s.
func (p Point) Div(s float32) (Point, error) {
	if s == 0 {
		return Point{}, errors.New("деление на ноль")
	}
	return Point{X: p.X / s, Y: p.Y / s}, nil
}

// Distance возвращает расстояние между текущей точкой и заданной точкой.
func (p Point) Distance(point Point) float32 {
	dx := p.X - point.X
	dy := p.Y - point.Y
	return float32(math.Sqrt(float64(dx*dx + dy*dy)))
}

// Dot возвращает скалярное произведение двух векторов.
func (p Point) Dot(point Point) float32 {
	return p.X*point.X + p.Y*point.Y
}

// Cross возвращает псевдовекторное произведение (детерминант) двух векторов в 2М.
func (p Point) Cross(point Point) float32 {
	return p.X*point.Y - p.Y*point.X
}

// Magnitude возвращает длину вектора.
func (p Point) Magnitude() float32 {
	return float32(math.Sqrt(float64(p.X*p.X + p.Y*p.Y)))
}

// Round округляет координаты точки до ближайших целых значений и возвращает объект image.Point.
func (p Point) Round() image.Point {
	return image.Point{
		X: int(math.Round(float64(p.X))),
		Y: int(math.Round(float64(p.Y))),
	}
}

// String возвращает строковое представление точки в формате "(X, Y)".
func (p Point) String() string {
	return fmt.Sprintf("(%.6f, %.6f)", p.X, p.Y)
}
