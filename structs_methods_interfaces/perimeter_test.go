package shapes

import (
    "testing"
    "math"
)

func TestPerimeter(t *testing.T) {
    t.Run("it should calculate the perimeter of a rectangle", func(t *testing.T) {
        got := Rectangle{120, 50}.Perimeter()
        want := 2 * 120.00 + 2 * 50.00

        if got != want {
            t.Errorf("got %.2f, want %.2f.", got, want)
        }
    })

    t.Run("it should calculate the area of a rectangle", func(t *testing.T) {
        want := 120.00 * 50.00
        assertEquals(t, Rectangle{120, 50}, want)
    })

    t.Run("it should calculate the area of a circle", func(t *testing.T) {
        want := 12 * 12 * math.Pi
        assertEquals(t, Circle{12}, want)
    })
}

func assertEquals(t testing.TB, shape Shape, want float64) {
    t.Helper()
    got := shape.Area()
    if got != want {
        t.Errorf("got %g, want %g.", got, want)
    }
}
