package tochka

import (
	"math"
	"strconv"
	"strings"
)

// Affine2D представляет аффинное преобразование в двумерной системе координат.
// Содержит элементы матрицы преобразования, которые позволяют выполнять
// операции сдвига, масштабирования, вращения и сдвига по углу.
type Affine2D struct {
	a, b, c float32
	d, e, f float32
}

// NewAffine2D создает новое аффинное преобразование.
// sx, hx, ox — элементы первой строки матрицы (масштаб по X, сдвиг по X, перенос по X).
// hy, sy, oy — элементы второй строки матрицы (сдвиг по Y, масштаб по Y, перенос по Y).
// Возвращает объект Affine2D.
func NewAffine2D(sx, hx, ox, hy, sy, oy float32) Affine2D {
	return Affine2D{
		a: sx - 1, b: hx, c: ox,
		d: hy, e: sy - 1, f: oy,
	}
}

// Offset выполняет сдвиг матрицы преобразования на заданный вектор.
// offset — точка, определяющая величину сдвига по X и Y.
// Возвращает новую матрицу преобразования с учетом сдвига.
func (a Affine2D) Offset(offset Point) Affine2D {
	return Affine2D{
		a.a, a.b, a.c + offset.X, // сдвиг по оси Х
		a.d, a.e, a.f + offset.Y, // сдвиг по оси Y
	}
}

// Scale выполняет масштабирование преобразования вокруг заданной точки.
// origin — точка, вокруг которой выполняется масштабирование.
// factor — коэффициенты масштабирования по X и Y.
// Возвращает новую матрицу преобразования с учетом масштабирования.
func (a Affine2D) Scale(origin, factor Point) Affine2D {
	if origin == (Point{}) {
		return a.scale(factor)
	}
	a = a.Offset(origin.Mul(-1))
	a = a.scale(factor)
	return a.Offset(origin)
}

// Rotate выполняет вращение вокруг заданной точки на указанный угол.
// origin — точка, вокруг которой выполняется вращение.
// radians — угол вращения в радианах.
// Возвращает новую матрицу преобразования с учетом вращения.
func (a Affine2D) Rotate(origin Point, radians float32) Affine2D {
	if origin == (Point{}) {
		return a.rotate(radians)
	}
	a = a.Offset(origin.Mul(-1)) // сдвигаем в начало
	a = a.rotate(radians)        // вращаем
	return a.Offset(origin)      // возвращаем в исходное положение
}

// Shear выполняет сдвиг под заданными углами относительно указанной точки.
// origin — точка, вокруг которой выполняется сдвиг.
// radiansX, radiansY — углы сдвига по осям X и Y.
// Возвращает новую матрицу преобразования с учетом сдвига.
func (a Affine2D) Shear(origin Point, radiansX, radiansY float32) Affine2D {
	if origin == (Point{}) {
		return a.shear(radiansX, radiansY)
	}
	a = a.Offset(origin.Mul(-1))    // сдвигаем систему координат
	a = a.shear(radiansX, radiansY) // применяем сдвиг
	return a.Offset(origin)         // возвращаем в исходное положение
}

// Mul выполняет умножение текущей матрицы преобразования на другую.
// B — вторая матрица преобразования.
// Возвращает результат умножения (новую матрицу).
func (A Affine2D) Mul(B Affine2D) (r Affine2D) {
	r.a = (A.a+1)*(B.a+1) + A.b*B.d - 1
	r.b = (A.a+1)*B.b + A.b*(B.e+1)
	r.c = (A.a+1)*B.c + A.b*B.f + A.c
	r.d = A.d*(B.a+1) + (A.e+1)*B.d
	r.e = A.d*B.b + (A.e+1)*(B.e+1) - 1
	r.f = A.d*B.c + (A.e+1)*B.f + A.f
	return r
}

