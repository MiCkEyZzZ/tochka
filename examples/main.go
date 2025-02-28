package main

import (
	"fmt"

	"github.com/MiCkEyZzZ/tochka"
)

func main() {
	// Создание новых точек.
	p1 := tochka.NewPoint(2.5, 3.5)
	p2 := tochka.NewPoint(1.0, 1.0)

	fmt.Println("Точка 1:", p1)
	fmt.Println("Точка 2:", p2)

	// Операция с точками
	sum := p1.Add(p2)
	fmt.Println("Сумма точек:", sum)

	diff := p1.Sub(p2)
	fmt.Println("Разность точек:", diff)

	scaled := p1.Mul(2.0)
	fmt.Println("Масштабирование точки 1 и 2:", scaled)

	divided, err := p1.Div(2.0)
	if err != nil {
		fmt.Println("Ошибка при делении:", err)
	} else {
		fmt.Println("Точка 1, разделённая на 2:", divided)
	}

	// Расстояние между точками
	dist := p1.Distance(p2)
	fmt.Printf("Расстояние между точкой 1 и точкой 2: %.2f\n", dist)

	// Округление координат
	rounded := p1.Round()
	fmt.Println("Точка 1 с округлёнными координатами:", rounded)

	// Строковое представление точки
	fmt.Println("Строковое представление точки 1:", p1)

	// Скалярное произведение
	dotProduct := p1.Dot(p2)
	fmt.Println("Скалярное произведение:", dotProduct)

	// Псевдовекторное произведение
	crossProduct := p1.Cross(p2)
	fmt.Println("Псевдовекторное произведение:", crossProduct)

	// Длина вектора p1
	magnitude1 := p1.Magnitude()
	fmt.Println("Длина вектора p1:", magnitude1)

	// Длина вектора p2
	magnitude2 := p2.Magnitude()
	fmt.Println("Длина вектора p2:", magnitude2)
}
