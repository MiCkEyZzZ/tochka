## tochka — a package for 2D points and affine transformations

[![godoc](https://godoc.org/github.com/MiCkEyZzZ/tochka?status.svg)](https://pkg.go.dev/github.com/MiCkEyZzZ/tochka?tab=doc)

`tochka` is a Go package that provides functionality for working with
2D points and affine transformation matrices. This package is ideal for
graphical applications, geometry processing, and other tasks related to
2D space.

## Features

- **Working with Points:**
  - Operations for addition, subtraction, multiplication, and division.
  - Rounding coordinates to integers.
  - Calculating the dot product (`Dot`).
  - Calculating the pseudovector product (`Cross`).
  - Finding the length of a vector (`Magnitude`).
  - Convenient string representation of points in the format `(X, Y)`.

- **Affine Transformations:**
  - Operations for translation, scaling, rotation, and shear.
  - Combining transformations using matrix multiplication.
  - Inverting the transformation matrix.
  - Applying transformations to 2D points.

- A simple and intuitive API for developers.

## Installation

To install the package, use the following command:

```zsh
go get github.com/MiCkEyZzZ/tochka
```

## Example Usage

### Working with Points

```go
package main

import (
	"fmt"

	"github.com/MiCkEyZzZ/tochka"
)

func main() {
	p1 := tochka.NewPoint(2.5, 3.7)
	p2 := tochka.NewPoint(1.2, -0.5)

	sum := p1.Add(p2)
	fmt.Println("Sum of points:", sum)  // Sum of points: (3.700000, 3.200000)

	rounded := sum.Round()
	fmt.Println("Rounded coordinates:", rounded)  // Rounded coordinates: (4, 3)
}
```

### Affine Transformations

```go
package main

import (
	"fmt"

	"github.com/MiCkEyZzZ/tochka"
)

func main() {
	transform := tochka.NewAffine2D(1, 0, 10, 0, 1, 20) // Translation by (10, 20)

	point := tochka.NewPoint(5, 5)
	transformed := transform.Transform(point)

	fmt.Println("Original point:", point)  // Original point: (5.000000, 5.000000)
	fmt.Println("After transformation:", transformed)  // After transformation: (15.000000, 25.000000)
}
```

## API

The package provides the following key methods for working with points and affine transformations:

### For the `Point` type:

- `NewPoint(x, y float32) Point`: Creates a point with the specified coordinates.
- `Add(other Point) Point`: Adds the current point to another.
- `Sub(other Point) Point`: Subtracts another point from the current one.
- `Mul(scale Point) Point`: Multiplies the point coordinates by the given scale.
- `Div(scale Point) (Point, error)`: Divides the point coordinates, returning an error when dividing by 0.
- `Round() Point`: Rounds the coordinates to the nearest integers.
- `Dot(other Point) float32`: Computes the dot product of two vectors.
- `Cross(other Point) float32`: Computes the pseudovector product (determinant) of two vectors.
- `Magnitude() float32`: Returns the length of the vector.

### For the `Affine2D` type:

- `NewAffine2D(sx, hx, ox, hy, sy, oy float32) Affine2D`: Creates an affine transformation.
- `Offset(offset Point) Affine2D`: Shifts the matrix by the specified vector.
- `Scale(origin, factor Point) Affine2D`: Scales relative to the specified point.
- `Rotate(origin Point, radians float32) Affine2D`: Rotates a point by the given angle in radians.
- `Shear(origin Point, radiansX, radiansY float32) Affine2D`: Applies a shear transformation.
- `Mul(other Affine2D) Affine2D`: Multiplies matrices to combine transformations.
- `Invert() Affine2D`: Computes the inverse of the transformation.
- `Transform(p Point) Point`:  Applies the transformation to a point.
- `Elems() (sx, hx, ox, hy, sy, oy float32)`: Returns the matrix elements.
- `Split() (Affine2D, Point)`: Splits the transformation into a matrix and a translation vector.

The full list of methods and their descriptions can be found in the [documentation](https://pkg.go.dev/github.com/MiCkEyZzZ/tochka).

## License

This project is licensed under the MIT License. The full license text is available in the [License](./LICENSE).