// Invert вычисляет обратное преобразование для текущей матрицы.
// Если матрица близка к сингулярной, результаты могут быть неточными.
// Возвращает новую матрицу преобразования.
func (a Affine2D) Invert() Affine2D {
	if math.Abs(float64(a.a)) < 1e-6 && math.Abs(float64(a.b)) < 1e-6 &&
		math.Abs(float64(a.d)) < 1e-6 && math.Abs(float64(a.e)) < 1e-6 {
		return Affine2D{a: 0, b: 0, c: -a.c, d: 0, e: 0, f: -a.f}
	}
	a.a += 1
	a.e += 1
	det := a.a*a.e - a.b*a.d
	if math.Abs(float64(det)) < 1e-6 {
		return Affine2D{} // матрица сингулярна
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

// Transform применяет текущее преобразование к заданной точке.
// p — точка, к которой применяется преобразование.
// Возвращает новую точку с преобразованными координатами.
func (a Affine2D) Transform(p Point) Point {
	return Point{
		X: p.X*(a.a+1) + p.Y*a.b + a.c,
		Y: p.Y*a.d + p.Y*(a.e+1) + a.f,
	}
}

// Elems возвращает элементы матрицы преобразования в порядке строк.
// sx, hx, ox — элементы первой строки.
// hy, sy, oy — элементы второй строки.
func (a Affine2D) Elems() (sx, hx, ox, hy, sy, oy float32) {
	return a.a + 1, a.b, a.c, a.d, a.e + 1, a.f
}

// Split разделяет преобразование на две части: чистый сдвиг и остальную часть.
// Возвращает матрицу без сдвига и вектор сдвига.
func (a Affine2D) Split() (src Affine2D, offset Point) {
	return Affine2D{
		a: a.a, b: a.b, c: 0,
		d: a.d, e: a.e, f: 0,
	}, Point{X: a.c, Y: a.f}
}

// scale выполняет внутреннее масштабирование текущей матрицы преобразования.
// factor — точка, задающая коэффициенты масштабирования по X и Y.
// Возвращает новую матрицу преобразования с учетом масштабирования.
func (a Affine2D) scale(factor Point) Affine2D {
	return Affine2D{
		(a.a+1)*factor.X - 1, a.b * factor.X, a.c * factor.X, // масштабируем по X
		a.d * factor.Y, (a.e+1)*factor.Y - 1, a.f * factor.Y, // масштабируем по Y
	}
}

// rotate выполняет внутреннее вращение текущей матрицы преобразования на заданный угол.
// radians — угол в радианах, на который нужно выполнить вращение.
// Возвращает новую матрицу преобразования с учетом вращения.
func (a Affine2D) rotate(radians float32) Affine2D {
	// находим синус и косинус
	sin, cos := math.Sincos(float64(radians))
	s, c := float32(sin), float32(cos)
	return Affine2D{
		(a.a+1)*c - a.d*s - 1, a.b*c - (a.e+1)*s, a.c*c - a.f*s, // применяем вращение
		(a.a+1)*s + a.d*c, a.b*s + (a.e+1)*c - 1, a.c*s + a.f*c, // обновляем элементы для сдвига по Y
	}
}

// shear выполняет внутренний сдвиг текущей матрицы преобразования под заданными углами.
// radiansX — угол наклона относительно оси X в радианах.
// radiansY — угол наклона относительно оси Y в радианах.
// Возвращает новую матрицу преобразования с учетом сдвига.
func (a Affine2D) shear(radiansX, radiansY float32) Affine2D {
	// вычисляем тангенсы углов для сдвига.
	tx := float32(math.Tan(float64(radiansX)))
	ty := float32(math.Tan(float64(radiansY)))

	return Affine2D{
		(a.a + 1) + a.d*tx - 1, a.b + (a.e+1)*tx, a.c + a.f*tx, // обновляем элементы для сдвига по X
		(a.a+1)*ty + a.d, a.b*ty + (a.e + 1) - 1, a.f*ty + a.f, // обновляем элементы для сдвига по Y
	}
}

// String возвращает строковое представление матрицы преобразования.
// Формат: "[[sx hx ox] [hy sy oy]]".
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
