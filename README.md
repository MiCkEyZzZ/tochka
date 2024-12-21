## tochka — пакет для двумерных точек и аффинных преобразований.

`tochka` — это пакет на языке Go, которая предоставляет функционал для работы
с двухмерными точками и матрицами аффинных преобразований. Этот пакет идеально
подходит для графических приложений, обработки геометрии и других задач,
связанных с 2М-пространством.

## Особенности

- **Работа с точками:**
  - Операции сложения, вычитания, умножения и деления.
  - Округление координат до целых чисел.
  - Удобное строковое представление.

- **Аффинные преобразования:**
  - Сдвиг, масштабирование, вращение и сдвиг по углу (shear).
  - Умножение матриц для комбинирования преобразований.
  - Инверсия матрицы преобразования.
  - Применение преобразований к двумерным точкам.

- Простое и интуитивное API для разработчиков.

## Установка

Для установки пакета используйте команду:

```zsh
go get github.com/MiCkEyZzZ/tochka
```

## Пример использования

### Работа с точками

```go
package main

import (
	"fmt"

	"github.com/MiCkEyZzZ/tochka/geom"
)

func main() {
	p1 := geom.NewPoint(2.5, 3.7)
	p2 := geom.NewPoint(1.2, -0.5)

	sum := p1.Add(p2)
	fmt.Println("Сумма точек:", sum)

	rounded := sum.Round()
	fmt.Println("Округленные координаты: ", rounded)
}
```

### Вывод программы:

```zsh
Сумма точек: (3.700000, 3.200000)
Округленные координаты: (4, 3)
```

### Аффинные преобразования

```go
package main

import (
	"fmt"

	"github.com/MiCkEyZzZ/tochka/geom"
)

func main() {
	transform := geom.NewAffine2D(1, 0, 10, 0, 1, 20) // Сдвиг на (10, 20)

	point := geom.NewPoint(5, 5)
	transformed := transform.Transform(point)

	fmt.Println("Исходная точка:", point)
	fmt.Println("После преобразования:", transformed)
}
```

### Вывод программы:

```zsh
Исходная точка: (5.000000, 5.000000)
После преобразования: (15.000000, 25.000000)
```

## API

Пакет предоставляет следующие ключевые методы для преобразования:

Полный список методов и их описание можно найти в [документации]().

## Лицензия

Этот проект распространяется под лицензией MIT. Полный текст лицензии доступен в файле [ЛИЦЕНЗИЯ](./LICENSE).
