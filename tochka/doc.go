// Пакет tochka предоставляет инструменты для работы с двухмерными точками и
// аффинными преобразованиями в двумерной системе координат. Пакет подходит для
// выполнения операций сдвига, масштабирования, вращения и сдвига по углу,
// что делает его полезным для манипуляции 2М-геометрией в графических приложениях.
//
// # Тип Point
//
// Point представляет собой точку в двумерной системе координат с координатами
// X и Y (тип float32). Этот тип включает основные арифметические операции и
// вспомогательные методы для работы с точками.
//
// Основные методы:
//   - NewPoint(x, y float32) Point: Создает точку с заданными координатами.
//   - String() string: Возвращает строковое представление точки в формате "(X, Y)".
//   - Add(other Point) Point: Складывает текущую точку с другой.
//   - Sub(other Point) Point: Вычитает другую точку из текущей.
//   - Mul(scale Point) Point: Умножает координаты точки на коэффициенты.
//   - Div(scale Point) (Point, error): Делит координаты точки, возвращая ошибку при делении на 0.
//   - Round() Point: Округляет координаты до ближайших целых.
//
// # Тип Affine2D
//
// Affine2D представляет аффинное преобразование, описанное матрицей 3x3, которая
// поддерживает операции трансформации в двумерном пространстве.
//
// Основные методы:
//   - NewAffine2D(sx, hx, ox, hy, sy, oy float32) Affine2D: Создает преобразование из элементов матрицы.
//   - Offset(offset Point) Affine2D: Выполняет сдвиг матрицы на указанный вектор.
//   - Scale(origin, factor Point) Affine2D: Масштабирует относительно заданной точки.
//   - Rotate(origin Point, radians float32) Affine2D: Вращает вокруг заданной точки на угол в радианах.
//   - Shear(origin Point, radiansX, radiansY float32) Affine2D: Применяет сдвиг по заданным углам.
//   - Mul(other Affine2D) Affine2D: Умножает текущее преобразование на другое.
//   - Invert() Affine2D: Вычисляет обратное преобразование (если возможно).
//   - Transform(p Point) Point: Применяет преобразование к точке.
//   - Elems() (sx, hx, ox, hy, sy, oy float32): Возвращает элементы матрицы в строковом формате.
//   - Split() (Affine2D, Point): Разделяет матрицу на сдвиг и остальную часть.
//   - String() string: Возвращает строковое представление матрицы.
//
// Пакет разработан для интеграции в графические приложения и обработки
// 2М-геометрии. Его легко использовать как в учебных, так и в производственных
// проектах.
package tochka
