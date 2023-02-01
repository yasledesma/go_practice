package main

import "testing"

func TestSwap(t *testing.T) {
    x := 1
    y := 2

    want := [2]int{2, 1}
    
    Swap(&x, &y)

    if x != want[0] && y != want[1] {
        t.Errorf("got %v, want %v.", [2]int{x, y}, want)
    }
}
