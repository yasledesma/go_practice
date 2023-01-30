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
        {"it should calculate the area of a rectangle", Rectangle{120, 50}, 120.00 * 50.00},
        {"it should calculate the area of a circle", Circle{12}, 12 * 12 * math.Pi},
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


