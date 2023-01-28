package main

import "testing"

func TestSum(t *testing.T) {
    t.Run("add an array of random size", func (t *testing.T) {
            got := Sum([]int {1, 2, 3, 4, 5, 6})
            want := 21

            assert(t, got, want)
    })
}


func assert(t testing.TB, got int, want int) {
    t.Helper()

    if got != want {
        t.Errorf("got %d, want %d.", got, want)
    }
}
