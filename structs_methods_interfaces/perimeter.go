package shapes

import "math"

type Shape interface {
    Area() float64
}

type Rectangle struct {
    Width float64
    Height float64
}

type Circle struct {
    Radius float64
}

func (rectangle Rectangle) Area() float64 {
    return rectangle.Width * rectangle.Height
}

func (rectangle Rectangle) Perimeter() float64 {
    return 2 * (rectangle.Width + rectangle.Height)
}

func (circle Circle) Area() float64 {
    return math.Pi * math.Pow(circle.Radius, 2)
}
