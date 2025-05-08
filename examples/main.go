package main

import (
	"fmt"

	"github.com/MiCkEyZzZ/tochka"
)

func main() {
	// Creating new points.
	p1 := tochka.NewPoint(2.5, 3.5)
	p2 := tochka.NewPoint(1.0, 1.0)

	fmt.Println("Point 1:", p1)
	fmt.Println("Point 2:", p2)

	// Operations with points
	sum := p1.Add(p2)
	fmt.Println("Sum of points:", sum)

	diff := p1.Sub(p2)
	fmt.Println("Difference of points:", diff)

	scaled := p1.Mul(2.0)
	fmt.Println("Scaling of point 1 by 2:", scaled)

	divided, err := p1.Div(2.0)
	if err != nil {
		fmt.Println("Error during division:", err)
	} else {
		fmt.Println("Point 1 divided by 2:", divided)
	}

	// Distance between points
	dist := p1.Distance(p2)
	fmt.Printf("Distance between point 1 and point 2: %.2f\n", dist)

	// Rounding coordinates
	rounded := p1.Round()
	fmt.Println("Point 1 with rounded coordinates:", rounded)

	// String representation of the point
	fmt.Println("String representation of point 1:", p1)

	// Dot product
	dotProduct := p1.Dot(p2)
	fmt.Println("Dot product:", dotProduct)

	// Cross product
	crossProduct := p1.Cross(p2)
	fmt.Println("Cross product:", crossProduct)

	// Magnitude of vector p1
	magnitude1 := p1.Magnitude()
	fmt.Println("Magnitude of vector p1:", magnitude1)

	// Magnitude of vector p2
	magnitude2 := p2.Magnitude()
	fmt.Println("Magnitude of vector p2:", magnitude2)
}
