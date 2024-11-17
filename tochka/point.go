package tochka

import (
	"errors"
	"fmt"
	"image"
	"math"
)

// Point представляет точку в двумерной системе координат.
// Содержит координаты X и Y в виде чисел с плавающей запятой (float32).
type Point struct {
	X, Y float32
}

// NewPoint создает новую точку с заданными координатами.
// x и y — координаты точки.
// Возвращает объект Point.
func NewPoint(x, y float32) Point {
	return Point{X: x, Y: y}
}

// Add возвращает новую точку, полученную сложением текущей точки с другой.
// point — точка, координаты которой будут добавлены к текущей.
// Возвращает новую точку с обновленными координатами.
func (p Point) Add(point Point) Point {
	return Point{X: p.X + point.X, Y: p.Y + point.Y}
}

// Sub возвращает новую точку, координаты которой равны разности текущей точки и другой.
// point — точка, координаты которой будут вычтены из текущей.
// Возвращает новую точку с результатом вычитания.
func (p Point) Sub(point Point) Point {
	return Point{X: p.X - point.X, Y: p.Y - point.Y}
}

// Mul возвращает новую точку, координаты которой умножены на заданный коэффициент.
// s — множитель для обеих координат.
// Возвращает новую точку с умноженными координатами.
func (p Point) Mul(s float32) Point {
	return Point{X: p.X * s, Y: p.Y * s}
}

// Div делит координаты текущей точки на заданный коэффициент.
// s — делитель для обеих координат.
// Возвращает новую точку с разделенными координатами или ошибку, если делитель равен нулю.
func (p Point) Div(s float32) (Point, error) {
	if s == 0 {
		return Point{}, errors.New("деление на ноль")
	}
	return Point{X: p.X / s, Y: p.Y / s}, nil
}

// Round преобразует координаты точки в целочисленные значения путем округления.
// Возвращает объект image.Point с целыми координатами.
func (p Point) Round() image.Point {
	return image.Point{
		X: int(math.Round(float64(p.X))),
		Y: int(math.Round(float64(p.Y))),
	}
}

// String возвращает строковое представление точки в формате "(X, Y)".
// Координаты округляются до шести знаков после запятой.
func (p Point) String() string {
	return fmt.Sprintf("(%.6f, %.6f)", p.X, p.Y)
}
