package tochka

import (
	"image"
	"testing"
)

// TestNewPoint проверяет создание новой точки с заданными координатами.
func TestNewPoint(t *testing.T) {
	p := NewPoint(1.5, 2.5)
	if p.X != 1.5 || p.Y != 2.5 {
		t.Errorf("NewPoint(1.5, 2.5) = %v; want (1.5, 2.5)", p)
	}
}

// TestAdd проверяет сложение двух точек.
func TestAdd(t *testing.T) {
	p1 := NewPoint(1, 2)
	p2 := NewPoint(3, 4)
	result := p1.Add(p2)
	expected := NewPoint(4, 6)

	if result != expected {
		t.Errorf("Add() = %v; want %v", result, expected)
	}
}

// TestSub проверяет вычитание одной точки из другой.
func TestSub(t *testing.T) {
	p1 := NewPoint(5, 7)
	p2 := NewPoint(2, 3)
	result := p1.Sub(p2)
	expected := NewPoint(3, 4)

	if result != expected {
		t.Errorf("Sub() = %v; want %v", result, expected)
	}
}

// TestMul проверяет умножение точки на скаляр.
func TestMul(t *testing.T) {
	p1 := NewPoint(2, 3)
	result := p1.Mul(2)
	expected := NewPoint(4, 6)

	if result != expected {
		t.Errorf("Mul() = %v; want %v", result, expected)
	}
}

// TestDiv проверяет деление точки на скаляр, включая проверку на деление на ноль.
func TestDiv(t *testing.T) {
	p := NewPoint(4, 6)
	result, err := p.Div(2)
	if err != nil {
		t.Fatalf("Div() returned an error: %v", err)
	}
	expected := NewPoint(2, 3)
	if result != expected {
		t.Fatalf("Div() = %v; want %v", result, expected)
	}

	_, err = p.Div(0)
	if err == nil {
		t.Fatalf("Div() did not return an error for division by zero")
	}
}

// TestRound проверяет округление точки, включая случай с отрицательными значениями.
func TestRound(t *testing.T) {
	p := NewPoint(1.5, 2.5)
	result := p.Round()
	expected := image.Point{X: 2, Y: 3}
	if result != expected {
		t.Errorf("Round() = %v; want %v", result, expected)
	}

	p2 := NewPoint(-1.5, -2.5)
	result2 := p2.Round()
	expected2 := image.Point{X: -2, Y: -3}
	if result2 != expected2 {
		t.Errorf("Round() = %v; want %v", result2, expected2)
	}
}

// TestString проверяем, что строковое представление соответствует ожидаемому формату, включая случай, когда координаты равны нулю.
func TestString(t *testing.T) {
	p := NewPoint(1.234567, 2.345678)
	result := p.String()
	expected := "(1.234567, 2.345678)"
	if result != expected {
		t.Errorf("String() = %v; want %v", result, expected)
	}

	p2 := NewPoint(0, 0)
	result2 := p2.String()
	expected2 := "(0.000000, 0.000000)"
	if result2 != expected2 {
		t.Errorf("String() = %v; want %v", result2, expected2)
	}
}
