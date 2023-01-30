package shapes

import (
    "testing"
    "math"
)

func TestPerimeter(t *testing.T) {
    testCases := []struct {
        name string
        shape Shape
        want float64
    } {
        {name: "it should calculate the area of a triangle", shape: Triangle{Base: 12, Height: 6}, want: 36.0},
        {name: "it should calculate the area of a rectangle", shape: Rectangle{Width: 120, Height: 50}, want: 120.00 * 50.00},
        {name: "it should calculate the area of a circle", shape: Circle{Radius: 12}, want: 12 * 12 * math.Pi},
    }

    for _, testCase := range testCases {
      t.Run(testCase.name, func(t *testing.T) {
            got := testCase.shape.Area()
            if got != testCase.want {
                t.Errorf("got %g, want %g.", got, testCase.want)
            }
        })
    }

    t.Run("it should calculate the perimeter of a rectangle", func(t *testing.T) {
        got := Rectangle{120, 50}.Perimeter()
        want := 2 * 120.00 + 2 * 50.00

        if got != want {
            t.Errorf("got %.2f, want %.2f.", got, want)
        }
    })
}


